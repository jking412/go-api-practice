package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-practice/bootstrap"
	btsconfig "go-api-practice/config"
	"go-api-practice/pkg/config"
	"go-api-practice/pkg/logger"
)

func init() {
	btsconfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "")
	flag.Parse()
	config.InitConfig(env)

	bootstrap.SetupLogger()

	logger.Dump("test")

	router := gin.New()

	bootstrap.SetupDB()

	bootstrap.SetupRoute(router)

	gin.SetMode(gin.ReleaseMode)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err)
	}
}
