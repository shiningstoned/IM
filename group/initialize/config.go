package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	"log"
)

var c MyConfig

func GetConfig() MyConfig {
	return c
}

type MyConfig struct {
	Mysql MysqlConfig
	Redis RedisConfig
}

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       string
}

type RedisConfig struct {
	Host     string
	Password string
	DB       int
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./group/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("load config failed")
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		klog.Fatal("unmarshal config failed")
	}
}
