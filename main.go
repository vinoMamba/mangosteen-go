package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/vinoMamba/mangosteen-go/bootstrap"
	pkgConfig "github.com/vinoMamba/mangosteen-go/config"
	"github.com/vinoMamba/mangosteen-go/pkg/config"
)

func init() {
	pkgConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	bootstrap.SetupDB()

	r := gin.New()
	bootstrap.SetupRouter(r)

	r.Run(":" + config.Get("route.port"))
}
