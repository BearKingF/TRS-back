package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func GetTeamMember(teamID uint) ([]models.User, error) {

	// 根据主键检索
	result := database.DB.First(&models.User{}, teamID) // TeamID 为 models.User 的主键
	if result.Error != nil {                            //团队为空
		return nil, result.Error
	}
	var teamMemberList []models.User //切片
	result = database.DB.Find(&teamMemberList, teamID)
	if result.Error != nil {
		return nil, result.Error
	}
	return teamMemberList, nil
}

func GetTeamByTeamID(teamID uint) (*models.Team, error) {
	var team models.Team

	//通过主键查询指定的某条记录，且主键为整型
	result := database.DB.First(&team, teamID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}

type TeamInfo struct {
	TeamID    uint   `json:"team_id"`
	TeamName  string `json:"team_name"`
	CaptainID uint   `json:"captain_id"`
	Total     uint   `json:"total"`
}

func GetAllIsCommittedTeam(R uint) ([]TeamInfo, int64, error) { //R取1或2
	//无需区分有无记录存在的情况

	var teamList []TeamInfo

	//find 方法在查询不到记录时不会报错！（此处利用这个特点）
	result := database.DB.Model(&models.Team{}).Where(&models.Team{Status: R}).Select("team_id", "team_name", "captain_id", "total").Find(&teamList)
	//Debug(): SELECT `team_id`,`team_name`,`captain_id`,`total` FROM `teams` WHERE `teams`.`status` = R

	//或 result := database.DB.Model(&models.Team{}).Where(&models.Team{Status: R}).Omit("password", "status").Find(&teamList)

	//不加Select其实也可以……
	//result = database.DB.Model(&models.Team{}).Where(&models.Team{Status: R}).Debug().Find(&teamList)
	//SELECT `teams`.`team_id`,`teams`.`team_name`,`teams`.`captain_id`,`teams`.`total` FROM `teams` WHERE `teams`.`status` = R

	return teamList, result.RowsAffected, result.Error // result.RowsAffected 为记录条数
}

func GetAllTeamCount() (int64, error) { //获取记录总数
	var count int64                                           //要求参数int64
	result := database.DB.Model(&models.Team{}).Count(&count) //此处必须要写 Model(&models.Team{})
	//fmt.Println(result.Error)
	// Count() 在查询不到记录时也不会报错
	return count, result.Error
}
