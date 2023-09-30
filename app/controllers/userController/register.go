package userController

import (
	"TRS/app/models"
	"TRS/app/services/userService"
	"TRS/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	//ID       uint   `json:"id"`
	Username        string `json:"username" binding:"required"`
	Sex             string `json:"sex" binding:"required"`
	PhoneNum        string `json:"phone_num" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Major           string `json:"major" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
	Type            uint   `json:"type" binding:"required"`
}

/* 注册(POST) */

func Register(c *gin.Context) {

	//1. 接收参数
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//2. 判断手机号是否已被注册
	err = userService.CheckUserExistByPhoneNum(data.PhoneNum)
	if err == nil { //说明该手机号已存在
		utils.JsonErrorResponse(c, 200504, "手机号已注册")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//3. 判断邮箱是否已被注册
	err = userService.CheckUserExistByEmail(data.Email)
	if err == nil { //说明该邮箱已存在
		utils.JsonErrorResponse(c, 200505, "邮箱已注册")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//4. 判断两次输入的密码是否一致
	flag := userService.ComparePwd(data.Password, data.ConfirmPassword)
	if !flag {
		utils.JsonErrorResponse(c, 200506, "密码不一致")
		return
	}

	//5. 判断账户类型是否合法
	if data.Type != 1 && data.Type != 2 {
		utils.JsonErrorResponse(c, 200511, "类型不合法")
		return
	}

	//5. 注册用户
	err = userService.Register(models.User{
		Username:  data.Username,
		Sex:       data.Sex,
		PhoneNum:  data.PhoneNum,
		Email:     data.Email,
		Major:     data.Major,
		Password:  data.Password,
		Type:      data.Type,
		IsCaptain: false,
		TeamID:    -1,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//6. 注册成功
	utils.JsonSuccessResponse(c, nil)
}
