package confhotload

import (
	"log"
	"this_is_a_explame/lib/configloader"
	"time"

	"this_is_a_explame/lib/configloader/loader"
	"this_is_a_explame/lib/watcher/filewatcher"

	"github.com/radovskyb/watcher"
)

type HotLoadingConf struct {
	loader.ConfLoader
	w *filewatcher.FileWatcher
}

func (h *HotLoadingConf) Init(fileName string) error {
	conf, err := configloader.LoadConfig(fileName)
	if err != nil {
		return err
	}
	h.ConfLoader = conf
	pFunc := func(w *watcher.Watcher) {
		for {
			select {
			case event := <-w.Event:
				if event.Op == watcher.Write {
					h.ConfLoader.ReLoadConf()
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}
	h.w = filewatcher.RunFileWatcher(conf.GetFileName(), pFunc, 1*time.Second)
	return nil
}
