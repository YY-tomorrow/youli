package store

import (
	"anest.top/youli/api/proto/config"
	"anest.top/youli/internal/data/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewMysql(config *config.Config) (*gorm.DB, error) {
	mysqlConf := config.MysqlConf
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConf.Dsn,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(
		&models.Auth{}); err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(int(mysqlConf.MaxOpenConn))
	sqlDB.SetMaxIdleConns(int(mysqlConf.MaxIdleConn))
	sqlDB.SetConnMaxLifetime(
		time.Duration(mysqlConf.MaxLifeMinute) * time.Minute)

	return db, err

}
