package HLConf

import (
	"fmt"
	"testing"
	"time"

	"this_is_a_explame/lib/configloader/loader"
)

func TestHotLoadingConf_Iint(t *testing.T) {
	hotLoader := &HotLoadingConf{}
	conf, err := loader.LoadConfig("conf.yml")
	if err != nil {

	}
	hotLoader.Iint(conf)
	msg := conf.GetString("msg", "")
	fmt.Println(msg)
	time.Sleep(30 * time.Second)
	msg = conf.GetString("msg", "")
	fmt.Println(msg)
}
