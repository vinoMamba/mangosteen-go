package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitPostgres() (err error) {
	dsn := "host=localhost user=mangosteen password=123456 dbname=mangosteen_dev port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.Error
}

func ClosePostgres() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}
