package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "go-api-practice/app/http/controllers/api/v1"
	"go-api-practice/app/models/user"
	"go-api-practice/app/requests"
	"net/http"
)

type SignupController struct {
	v1.BaseController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		fmt.Println(err.Error())

		return
	}

	errs := requests.ValidatePhoneExist(&request, c)

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
