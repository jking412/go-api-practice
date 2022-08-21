package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-api-practice/routes"
	"net/http"
)

func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleware(router)
	routes.RegisterAPIRoutes(router)
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error_code":    404,
			"error_message": "页面不存在",
		})
	})
}
