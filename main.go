package main

import (
	"github.com/gin-gonic/gin"
	"westonline/models"
	"westonline/routers"
)

func main() {
	//创建一个默认路由
	r := gin.Default()
	//连接数据库
	models.MysqlInit()
	//访问接口接收数据
	routers.Routersinit(r)
	r.Run()
}
