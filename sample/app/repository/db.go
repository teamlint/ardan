package repository

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/teamlint/ardan/config"
	"github.com/teamlint/ardan/server"
	"xorm.io/xorm"
)

func NewDB() (*xorm.Engine, error) {
	db, err := xorm.NewEngine(config.Databases("UserDB").DriverName, config.Databases("UserDB").ConnString)
	if err != nil {
		return nil, err
	}
	// ping
	err = db.Ping()
	if err != nil {
		log.Printf("[repository.NewDB] err=%v\n", err)
	}
	// conn
	db.SetMaxOpenConns(10000)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(3 * time.Minute)
	// log
	if server.Mode() == server.DebugMode {
		db.ShowSQL(true)
	}
	// db.SetLogLevel(core.LOG_DEBUG)
	return db, err
}
