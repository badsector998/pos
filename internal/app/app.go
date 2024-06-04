package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
}

func NewApp() *App {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &App{}
}
