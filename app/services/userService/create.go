package userService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func CreateUser(user models.User) error { //注册用户（往数据库中添加新的用户）
	result := database.DB.Create(&user)
	return result.Error //创建过程中可能会出现错误
}
