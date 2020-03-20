package repository

import (
	"log"
	"time"
    
	
	_ "github.com/lib/pq"
	
	"github.com/teamlint/ardan/config"
	"xorm.io/xorm"
)

func NewDB() (*xorm.Engine, error) {
	db, err := xorm.NewEngine(config.Databases("Ardan").DriverName, config.Databases("Ardan").ConnString)
	if err != nil {
		return nil, err
	}
	db.SetMapper(core.GonicMapper{})
	// ping
	err = db.Ping()
	if err != nil {
		log.Fatalf("[repository.NewDB] err=%v\n", err)
	}
	// conn
	db.SetMaxOpenConns(config.Databases("Ardan").MaxOpenConns)
	db.SetMaxIdleConns(config.Databases("Ardan").MaxIdleConns)
	dur, err := time.ParseDuration(config.Databases("Ardan").ConnMaxLifetime)
	if err != nil {
		log.Fatalf("[repository.NewDB] SetConnMaxLifetime err=%v\n", err)
	}
	db.SetConnMaxLifetime(dur)
	// log
	if config.App().Debug {
		db.ShowSQL(true)
	}

	return db, err
}
