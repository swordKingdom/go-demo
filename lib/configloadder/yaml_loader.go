package configloadder

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

type YamlLoader struct {
	BaseConfLoader
	confPath string
}

func (y *YamlLoader) LoadConfigFromFile(fileName string) {
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

func (y *YamlLoader) LoadConfigFromFileReader(file *os.File) {
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
