package router

import (
	"TRS/app/controllers/userController"
	"github.com/gin-gonic/gin"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/login", userController.Login)       //登录
		user.POST("/register", userController.Register) //注册
	}
}
