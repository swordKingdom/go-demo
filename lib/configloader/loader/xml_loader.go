package loader

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type XmlLoader struct {
	BaseConfLoader
	confPath string
}

//TODO：xml读取
func (x *XmlLoader) LoadConfigFromFile(fileName string) {
	x.BaseConfLoader.lock.Lock()
	defer x.BaseConfLoader.lock.Unlock()
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
	err = xml.Unmarshal(bs, &x.BaseConfLoader.confMap)
	if err != nil {
		panic(err)
	}
	x.confPath = path
}

func (x *XmlLoader) LoadConfigFromFileReader(file *os.File) {
	x.BaseConfLoader.lock.Lock()
	defer x.BaseConfLoader.lock.Unlock()
	bs, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(bs, &x.BaseConfLoader.confMap)
	if err != nil {
		panic(err)
	}
	x.confPath = file.Name()
}
