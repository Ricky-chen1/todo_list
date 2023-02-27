package service

import (
	"westonline/models"
	"westonline/pkg/serializer"
	"westonline/pkg/utils"
)

type Userservice struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
}

// 用户注册
func (service *Userservice) Register() serializer.Response {
	var user models.User
	var count int64 = 0
	models.DB.Where("username=?", service.Username).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户已存在",
		}
	}
	user.Username = service.Username
	user.Email = service.Email
	user.SetPassword(service.Password)
	//创建用户
	if err := models.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "注册用户失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
		Data:   service,
	}
}

// 用户登录
func (service *Userservice) Login() serializer.Response {
	user := models.User{}
	//查找用户
	var count int64 = 0
	models.DB.Where("username=?", service.Username).First(&user).Count(&count)
	if count == 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户未注册",
		}
	}
	//校验密码
	if err := user.CheckPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//用户登录时签发token
	claims := utils.Claims{
		Id:       user.ID,
		Username: user.Username,
	}
	token, err := utils.GenerateToken(claims, "golang")
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户登录成功",
		Data: serializer.TokenData{
			Username: service.Username,
			Email:    user.Email,
			Token:    token,
		},
	}
}
