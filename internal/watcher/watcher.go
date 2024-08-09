package watcher

import (
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"path/filepath"
	"sync"
)

type Watcher interface {
	Events() <-chan fsnotify.Event
	Errors() <-chan error
	Stop()
}

type watcher struct {
	mu      sync.Mutex
	watcher *fsnotify.Watcher
	pathSet mapset.Set[string]
}

func NewWatcher() Watcher {
	w := lo.Must(fsnotify.NewWatcher())
	set := mapset.NewSet[string]()
	for _, dir := range viper.GetStringSlice(macro.ConfigKeyIncludeDirs) {
		dir = lo.Must(filepath.Abs(dir))
		if !set.Contains(dir) {
			lo.Must0(w.Add(dir))
			set.Add(dir)
		}
	}
	wat := &watcher{
		watcher: w,
		pathSet: set,
	}
	viper.OnConfigChange(wat.onConfigChange)
	return wat
}

func (w *watcher) onConfigChange(_ fsnotify.Event) {
	w.mu.Lock()
	defer w.mu.Unlock()

	newSet := mapset.NewSet[string]()
	for _, dir := range viper.GetStringSlice(macro.ConfigKeyIncludeDirs) {
		dir, err := filepath.Abs(dir)
		if err != nil {
			continue
		}
		newSet.Add(dir)
	}
	remove := w.pathSet.Difference(newSet)
	add := newSet.Difference(w.pathSet)
	for dir := range remove.Iter() {
		_ = w.watcher.Remove(dir)
	}
	for dir := range add.Iter() {
		_ = w.watcher.Add(dir)
	}
	w.pathSet = newSet
}

func (w *watcher) Events() <-chan fsnotify.Event {
	return w.watcher.Events
}

func (w *watcher) Errors() <-chan error {
	return w.watcher.Errors
}

func (w *watcher) Stop() {
	w.watcher.Close()
}
