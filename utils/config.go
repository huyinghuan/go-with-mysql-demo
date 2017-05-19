package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//DbConfig 数据库配置
type DbConfig struct {
	Driver  string
	Connect string
}

//Config 配置文件
type Config struct {
	Db  DbConfig
	Dev bool
}

//ReadConfig 读取配置文件
func ReadConfig(source string, config *Config) error {
	configBytes, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return err
	}
	return nil
}

//GetConfig 获取配置文件
func GetConfig() (config Config) {
	//configFilePath, _ := filepath.Abs("../conf/config.yaml")
	ReadConfig("conf/config.yaml", &config)
	return config
}
