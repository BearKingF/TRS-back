package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func GetTeamMember(teamID uint) ([]models.User, error) {
	result := database.DB.Where(&models.Team{TeamID: teamID}).First(&models.User{})
	if result.Error != nil { //团队为空
		return nil, result.Error
	}
	var teamMemberList []models.User //切片
	result = database.DB.Where(&models.Team{TeamID: teamID}).Find(&teamMemberList)
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
