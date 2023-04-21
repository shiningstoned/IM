package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	"log"
)

type MyConfig struct {
	Mysql MysqlConfig
}

var c MyConfig

type MysqlConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       string
}

func GetConfig() MyConfig {
	return c
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./message/config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("init config failed, err: %s", err.Error())
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		klog.Fatalf("unmarshal config failed: %s", err.Error())
	}
}
