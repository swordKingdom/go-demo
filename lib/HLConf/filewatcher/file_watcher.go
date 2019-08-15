package filewatcher

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

//ProcessFunc 监听事件的处理函数
type ProcessFunc func(*watcher.Watcher)

//RunFileWatcher 启动文件监控
func RunFileWatcher(fileName string, processer ProcessFunc, interval time.Duration) {
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Write)
	w.Add(fileName)
	if err := w.Start(interval); err != nil {
		log.Fatalln(err)
	}
	go processer(w)
}
