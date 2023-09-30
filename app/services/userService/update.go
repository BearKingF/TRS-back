package userService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func UpdateUserInfo(user models.User) error {
	result := database.DB.Updates(&user)
	return result.Error
}

func UpdateTeamID(id uint, teamID uint) error {
	user, _ := GetUserByID(id)
	user.TeamID = int(teamID)
	result := database.DB.Updates(&user)
	return result.Error
}

func UpdateCaptainFlag(id uint) error { //更改isCaptain信息
	user, _ := GetUserByID(id)
	if user.IsCaptain == false {
		user.IsCaptain = true
	} else {
		user.IsCaptain = false
	}
	result := database.DB.Updates(&user)
	return result.Error
}
