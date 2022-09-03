// Package user 存放用户 Model 相关逻辑
package user

import (
	"go-api-practice/app/models"
	"go-api-practice/pkg/database"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (user *User) Create() {
	database.DB.Create(&user)
}
