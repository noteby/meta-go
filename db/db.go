package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"meta-go/model"
)

var conn *gorm.DB

func Init() {
	var err error
	conn, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@123@tcp(127.0.0.1:3306)/meta_go?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Panic(err)
	}
	Migrate()

}

func Conn() *gorm.DB {
	return conn
}

func Migrate() {
	conn.AutoMigrate(
		&model.User{},
		&model.Star{},
		&model.Media{},
	)
}
