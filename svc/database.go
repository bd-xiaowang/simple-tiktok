package svc

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"time"
)

// DB 定义一个全局对象
var DB *gorm.DB

// Init 定义一个初始化数据库的函数
func Init() {
	dsn := "shengyi:123456@tsy@tcp(rm-2vc34w5spf5nm2992eo.mysql.cn-chengdu.rds.aliyuncs.com)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1000)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(2000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}

// 定义一个全局对象db
var db *sql.DB

func initDb(env string) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("error:", err)
		}
	}()
	var err error
	dsn := "shengyi:123456@tsy@tcp(rm-2vc34w5spf5nm2992eo.mysql.cn-chengdu.rds.aliyuncs.com)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Error("数据库连接出现问题，请检查dsn格式")
		panic(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Error("数据库连接出现问题，请检查数据库连接")
		panic(err)
	}
	log.Info("初始化数据库成功")
}
