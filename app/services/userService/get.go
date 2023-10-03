package userService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func GetUserByAccount(account string) (*models.User, error) {
	var user models.User // 创建一个User的实例
	result := database.DB.Where(&models.User{Email: account}).Or(&models.User{PhoneNum: account}).First(&user)
	//result := database.DB.Where("email = ? or phone_num = ?", account, account).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	//result := database.DB.Where(&models.User{ID: id}).First(&user)
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
