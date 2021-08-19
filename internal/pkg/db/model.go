package db

import (
	"encoding/gob"

	"github.com/SYSU-ECNC/workspace-be/internal/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NetID string `gorm:"unique"`
	Name  string
	Level int
	Email string
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open(postgres.Open(config.Get("db_dsn")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	gob.Register(&User{})
}
