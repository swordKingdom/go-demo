package loader

import (
	"os"

	"github.com/BurntSushi/toml"
)

//TomlLoader toml配置文件加载对象
type TomlLoader struct {
	BaseConfLoader
	confPath string
}

//LoadConfigFromFile 读取配置
//TODO：toml读取
func (t *TomlLoader) LoadConfigFromFile(fileName string) {
	t.BaseConfLoader.lock.Lock()
	defer t.BaseConfLoader.lock.Unlock()
	path := os.Getenv(EnvConfBasePath)
	if path == "" {
		path = EnvConfBasePath
	}
	if info, _ := os.Stat(path); info == nil {
		return
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err := toml.DecodeReader(file, &t.BaseConfLoader.confMap); err != nil {
		// handle error
	}
	t.confPath = path
}

//LoadConfigFromFileReader 通过fileReader对象读取对象
func (t *TomlLoader) LoadConfigFromFileReader(file *os.File) {
	t.BaseConfLoader.lock.Lock()
	defer t.BaseConfLoader.lock.Unlock()
	if _, err := toml.DecodeReader(file, &t.BaseConfLoader.confMap); err != nil {
		// handle error
	}
	t.confPath = file.Name()
}

func (t *TomlLoader) ReLoadConf() {
	t.LoadConfigFromFile(t.confPath)
}
