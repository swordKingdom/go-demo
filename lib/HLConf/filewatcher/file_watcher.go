package filewatcher

import (
	"log"

	"github.com/howeyc/fsnotify"
)

func RunFileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan bool)
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Watch("testDir")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
