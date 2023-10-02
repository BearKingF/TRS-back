package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func CheckTeamExistByTeamID(teamID uint) error {
	result := database.DB.Where(&models.Team{TeamID: teamID}).First(&models.Team{})
	return result.Error
}

func CheckTeamExistByTeamName(teamName string) error {
	result := database.DB.Where(&models.Team{TeamName: teamName}).First(&models.Team{})
	return result.Error
}

func CheckPwd(pwd1 string, pwd2 string) bool {
	return pwd1 == pwd2
}

func CheckStatus(teamID uint) bool { //判断团队是否已提交报名申请（在团队存在的前提下）
	team, _ := GetTeamByTeamID(teamID)
	if team.Status == 1 {
		return true
	}
	return false
}
