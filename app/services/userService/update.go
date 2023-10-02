package userService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func UpdateUserInfo(user models.User) error {
	result := database.DB.Updates(&user)
	return result.Error
}

func UpdateTeamID(id uint, teamID int) error {
	user, _ := GetUserByID(id)
	user.TeamID = teamID
	result := database.DB.Updates(&user)
	return result.Error
}

func UpdateCaptainFlag(id uint) error { //更改isCaptain信息
	user, _ := GetUserByID(id)
	if user.IsCaptain == 2 {
		user.IsCaptain = 1
	} else {
		user.IsCaptain = 2
	}
	result := database.DB.Updates(&user)
	return result.Error
}

func UpdateAllTeamID(teamID uint) error { //批量更新
	//db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	//不知道对不对？
	result := database.DB.Where(&models.User{TeamID: int(teamID)}).Updates(models.User{TeamID: -1})
	return result.Error
}
