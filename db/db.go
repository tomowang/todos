package db

import (
	"todos/config"
	"todos/core"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()
	var err error
	db, err = gorm.Open(sqlite.Open(c.Database.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&core.User{})
	db.AutoMigrate(&core.Todo{})
}

func GetDB() *gorm.DB {
	return db
}
