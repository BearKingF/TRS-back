package midwares

import "TRS/app/services/userService"

// 判断是否为管理员用户

func CheckAdmin(id uint) bool {
	user, err := userService.GetUserByID(id)
	if err == nil && user.Type == 2 {
		return true
	}
	return false
}
