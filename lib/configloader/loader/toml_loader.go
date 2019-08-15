package loader

import (
	"os"

	"github.com/BurntSushi/toml"
)

type TomlLoader struct {
	BaseConfLoader
	confPath string
}

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

func (t *TomlLoader) LoadConfigFromFileReader(file *os.File) {
	t.BaseConfLoader.lock.Lock()
	defer t.BaseConfLoader.lock.Unlock()
	if _, err := toml.DecodeReader(file, &t.BaseConfLoader.confMap); err != nil {
		// handle error
	}
	t.confPath = file.Name()
}
