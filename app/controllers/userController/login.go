package userController

import (
	"TRS/app/models"
	"TRS/app/services/userService"
	"TRS/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginData struct {
	Account  string `json:"account" binding:"required"` //使用手机号或邮箱登录（唯一标识）
	Password string `json:"password" binding:"required"`
}

/* 登录(POST) */

func Login(c *gin.Context) {

	//1. 接收参数
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//2. 判断用户是否存在
	err = userService.CheckUserExistByAccount(data.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	//3. 获取用户信息
	var user *models.User
	user, err = userService.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//4. 判断密码是否正确
	flag := userService.CheckPwd(data.Password, user.Password)
	if !flag {
		utils.JsonErrorResponse(c, 200503, "密码错误")
		return
	}

	//5. 登录成功，返回用户信息
	utils.JsonSuccessResponse(c, user)
}
