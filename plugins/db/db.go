package db

import (
	"books/basic"
	"database/sql"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func init() {
	basic.Register(initDB)
}

// Init 初始化数据库
func initDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db 已经初始化过")
		log.Logf(err.Error())
		return
	}

	// 如果配置声明使用mysql
	initMysql()

	inited = true
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}
