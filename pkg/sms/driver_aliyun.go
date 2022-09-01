package sms

import (
	"encoding/json"
	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
	"go-api-practice/pkg/logger"
)

type Aliyun struct{}

func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	smsClient := aliyunsmsclient.New("http://dysmsapi.aliyuncs.com/")

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("Aliyun sms error", "json.Marshal error", err.Error())
		return false
	}

	logger.DebugJSON("Aliyun sms", "config", config)

	result, err := smsClient.Execute(
		config["access_key_id"],
		config["access_key_secret"],
		phone,
		config["sign_name"],
		message.Template,
		string(templateParam),
	)

	logger.DebugJSON("Aliyun sms", "request", smsClient.Request)
	logger.DebugJSON("Aliyun sms", "result", result)

	if err != nil {
		logger.ErrorString("Aliyun sms error", "Execute error", err.Error())
		return false
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		logger.ErrorString("Aliyun sms error", "json.Marshal error", err.Error())
		return false
	}

	if result.IsSuccessful() {
		logger.DebugString("Aliyun sms", "success", "")
		return true
	} else {
		logger.ErrorString("Aliyun sms", "error", string(resultJSON))
		return false
	}

}
