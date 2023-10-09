package midwares

import (
	"TRS/app/services/sessionService"
	"TRS/app/utils"
	"github.com/gin-gonic/gin"
)

// 中间件：判断用户是否登录(预处理)

func CheckLogin(c *gin.Context) {

	isLogin := sessionService.CheckUserSession(c)
	if !isLogin {
		utils.JsonErrorResponse(c, 200507, "未登录")
		c.Abort()
	}
	c.Next()
}
