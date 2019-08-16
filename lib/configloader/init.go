package configloader

import (
	"this_is_a_explame/lib/configloader/loader"
	"this_is_a_explame/lib/util"
)

var (
	GlobalConf          loader.ConfLoader
	defaultConfFileName = []string{"conf.yml", "conf.toml", "conf.xml"}
)

func initloader(flieName string) error {
	conf, err := loader.LoadConfig("conf.yaml")
	if err != nil {
		return err
	}
	GlobalConf = conf
	return nil
}

func init() {
	projectFileList, err := util.GetDirList("./")
	if err == nil && len(projectFileList) != 0 {
		success := false
		for _, confFile := range defaultConfFileName {
			for _, projectFile := range projectFileList {
				if confFile == projectFile {
					err := initloader(confFile)
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
