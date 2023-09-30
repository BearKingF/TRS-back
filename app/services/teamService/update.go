package teamService

import "TRS/config/database"

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
