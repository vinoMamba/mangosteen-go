package database

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dbConfig gorm.Dialector) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
