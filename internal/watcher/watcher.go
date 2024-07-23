package watcher

import (
	"github.com/KawashiroNitori/MoeManager/internal/macro"
	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

type Watcher interface {
	Events() <-chan fsnotify.Event
	Errors() <-chan error
	Stop()
}

type watcher struct {
	watcher *fsnotify.Watcher
}

func NewWatcher() Watcher {
	w := lo.Must(fsnotify.NewWatcher())
	for _, dir := range viper.GetStringSlice(macro.ConfigKeyIncludeDirs) {
		lo.Must0(w.Add(dir))
	}
	return &watcher{
		watcher: w,
	}
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
