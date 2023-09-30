package midwares

import "TRS/app/services/userService"

// 判断用户是否存在

func CheckLogin(id uint) bool {
	_, err := userService.GetUserByID(id)
	if err != nil {
		return false
	}
	return true
}
