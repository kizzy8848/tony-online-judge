package models

import (
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

var RDB = InitRedisDB()

func Init() *gorm.DB {
	dsn :="root:admin@tcp(127.0.0.1:3306)/toj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	return db
}

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //  password 884899
		DB:       0,        // use default DB
	})
}
