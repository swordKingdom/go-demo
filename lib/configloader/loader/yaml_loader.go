package loader

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

//YamlLoader yaml配置文件加载对象
type YamlLoader struct {
	BaseConfLoader
	confPath string
}

//LoadConfigFromFile 读取配置文件
func (y *YamlLoader) LoadConfigFromFile(fileName string) {
	y.BaseConfLoader.lock.Lock()
	defer y.BaseConfLoader.lock.Unlock()
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

	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bs, &y.BaseConfLoader.confMap)
	if err != nil {
		panic(err)
	}
	y.confPath = path
}

//LoadConfigFromFileReader 从fileReader对象中读取配置
func (y *YamlLoader) LoadConfigFromFileReader(file *os.File) {
	y.BaseConfLoader.lock.Lock()
	defer y.BaseConfLoader.lock.Unlock()
	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(bs, &y.BaseConfLoader.confMap)
	if err != nil {
		panic(err)
	}
	y.confPath = file.Name()
}

func (y *YamlLoader) ReLoadConf() {
	y.LoadConfigFromFile(y.confPath)
}
