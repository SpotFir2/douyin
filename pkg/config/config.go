package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	Confi *Config
)

type Config struct {
	Site struct {
		IP     string `yaml:"ip"`
		Port   int    `yaml:"port"`
		Domain string `yaml:"domain"`
	} `yaml:"site"`
	MySQL struct {
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		IP       string `yaml:"ip"`
		Port     int    `yaml:"port"`
	} `yaml:"mysql"`
	Salt string `yaml:"salt"`
	v    *viper.Viper
}

/*
InitConfig 初始化配置文件
path 预设配置文件地址
*/
func InitConfig(path string) {
	c := &Config{}
	c.Site.IP = "0.0.0.0"
	c.Site.Port = 8080
	c.Site.Domain = "localhost:8080"
	c.MySQL.Name = "douyin"
	c.MySQL.User = "douyin"
	c.MySQL.Password = "douyin"
	c.Salt = "douyin"
	data, err := yaml.Marshal(c)
	if err != nil {
		log.Println("生成配置文件失败:", err)
		return
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		log.Println("生成配置文件失败:", err)
		return
	}
	log.Println("按 ", path, " 格式建立配置文件")
}

/*
LoadConfig 加载配置文件
path 配置文件地址
*/
func LoadConfig(path string) (err error) {
	c := &Config{}
	c.v = viper.New()
	c.v.SetConfigFile(path)
	if err = c.v.ReadInConfig(); err != nil {
		return
	}
	if err = c.v.Unmarshal(c); err != nil {
		return
	}
	Confi = c
	return nil
}
