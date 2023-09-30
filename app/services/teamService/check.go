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
