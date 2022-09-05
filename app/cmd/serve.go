package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-api-practice/bootstrap"
	"go-api-practice/pkg/config"
	"go-api-practice/pkg/console"
	"go-api-practice/pkg/logger"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD", "serve", err.Error())
		console.Exit(err.Error())
	}
}
