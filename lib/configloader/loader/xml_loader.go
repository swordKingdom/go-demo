package loader

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"os"
)

//XMLLoader xml配置文件加载队对象
type XMLLoader struct {
	BaseConfLoader
	confPath string
}

//LoadConfigFromFile 读取配置
//TODO：xml读取
func (x *XMLLoader) LoadConfigFromFile(fileName string) {
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

//LoadConfigFromFileReader 从fileReader对象中读取对象
func (x *XMLLoader) LoadConfigFromFileReader(file *os.File) {
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

func (x *XMLLoader) ReLoadConf() {
	x.LoadConfigFromFile(x.confPath)
}
