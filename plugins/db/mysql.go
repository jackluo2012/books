package db

import (
	"books/basic/config"
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type db struct {
	Mysql Mysql `json："mysql"`
}

// Mysql mySQL 配置
type Mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	MaxIdleConnection int    `json:"maxIdleConnection"`
	MaxOpenConnection int    `json:"maxOpenConnection"`
}

func initMysql() {
	log.Info("[initMysql] 初始化Mysql")

	c := config.C()
	cfg := &db{}

	err := c.App("db", cfg)
	if err != nil {
		log.Warnf("[initMysql] %s", err)
	}

	if !cfg.Mysql.Enable {
		log.Warnf("[initMysql] 未启用Mysql")
		return
	}

	// 创建连接
	mysqlDB, err = sql.Open("mysql", cfg.Mysql.URL)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(cfg.Mysql.MaxOpenConnection)

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(cfg.Mysql.MaxIdleConnection)

	//连接数据库闲置断线的问题
	mysqlDB.SetConnMaxLifetime(time.Second * 100)
	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Warnf("[initMysql] Mysql 连接成功")
}
