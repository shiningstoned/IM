package initialize

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
)

func InitMysql() *gorm.DB {
	username := GetConfig().Mysql.Username
	password := GetConfig().Mysql.Password
	host := GetConfig().Mysql.Host
	port := GetConfig().Mysql.Port
	DB := GetConfig().Mysql.DB
	dns := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", DB, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		klog.Fatalf("init mysql failed: %s", err.Error())
	}
	return db
}

func InitRedis() *redis.Client {
	addr := GetConfig().Redis.Host
	password := GetConfig().Redis.Password
	DB := GetConfig().Redis.DB
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       DB,
	})
	return client
}
