package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func Delete(teamID uint) error {
	team, _ := GetTeamByTeamID(teamID)
	result := database.DB.Where(models.Team{}).Delete(&team)
	return result.Error
}
