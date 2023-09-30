package userService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func CheckUserExistByPhoneNum(phoneNum string) error {
	result := database.DB.Where(&models.User{PhoneNum: phoneNum}).First(&models.User{})
	return result.Error
}

func CheckUserExistByEmail(email string) error {
	result := database.DB.Where(&models.User{Email: email}).First(&models.User{})
	return result.Error
}

func CheckUserExistByAccount(account string) error {
	result := database.DB.Where(&models.User{Email: account}).Or(&models.User{PhoneNum: account}).First(&models.User{})
	//result := database.DB.Where("email = ? or phone_num = ?", account, account).First(&models.User{})
	return result.Error
}

func CheckPwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}
