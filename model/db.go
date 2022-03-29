package model

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myBlog/config"
	"time"
)

var (
	db  *gorm.DB
	err error
)

// Init 初始化数据库连接
func Init() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.Name,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("connect to mysql error")
	}

	db.AutoMigrate(&User{}, &Article{}, &Category{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Error("get sql.DB error")
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
