package confhotload

import (
	"fmt"
	"testing"
	"time"
)

func TestHotLoadingConf_Iint(t *testing.T) {
	hotLoader := &HotLoadingConf{}
	err := hotLoader.Init("conf.yml")
	if err != nil {
		return
	}
	msg := hotLoader.GetString("msg", "")
	fmt.Println(msg)
	time.Sleep(30 * time.Second)
	msg = hotLoader.GetString("msg", "")
	fmt.Println(msg)
}
