package utils

import "sync"

type Observer interface {
	NotifyCallback(event interface{})
}

type Observable interface {
	Add(observer Observer)
	Remove(observer Observer)
	Notify(event interface{})
}

type Watcher struct {
	sync.Mutex
	observers map[Observer]struct{}
}

func NewWatcher() *Watcher {
	return &Watcher{
		observers: make(map[Observer]struct{}),
	}
}

func (w *Watcher) Add(observer Observer) {
	w.Lock()
	w.observers[observer] = struct{}{}
	w.Unlock()
}

func (w *Watcher) Remove(observer Observer) {
	w.Lock()
	delete(w.observers, observer)
	w.Unlock()
}

func (w *Watcher) Notify(event interface{}) {
	w.Lock()
	defer w.Unlock()

	for w := range w.observers {
		w.NotifyCallback(event)
	}
}
