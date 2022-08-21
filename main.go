package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api-practice/bootstrap"
	btsconfig "go-api-practice/config"
	"go-api-practice/pkg/config"
)

func init() {
	btsconfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "")
	flag.Parse()
	config.InitConfig(env)

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err)
	}
}
