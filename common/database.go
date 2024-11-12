package common

import (
	"go-test-basic/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gloablDB *gorm.DB

func GetDB() *gorm.DB {
	return gloablDB
}

func NewDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/hello?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),
		TranslateError: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Item{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}

func InitDB() {
	gloablDB = NewDB()
}
