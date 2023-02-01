package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"litemall_servers/user_server/model"

	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"crypto/md5"
	"crypto/sha512"

	"github.com/anaskhan96/go-password-encoder"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	dsn := "root:mysqlpw@tcp(127.0.0.1:55000)/litemall_user_server?charset=utf8mb4&parseTime=True&loc=Local"

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
		panic(err)
	}

	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("admin123", options)
	//newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(newPassword)
	//
	//for i := 0; i < 10; i++ {
	//	user := model.User{
	//		NickName: fmt.Sprintf("bobby%d", i),
	//		Mobile:   fmt.Sprintf("1878222222%d", i),
	//		Password: newPassword,
	//	}
	//	db.Save(&user)
	//}

	////设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	////sql才是最重要的，本着这个原则我尽量的给大家看到每个api背后的sql语句是什么
	//
	//定义一个表结构， 将表结构直接生成对应的表 - migrations
	// 迁移 schema
	//此处应该有sql语句
	_ = db.AutoMigrate(&model.User{})
	_ = db.AutoMigrate(&model.Address{})
	_ = db.AutoMigrate(&model.Cart{})
	_ = db.AutoMigrate(&model.Collect{})
	_ = db.AutoMigrate(&model.Comment{})
	_ = db.AutoMigrate(&model.Coupon{})
	_ = db.AutoMigrate(&model.CouponUser{})
	_ = db.AutoMigrate(&model.Feedback{})
	_ = db.AutoMigrate(&model.Footprint{})
	_ = db.AutoMigrate(&model.SearchHistory{})

	fmt.Println(genMd5("xxxxx_123456"))

	// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(len(newPassword))
	fmt.Println(newPassword)

	//passwordInfo := strings.Split(newPassword, "$")
	//fmt.Println(passwordInfo)
	//check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check) // true
}
