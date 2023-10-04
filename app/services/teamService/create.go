package teamService

import (
	"TRS/app/models"
	"TRS/config/database"
)

func CreateTeam(team models.Team) (uint, error) { //返回创建团队的团队编号
	result := database.DB.Create(&team) //通过数据的指针来创建
	return team.TeamID, result.Error    //team.TeamID: 返回主键
}
