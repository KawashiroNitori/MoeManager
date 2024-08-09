package manager

import (
	"context"
	"fmt"
	"github.com/KawashiroNitori/MoeManager/internal/dao"
	daoImpl "github.com/KawashiroNitori/MoeManager/internal/dao/impl"
	"github.com/KawashiroNitori/MoeManager/internal/ent"
	"github.com/KawashiroNitori/MoeManager/internal/ent/picture"
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	"github.com/KawashiroNitori/MoeManager/internal/renamer"
	"github.com/KawashiroNitori/MoeManager/internal/upscaler"
	"github.com/KawashiroNitori/MoeManager/internal/util"
	"github.com/KawashiroNitori/MoeManager/internal/watcher"
	"github.com/alexflint/go-filemutex"
	"github.com/fsnotify/fsnotify"
	"github.com/kardianos/service"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Event struct {
	fsnotify.Event
	Time time.Time
}

type Manager interface {
	service.Interface
}

type manager struct {
	mu      sync.Mutex
	started bool
	stopped bool
	stop    chan struct{}
	done    chan struct{}

	logger service.Logger

	watcher  watcher.Watcher
	renamer  renamer.Renamer
	upscaler upscaler.Upscaler

	pictureDAO dao.PictureDAO
}

func NewManager() Manager {
	return &manager{
		pictureDAO: daoImpl.NewPictureDAO(),
		renamer:    renamer.NewRenamer(),
		upscaler:   upscaler.NewUpscaler(),
	}
}

func (m *manager) Start(s service.Service) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.started {
		return fmt.Errorf("manager has been started")
	}
	if m.stopped {
		return fmt.Errorf("manager has been stopped")
	}
	m.started = true
	m.logger = lo.Must(s.Logger(nil))
	m.stop = make(chan struct{})
	m.done = make(chan struct{})
	ctx := context.Background()
	go m.Run(ctx)

	return nil
}

func (m *manager) Run(ctx context.Context) {
	queue := make(chan Event, 1024)
	go func() {
		defer close(m.done)

		for {
			select {
			case _, ok := <-m.stop:
				if !ok {
					_ = m.logger.Info("Manager stopped.")
					return
				}
			case event := <-queue:
				if err := m.Workflow(ctx, event); err != nil {
					_ = m.logger.Errorf("workflow %s error occured: %v", event.Name, err)
				}
			}
		}
	}()

	for _, includeDir := range viper.GetStringSlice(macro.ConfigKeyIncludeDirs) {
		entries := lo.Must(os.ReadDir(includeDir))
		for _, entry := range entries {
			select {
			case _, ok := <-m.stop:
				if !ok {
					queue = nil
					return
				}
			default:
			}
			if entry.IsDir() {
				continue
			}
			filename := filepath.Join(includeDir, entry.Name())
			queue <- Event{Event: fsnotify.Event{Name: filename, Op: fsnotify.Create}}
		}
	}

	m.watcher = watcher.NewWatcher()
	for {
		select {
		case _, ok := <-m.stop:
			if !ok {
				queue = nil
				return
			}
		case err := <-m.watcher.Errors():
			if err != nil {
				_ = m.logger.Errorf("watcher error occured: %v", err)
			}
		case event := <-m.watcher.Events():
			pic := m.pictureDAO.GetPictureByFilename(ctx, event.Name)
			if pic != nil && pic.Status == picture.StatusProcessing {
				continue
			}
			queue <- Event{Event: event, Time: time.Now()}
		}
	}
}

func (m *manager) Workflow(ctx context.Context, event Event) error {
	if event.Op&(fsnotify.Create|fsnotify.Remove|fsnotify.Write) == 0 {
		return nil
	}
	_ = m.logger.Infof(event.String())
	path, err := filepath.Abs(event.Name)
	if err != nil {
		return err
	}
	if util.IsDir(path) || !util.IsExists(path) {
		return nil
	}
	if event.Op.Has(fsnotify.Remove) {
		return m.Remove(ctx, path)
	}
	if !util.IsSupportedExtension(util.GetAllSupportedExtensions(), path) {
		return nil
	}
	var fileMu *filemutex.FileMutex
	for {
		fileMu, err = filemutex.New(path)
		if err == nil {
			break
		}
	}
	if err := fileMu.RLock(); err != nil {
		return err
	}
	fileMu.Close()

	if event.Op.Has(fsnotify.Create) {
		if err := m.Create(ctx, fileMu, path); err != nil {
			return err
		}
	}
	if event.Op.Has(fsnotify.Write) {
		if err := m.Digest(ctx, path); err != nil {
			return err
		}
	}
	return nil
}

func (m *manager) Create(ctx context.Context, fileMu *filemutex.FileMutex, path string) (err error) {
	filename := filepath.Base(path)
	pic := m.pictureDAO.GetPictureByFilename(ctx, filename)
	digest, err := util.GetDigest(path)
	if err != nil {
		return err
	}
	if pic != nil {
		return nil
	}
	width, height, _ := util.GetImageSize(path)
	pic, err = m.pictureDAO.Create(ctx, &ent.Picture{
		Filename:         filename,
		OriginalFilename: filename,
		Dir:              filepath.Dir(path),
		Digest:           digest,
		OriginalWidth:    width,
		OriginalHeight:   height,
		Status:           picture.StatusProcessing,
		CreatedAt:        util.LastModify(path),
	})
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			pic.Status = picture.StatusDone
		} else {
			pic.Status = picture.StatusError
			pic.ErrorMessage = err.Error()
		}
		pic, _ = pic.Update().
			SetStatus(pic.Status).
			SetErrorMessage(pic.ErrorMessage).
			Save(ctx)
	}()
	if pic, err = m.renamer.Rename(ctx, path, fileMu, pic); err != nil {
		return err
	}
	path = filepath.Join(filepath.Dir(path), pic.Filename)
	if pic, err = m.upscaler.Upscale(ctx, path, pic); err != nil {
		return err
	}

	return nil
}

func (m *manager) Remove(ctx context.Context, path string) error {
	filename := filepath.Base(path)
	pic := m.pictureDAO.GetPictureByFilename(ctx, filename)
	if pic == nil || pic.Status == picture.StatusProcessing {
		return nil
	}
	return m.pictureDAO.Remove(ctx, filename)
}

func (m *manager) Digest(ctx context.Context, path string) error {
	filename := filepath.Base(path)
	pic := m.pictureDAO.GetPictureByFilename(ctx, filename)
	if pic == nil {
		return nil
	}
	if pic.Status == picture.StatusProcessing {
		return nil
	}
	digest, err := util.GetDigest(path)
	if err != nil {
		return err
	}
	if pic.Digest == digest {
		return nil
	}
	pic, err = pic.Update().SetDigest(digest).Save(ctx)
	return err
}

func (m *manager) Stop(s service.Service) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !m.started {
		return fmt.Errorf("manager has not been started")
	}
	if m.stopped {
		return fmt.Errorf("manager has been stopped")
	}

	if m.watcher != nil {
		m.watcher.Stop()
	}
	close(m.stop)
	<-m.done

	lo.Must0(m.pictureDAO.Close())
	m.stopped = true

	return nil
}
