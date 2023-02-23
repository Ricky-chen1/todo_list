package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"westonline/api"
	"westonline/middleware"
)

func Routersinit(r *gin.Engine) {
	store := cookie.NewStore([]byte("secret11111"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("/api/v1")
	{
		//用户注册
		v1.POST("/register", api.RegisterUser)
		//用户登录
		v1.POST("/login", api.LoginUser)
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			//增加一条待办事务
			authed.POST("/task", api.CreateTask)
			//查询待办/未办/全部事务
			authed.GET("/tasks", api.ListAllTask)
			//关键字搜索
			authed.GET("/task", api.ReadTask)
			//将一条事务设为代办/已完成
			authed.PUT("/task/:id", api.UpdateTask)
			//将所有事务设为待办/已完成
			authed.PUT("/tasks", api.UpdateAllTask)
			//删除一条待办/已完成事务
			authed.DELETE("/task/:id", api.DeleteTask)
			//删除所有待办/已完成事务
			//使用post请求来使用请求体传参
			authed.POST("/tasksdelete", api.DeleteAllTask)
		}
	}
}
