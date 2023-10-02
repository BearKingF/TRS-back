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

func GetAllIsCommittedTeam(R uint) ([]models.Team, int64, error) { //R取1或2
	//确定有无记录存在
	result := database.DB.Model(&models.Team{}).Where(&models.Team{Status: R}).First(&models.Team{})
	if result.Error != nil {
		return nil, 0, result.Error
	}
	var count int64
	var teamList []models.Team
	result = database.DB.Model(&models.Team{}).Where(&models.Team{Status: R}).Count(&count)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	result = database.DB.Model(&models.Team{}).Where(&models.Team{Status: R}).Find(&teamList) //可以这样写吗（怎么知道是team的那张表）
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return teamList, count, nil
}

func GetAllTeamCount() (int64, error) { //获取记录总数
	var count int64 //要求参数int64
	result := database.DB.Model(&models.Team{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
