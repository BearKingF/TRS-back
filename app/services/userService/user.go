package userService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func CheckUserExistByPhoneNum(phoneNum string) error {
	result := database.DB.Where("phone_num = ?", phoneNum).First(&models.User{})
	return result.Error
}

func CheckUserExistByEmail(email string) error {
	result := database.DB.Where("email = ?", email).First(&models.User{})
	return result.Error
}

func CheckUserExistByAccount(account string) error {
	result := database.DB.Where("email = ? or phone_num = ?", account, account).First(&models.User{})
	return result.Error
}

func GetUserByAccount(account string) (*models.User, error) {
	var user models.User // 创建一个User的实例
	result := database.DB.Where("email = ? or phone_num = ?", account, account).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func Register(user models.User) error { //注册用户（往数据库中添加新的用户）
	result := database.DB.Create(&user)
	return result.Error //创建过程中可能会出现错误
}
