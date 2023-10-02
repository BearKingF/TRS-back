package midwares

import "TRS/app/services/userService"

func CheckCaptain(id uint) bool {
	user, err := userService.GetUserByID(id)
	if err == nil && user.IsCaptain == 1 {
		return true
	}
	return false
}
