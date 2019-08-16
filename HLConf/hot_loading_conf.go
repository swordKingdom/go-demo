package HLConf

import (
	"errors"
	"log"
	"time"

	"this_is_a_explame/lib/configloader/loader"
	"this_is_a_explame/lib/watcher/filewatcher"

	"github.com/radovskyb/watcher"
)

type HotLoadingConf struct {
	loader loader.ConfLoader
	w      *filewatcher.FileWatcher
}

func (h *HotLoadingConf) Init(loader loader.ConfLoader) error {
	if loader == nil {
		return errors.New("conf loader is nil")
	}
	h.loader = loader
	pFunc := func(w *watcher.Watcher) {
		for {
			select {
			case event := <-w.Event:
				if event.Op == watcher.Write {
					h.loader.ReLoadConf()
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}
	h.w = filewatcher.RunFileWatcher(loader.GetFileName(), pFunc, 1*time.Second)
	return nil
}
