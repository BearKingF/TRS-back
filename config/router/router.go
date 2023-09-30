package router

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"

	api := r.Group(pre) //路由组
	{
		userRouterInit(api)
		clientRouterInit(api)
		adminRouterInit(api)
	}

}
