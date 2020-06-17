package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Database struct {
		Host     string
		Port     uint
		User     string
		Password string
		DbName   string
	}
	Server struct {
		Port string
	}
}

var c = new(Config)
var one sync.Once

func InitConfig(confPath string) {
	one.Do(func() {
		//vi := viper.New()
		viper.SetConfigFile(confPath)
		//viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("读取配置文件失败:%s\n", err.Error())
			panic("读取配置文件失败")
		}
		//fmt.Printf("参数：%v\n", viper.AllSettings())
		if err := viper.Unmarshal(c); err != nil {
			fmt.Printf("序列化配置文件失败：%s\n", err.Error())
			panic("序列化配置文件失败")
		}
	})
}

func GetConfig() *Config {
	return c
}
