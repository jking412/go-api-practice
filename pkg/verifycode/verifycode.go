package verifycode

import (
	"go-api-practice/helpers"
	"go-api-practice/pkg/app"
	"go-api-practice/pkg/config"
	"go-api-practice/pkg/logger"
	"go-api-practice/pkg/redis"
	"go-api-practice/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				keyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

func (vc *VerifyCode) SendSMS(phone string) bool {
	code := vc.generateVerifyCode(phone)

	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}

	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data: map[string]string{
			"code": code,
		},
	})
}

func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})
	if !app.IsProduction() &&
		(strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
			strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}
	return vc.Store.Verify(key, answer, false)
}

func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	vc.Store.Set(key, code)
	return code
}
