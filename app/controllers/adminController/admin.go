package adminController

import (
	"TRS/app/midwares"
	"TRS/app/services/sessionService"
	"TRS/app/services/teamService"
	"TRS/app/utils"
	"github.com/gin-gonic/gin"
)

/* 获取所有团队信息(GET) */

//type GetAllTeamInfoData struct {
//	UserID uint `form:"user_id" binding:"required"` //form: 对应c.ShouldBindQuery(&data)
//}

func GetAllTeamInfo(c *gin.Context) {
	//var data GetAllTeamInfoData
	//err := c.ShouldBindQuery(&data)
	//if err != nil {
	//	utils.JsonErrorResponse(c, 200501, "参数错误")
	//	return
	//}

	user, err := sessionService.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "未登录")
	}

	//判断是否为管理员账户
	flag := midwares.CheckAdmin(user.ID)
	if !flag {
		utils.JsonErrorResponse(c, 200520, "非管理员账户")
		return
	}

	//可以获得所有团队的信息
	committedTeamList, count1, err := teamService.GetAllIsCommittedTeam(1)
	if err != nil { //注意：Find 查询不到记录是不报错的，所以这样写
		utils.JsonInternalServerErrorResponse(c)
	}

	uncommittedTeamList, count2, err := teamService.GetAllIsCommittedTeam(2)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
	}

	count, err := teamService.GetAllTeamCount()
	if err != nil { //注意：Count 查询不到记录也是不报错的
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"all_team_count":         count,
		"committed_team_list":    committedTeamList,
		"committed_team_count":   count1,
		"uncommitted_team_list":  uncommittedTeamList,
		"uncommitted_team_count": count2,
	})
}
