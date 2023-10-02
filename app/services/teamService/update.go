package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

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

func UpdateTeamInfo(team models.Team) error {
	result := database.DB.Updates(&team)
	return result.Error
}

func UpdateStatus(teamID uint) error { //更改Status信息
	team, _ := GetTeamByTeamID(teamID)
	if team.Status == 2 {
		team.Status = 1
	} else {
		team.Status = 2
	}
	result := database.DB.Updates(&team)
	return result.Error
}
