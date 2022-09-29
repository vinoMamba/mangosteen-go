package main

import (
	"github.com/mangosteen-go/dao"
	"github.com/mangosteen-go/model"
	"github.com/mangosteen-go/router"
)

func main() {
	//init postgres
	err := dao.InitPostgres()
	if err != nil {
		panic(err)
	}

	defer dao.ClosePostgres()

	dao.DB.AutoMigrate(&model.ValidationCode{})
	r := router.SetupRouter()

	r.Run("localhost:3000")
}
