package conf

import (
	"flag"

	"github.com/apache/dubbo-go/common/yaml"
)

const defaultConfigFile = "../conf/data.yml"

func Init() (config *AppConfig, err error) {
	var (
		b          []byte
		configFile string
	)
	// 1.Specify the configuration file using command line parameters
	flag.StringVar(&configFile, "dataConf", "../conf/data.yml", "choose conf file.")
	flag.Parse()

	if configFile == "" {
		configFile = defaultConfigFile
	}

	// 2.Load yml config
	if b, err = yaml.LoadYMLConfig(configFile); err != nil {
		return
	}

	// 3.Serialize the contents of the configuration file to struct
	config = new(AppConfig)
	if err = yaml.UnmarshalYML(b, config); err != nil {
		return
	}
	return
}
