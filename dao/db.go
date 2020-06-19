package dao

import (
	"time"

	"github.com/stiwedung/blog/config"
	"github.com/stiwedung/blog/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stiwedung/libgo/log"
	"xorm.io/core"
	"xorm.io/xorm"
)

var db *xorm.Engine

func Connect() {
	if db != nil {
		return
	}
	mysqlURL := config.MysqlURL()
	if mysqlURL == "" {
		return
	}
	doConnect(mysqlURL)
}

func doConnect(url string) {
	var err error
	db, err = xorm.NewEngine("mysql", url)
	if err != nil {
		log.Fatalf("connect mysql failed: %v", err)
		return
	}
	db.SetLogger(&logger{
		Logger:    log.GetLogger(),
		isShowSQL: !config.Config.Common.ReleaseMode,
	})
	if err := db.Ping(); err != nil {
		log.Fatalf("mysql ping failed: %v", err)
		return
	}
	if _, err := db.Exec("use " + config.Config.DB.DBName); err != nil {
		log.Errorf("database %s not exist: %v", config.Config.DB.DBName, err)
		_, err = db.Exec("CREATE DATABASE " + config.Config.DB.DBName + " DEFAULT CHARACTER SET utf8mb4")
		if err != nil {
			log.Fatalf("create database %s error: %v", config.Config.DB.DBName, err)
			return
		}
	}
	db.Dialect().URI().DBName = config.Config.DB.DBName

	tblMapper := core.NewPrefixMapper(core.GonicMapper{}, "tbl_")
	db.SetTableMapper(tblMapper)
	db.SetColumnMapper(core.GonicMapper{})

	if err := db.Sync2(model.Models...); err != nil {
		log.Fatalf("create tables error: %v", err)
		return
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(30)
	db.SetConnMaxLifetime(5 * time.Minute)
}

func Disconnect() {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		log.Errorf("disconnect mysql error: %v", err)
	}
}

func Connected() bool {
	return db != nil
}
