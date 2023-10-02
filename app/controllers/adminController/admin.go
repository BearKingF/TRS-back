package adminController

import (
	"TRS/app/midwares"
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* 获取所有团队信息(GET) */

type GetAllTeamInfoData struct {
	UserID uint `form:"user_id" binding:"required"` //form: 对应c.ShouldBindQuery(&data)
}

func GetAllTeamInfo(c *gin.Context) {
	var data GetAllTeamInfoData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	//判断是否为管理员账户
	flag := midwares.CheckAdmin(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200520, "非管理员账户")
		return
	}

	//可以获得所有团队的信息
	//var committedTeamList []models.Team
	committedTeamList, count1, err := teamService.GetAllIsCommittedTeam(1)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			committedTeamList = make([]models.Team, 0) //空切片[] 不同于nil!!!
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	uncommittedTeamList, count2, err := teamService.GetAllIsCommittedTeam(2)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			uncommittedTeamList = make([]models.Team, 0) //空切片[]
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	count, err := teamService.GetAllTeamCount()
	if err != nil && err != gorm.ErrRecordNotFound {
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
