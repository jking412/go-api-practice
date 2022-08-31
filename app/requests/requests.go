package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"go-api-practice/pkg/response"
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
		response.BadRequest(c, err)
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}

	return true
}
