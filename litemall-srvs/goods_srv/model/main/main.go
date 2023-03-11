package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"litemall-srvs/goods_srv/model"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:mysqlpw@tcp(127.0.0.1:55002)/litemall_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(any(err))
	}

	_ = db.AutoMigrate(&model.Category{}, &model.Brand{}, &model.Goods{})
	_ = db.AutoMigrate(&model.GoodsAttribute{})
	_ = db.AutoMigrate(&model.GoodsProduct{})
	_ = db.AutoMigrate(&model.GoodsSpecification{})
	_ = db.AutoMigrate(&model.Banner{})
}
