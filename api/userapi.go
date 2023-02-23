package api

import (
	"github.com/gin-gonic/gin"
	"westonline/service"
)

func RegisterUser(c *gin.Context) {
	//创建服务对象接收数据
	var userRegister service.Userservice
	//操作数据并返回前端json数据
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(400, err)
	} else {
		res := userRegister.Register()
		c.JSON(200, res)
	}
}

func LoginUser(c *gin.Context) {
	//创建服务对象接收数据
	var userLogin service.Userservice
	//操作数据并返回前端json数据
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(400, err)
	} else {
		res := userLogin.Login()
		c.JSON(200, res)
	}
}
