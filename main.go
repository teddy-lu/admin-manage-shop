package main

import (
	"admin-manage-shop/bootstrap"
	btsConfig "admin-manage-shop/config"
	"admin-manage-shop/pkg/config"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	bootstrap.SetupLogger()
	gin.SetMode(gin.ReleaseMode)

	bootstrap.SetupDB()

	bootstrap.SetupRedis()

	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
