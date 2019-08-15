package loader

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

const (
	//EnvConfBasePath 默认配置文件路径，配置在系统环境变量中
	EnvConfBasePath = "ENV_CONF_FILE"
	//YamlConfType yaml文件类型
	YamlConfType = "yaml"
	//XMLConfType xml文件类型
	XMLConfType = "xml"
	//TomlConfType toml文件类型
	TomlConfType = "toml"

	//FileNameDefultStepSep 文件后缀分割符号
	FileNameDefultStepSep = "."
)

//ConfLoader 配置加载器接口
type ConfLoader interface {
	ConfLoaderOp
	LoadConfigFromFile(file string)
	LoadConfigFromFileReader(file *os.File)
	ReLoadConf()
}

//ConfLoaderOp 获取配置对象操作的接口
type ConfLoaderOp interface {
	GetInt(key string, defultValue int) int
	GetBool(key string, defultValue bool) bool
	GetString(key string, defultValue string) string
	GetFloat(key string, defultValue float32) float32
}

//BaseConfLoader 配置文件对象操作的实现对象
type BaseConfLoader struct {
	confMap map[interface{}]interface{}
	lock    *sync.Mutex
}

//GetInt 获取Int类型的配置参数
func (b *BaseConfLoader) GetInt(key string, defultValue int) int {
	if value, ok := b.confMap[key]; ok {
		intValue, ok := value.(int)
		if ok {
			return intValue
		}
		return defultValue
	}
	return defultValue
}

//GetBool 获取bool类型的配置参数
func (b *BaseConfLoader) GetBool(key string, defultValue bool) bool {
	if value, ok := b.confMap[key]; ok {
		bValue, ok := value.(bool)
		if ok {
			return bValue
		}
		return defultValue
	}
	return defultValue
}

//GetString 获取string类型的配置参数
func (b *BaseConfLoader) GetString(key string, defultValue string) string {
	if value, ok := b.confMap[key]; ok {
		bValue, ok := value.(string)
		if ok {
			return bValue
		}
		return defultValue
	}
	return defultValue
}

//GetFloat 获取float32类型的配置参数
func (b *BaseConfLoader) GetFloat(key string, defultValue float32) float32 {
	if value, ok := b.confMap[key]; ok {
		fValue, ok := value.(float32)
		if ok {
			return fValue
		}
		return defultValue
	}
	return defultValue
}

//LoadConfig 加载配置文件
func LoadConfig(file string) (ConfLoader, error) {
	tmpStr := strings.Split(file, FileNameDefultStepSep)
	if len(tmpStr) < 2 {
		return nil, fmt.Errorf("conf file name:%v  error", file)
	}
	confType := tmpStr[len(tmpStr)-1]
	switch confType {
	case YamlConfType:
		loader := &YamlLoader{}
		loader.LoadConfigFromFile(file)
		return loader, nil
	case XMLConfType:
		loader := &XMLLoader{}
		loader.LoadConfigFromFile(file)
		return loader, nil
	case TomlConfType:
		loader := &TomlLoader{}
		loader.LoadConfigFromFile(file)
		return loader, nil
	default:
		return nil, fmt.Errorf("unsupport conf type: %v", confType)
	}
}
