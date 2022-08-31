package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-api-practice/app/http/controllers/api/v1"
	"go-api-practice/pkg/captcha"
	"go-api-practice/pkg/logger"
	"net/http"
)

type VerifyCodeController struct {
	v1.BaseController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
