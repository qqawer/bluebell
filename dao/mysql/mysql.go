package mysql

import (
	"WebApp/settings"
	
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *settings.Config) (err error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	cfg.Mysql.User,
	// 	cfg.Mysql.Password,
	// 	cfg.Mysql.Host,
	// 	cfg.Mysql.Port,
	// 	cfg.Mysql.Name)
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	zap.L().Error("Failed to initialize database", zap.Error(err))
	// }
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(cfg.Mysql.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.Mysql.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil

}
func Close() {
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
}
