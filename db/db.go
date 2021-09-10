package db

import (
	"todos/config"
	"todos/core"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()
	var err error
	if c.DB_Driver == "sqlite" {
		db, err = gorm.Open(sqlite.Open(c.DB_DSN), &gorm.Config{})
	} else if c.DB_Driver == "postgres" {
		db, err = gorm.Open(postgres.Open(c.DB_DSN), &gorm.Config{})
	}
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&core.User{})
	db.AutoMigrate(&core.Todo{})
}

func GetDB() *gorm.DB {
	return db
}
