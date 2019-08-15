package configloadder

import (
	"fmt"
	"os"
	"strings"
)

const (
	EnvConfBasePath = "ENV_CONF_FILE"
	YamlConfType    = "yaml"
	XmlConfType     = "xml"
	TomlConfType    = "toml"

	FileNameDefultStepSep = "."
)

type ConfLoader interface {
	ConfLoaderOp
	LoadConfigFromFile(file string)
	LoadConfigFromFileReader(file *os.File)
}

type ConfLoaderOp interface {
	GetInt(key string, defultValue int) int
	GetBool(key string, defultValue bool) bool
	GetString(key string, defultValue string) string
	GetFloat(key string, defultValue float32) float32
}

type BaseConfLoader struct {
	confMap map[interface{}]interface{}
}

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
	case XmlConfType:
		loader := &XmlLoader{}
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
