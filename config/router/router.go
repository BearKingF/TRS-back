package router

import (
	"TRS/app/controllers/userController"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre) //路由组
	{
		api.POST("/login", userController.Login) //登录

		api.POST("/register", userController.Register) //注册
	}

}
