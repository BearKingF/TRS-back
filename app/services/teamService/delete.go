package teamService

import (
	"TRS/config/database"
)

func Delete(teamID uint) error {
	team, _ := GetTeamByTeamID(teamID)
	result := database.DB.Delete(&team) //根据team对象的主键teamID删除
	//注意：要删除的数据不存在不会报错
	return result.Error
}
