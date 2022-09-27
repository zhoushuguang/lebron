package orm

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Config struct {
	DSN         string
	Active      int
	Idle        int
	IdleTimeout time.Duration
}

func NewMysql(c *Config) *gorm.DB {
	if c == nil {
		panic("config cannot be nil")
	}

	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}

	db.DB().SetMaxIdleConns(c.Idle)
	db.DB().SetMaxOpenConns(c.Active)
	db.DB().SetConnMaxLifetime(c.IdleTimeout)

	return db
}
