package app

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
}

func NewApp() *App {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Panic("error opening db")
	}

	return &App{}
}
