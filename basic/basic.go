package basic

import (
	"books/basic/config"
	"books/basic/db"
	"books/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
