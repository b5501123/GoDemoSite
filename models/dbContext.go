package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

const (
	UserName     string = "aeon"
	Password     string = "aeon12345"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "GoDemo"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

func CreateDataBase() {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}
