package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func CreateTeam(team models.Team) (uint, error) { //返回创建团队的团队编号
	result := database.DB.Create(&team) //通过数据的指针来创建
	return team.TeamID, result.Error
}

func GetTeamMember(teamID uint) ([]models.User, error) {
	result := database.DB.Where("team_id = ?", teamID).First(&models.User{})
	if result.Error != nil { //团队为空
		return nil, result.Error
	}
	var teamMemberList []models.User //切片
	result = database.DB.Where("team_id = ?", teamID).Find(&teamMemberList)
	if result.Error != nil {
		return nil, result.Error
	}
	return teamMemberList, nil
}

func GetTeamByTeamID(teamID uint) (*models.Team, error) {
	var team models.Team
	result := database.DB.Where(&models.Team{TeamID: teamID}).First(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}

func CheckTeamExistByTeamID(teamID uint) error {
	result := database.DB.Where("team_id = ?", teamID).First(&models.Team{})
	return result.Error
}

func UpdateTotal(teamID uint) error { //团队人数 +1
	team, _ := GetTeamByTeamID(teamID)
	team.Total++
	result := database.DB.Updates(&team)
	return result.Error
}

func UpdateTotal2(teamID uint) error { //团队人数 -1
	team, _ := GetTeamByTeamID(teamID)
	team.Total--
	result := database.DB.Updates(&team)
	return result.Error
}

func CheckTeamExistByTeamName(teamName string) error {
	result := database.DB.Where("team_name = ?", teamName).First(&models.Team{})
	return result.Error
}
