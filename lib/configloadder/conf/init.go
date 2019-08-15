package conf

import (
	"awesomeProject12/configloadder"
	"awesomeProject12/file"
)

var (
	GlobalConf          configloadder.ConfLoader
	defaultConfFileName = []string{"conf.yml", "conf.toml", "conf.xml"}
)

func initWithFileName(flieName string) error {
	conf, err := configloadder.LoadConfig("conf.yaml")
	if err != nil {
		return err
	}
	GlobalConf = conf
	return nil
}

func init() {
	projectFileList, err := file.GetDirList("./")
	if err == nil && len(projectFileList) != 0 {
		success := false
		for _, confFile := range defaultConfFileName {
			for _, projectFile := range projectFileList {
				if confFile == projectFile {
					err := initWithFileName(confFile)
					if err != nil {
						continue
					}
				}
			}
		}
		if !success {
			//TODO：抛出异常
		}
	} else {
		//TODO：抛出异常
	}
}
