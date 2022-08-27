package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type ValidateFunc func(interface{}, *gin.Context) map[string][]string

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	ops := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	return govalidator.New(ops).ValidateStruct()
}

func Validate(c *gin.Context, obj interface{}, handler ValidateFunc) bool {
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "数据格式错误",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "上传的数据格式出错",
			"error":   errs,
		})
		return false
	}

	return true
}
