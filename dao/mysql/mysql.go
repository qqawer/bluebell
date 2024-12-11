package mysql

import (
	"WebApp/global"
	"fmt"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init()(err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.AppConfig.Mysql.User,
		global.AppConfig.Mysql.Password,
		global.AppConfig.Mysql.Host,
		global.AppConfig.Mysql.Port,
		global.AppConfig.Mysql.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	zap.L().Error("Failed to initialize database", zap.Error(err))
	// }

	sqlDB, err := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(global.AppConfig.Mysql.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(global.AppConfig.Mysql.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.Db=db
	return

}

