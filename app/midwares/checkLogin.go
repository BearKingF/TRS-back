package midwares

import (
	"TRS/app/services/sessionService"
	"github.com/gin-gonic/gin"
)

// 判断用户是否存在

func CheckLogin(c *gin.Context) bool {
	//_, err := userService.GetUserByID(id)
	isLogin := sessionService.CheckUserSession(c)
	if !isLogin {
		return false
	}
	return true
}
