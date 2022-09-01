package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"go-api-practice/pkg/captcha"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
		"phone":          []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{
		"captcha_id": []string{
			"required:请输入验证码",
		},
		"captcha_answer": []string{
			"required:请输入验证码",
			"digits:验证码长度必须是6位",
		},
		"phone": []string{
			"required:请输入手机号",
			"digits:手机号长度必须是11位",
		},
	}
	errs := validate(data, rules, messages)

	_data := data.(*VerifyCodePhoneRequest)
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "验证码错误")
	}
	return errs

}
