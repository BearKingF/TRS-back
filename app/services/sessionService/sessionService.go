package sessionService

import (
	"TRS/app/models"
	"TRS/app/services/userService"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ClearUserSession(c *gin.Context) { //清除用户session
	webSession := sessions.Default(c)
	webSession.Delete("id")
	webSession.Save()
	return
}

func SetUserSession(c *gin.Context, user *models.User) error {
	webSession := sessions.Default(c)
	webSession.Options(sessions.Options{MaxAge: 3600 * 24 * 7, Path: "/api"}) //Path: "/api"
	//对于浏览器来说，是要指定Path的
	//同源策略问题：解决———跨域资源共享（域，即Path前缀要一样）
	webSession.Set("id", user.ID)
	return webSession.Save()
}

func GetUserSession(c *gin.Context) (*models.User, error) {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return nil, errors.New("") //errors.New("") 创建一个错误对象，返回error类型
	}
	user, _ := userService.GetUserByID(id.(uint)) // id.(uint) 上面返回的id是 interface{} 类型的

	if user == nil {
		ClearUserSession(c)
		return nil, errors.New("")
	}
	return user, nil
}

func CheckUserSession(c *gin.Context) bool {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return false
	}
	return true
}
