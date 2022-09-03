package auth

import (
	"github.com/gin-gonic/gin"
	v1 "go-api-practice/app/http/controllers/api/v1"
	"go-api-practice/app/models/user"
	"go-api-practice/app/requests"
	"go-api-practice/pkg/response"
)

type SignupController struct {
	v1.BaseController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context) {
	request := requests.SignupUsingPhoneRequest{}

	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	_user := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}

	_user.Create()

	if _user.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": _user,
		})
	} else {
		response.Abort500(c, "创建用户失败")
	}
}
