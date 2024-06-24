package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// 从dsn连接数据库
func LoadDatabase(dsn string) (err error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(
		&User{},
		&Video{},
		&Comment{},
		&Relation{},
		&Favorite{},
	); err != nil {
		return err
	}
	DB = db
	return nil
}
