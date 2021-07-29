package main

import (
	"todos/config"
	"todos/db"
)

func main() {
	config.Init("")
	db.Init()
	c := config.GetConfig()
	router := NewRouter()
	router.Run(c.Listen.TCP)
}
