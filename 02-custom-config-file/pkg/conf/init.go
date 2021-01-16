package conf

import (
	"flag"
	"fmt"

	"github.com/apache/dubbo-go/common/yaml"
)

const defaultConfigFile = "../conf/data.yml"

func Init() (config *DataConfig, err error) {
	var (
		b          []byte
		configFile string
	)
	flag.StringVar(&configFile, "dataConf", "../conf/data.yml", "choose conf file.")
	flag.Parse()
	fmt.Printf("configFile :%v", configFile)

	if configFile == "" {
		configFile = defaultConfigFile
	}
	if b, err = yaml.LoadYMLConfig(configFile); err != nil {
		return
	}
	config = new(DataConfig)
	if err = yaml.UnmarshalYML(b, config); err != nil {
		return
	}
	return
}
