package db

import (
	"cyberpark/database/models"
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:greed9527@tcp(localhost:3306)/cyberpark?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	sqldb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	sqldb.SetConnMaxLifetime(time.Hour) // 每條連線的存活時間
	sqldb.SetMaxOpenConns(8)            // 最大連線數
	sqldb.SetMaxIdleConns(6)            // 最大閒置連線數

	return db
}
