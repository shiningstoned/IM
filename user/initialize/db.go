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
	database := GetConfig().Mysql.DB
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
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
	host := GetConfig().Redis.Host
	password := GetConfig().Redis.Password
	db := GetConfig().Redis.DB
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})
	return client
}
