package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Service struct {
		Dingding struct {
			Webhook string `yaml:"webhook" json:"webhook"`
		}
		Feishu struct {
			ApiUrl    string `yaml:"apiUrl" json:"apiUrl"`
			AppId     string `yaml:"appId" json:"appId"`
			AppSecret string `yaml:"appSecret" json:"appSecret"`
		}
	}
}

func (c *Config) GetConfig() *Config {
	ymlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln("读取配置文件错误: ", err.Error())
	}
	if err := yaml.Unmarshal(ymlFile, &c); err != nil {
		log.Fatalln("配置文件解析错误: ", err.Error())
	}
	return c
}
