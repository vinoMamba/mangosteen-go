package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/mangosteen-go/bootstrap"
)

func main() {
	r := gin.New()
	bootstrap.SetupRouter(r)

	r.Run(":3000")
}
