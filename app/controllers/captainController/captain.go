package captainController

import (
	"TRS/app/midwares"
	"TRS/app/models"
	"TRS/app/services/teamService"
	"TRS/app/services/userService"
	"TRS/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* 编辑团队信息(PUT)（可修改团队名称和密码）*/

type UpdateTeamInfoData struct {
	UserID uint `json:"user_id" binding:"required"`
	//TeamID          uint   `json:"team_id" binding:"required"`
	TeamName        string `json:"team_name"`    //团队名称不可重复
	OldPassword     string `json:"old_password"` //如果要修改密码，需要先输入旧密码
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

func UpdateTeamInfo(c *gin.Context) {
	var data UpdateTeamInfoData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	flag := midwares.CheckLogin(c)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "未登录")
		return
	}
	//判断是否为队长账户
	flag = midwares.CheckCaptain(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200514, "非队长账户")
		return
	}
	user, _ := userService.GetUserByID(data.UserID)

	//判断团队名是否重复
	if data.TeamName != "" {
		err = teamService.CheckTeamExistByTeamName(data.TeamName)
		if err == nil {
			utils.JsonErrorResponse(c, 200512, "团队名已注册")
			return
		}
	}
	team, _ := teamService.GetTeamByTeamID(uint(user.TeamID))

	//判断若要修改密码，输入的原密码是否正确及两次输入的新密码是否一致
	if data.OldPassword != "" || data.NewPassword != "" || data.ConfirmPassword != "" {
		if data.OldPassword != "" && data.NewPassword != "" && data.ConfirmPassword != "" {
			flag = teamService.CheckPwd(data.OldPassword, team.Password) && teamService.CheckPwd(data.NewPassword, data.ConfirmPassword)
			if !flag {
				utils.JsonErrorResponse(c, 200506, "密码不一致")
				return
			}
		} else {
			utils.JsonErrorResponse(c, 200506, "密码不一致")
			return
		}
	}

	//更新团队信息（为零值的属性不会更新 --> 保证了信息正确性）
	err = teamService.UpdateTeamInfo(models.Team{
		TeamID:   uint(user.TeamID),
		TeamName: data.TeamName,
		Password: data.NewPassword,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

/* 提交报名(PUT) */

type ApplyData struct {
	UserID uint `json:"user_id" binding:"required"`
}

func Apply(c *gin.Context) {
	var data ApplyData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
	}

	flag := midwares.CheckLogin(c)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "未登录")
		return
	}
	//判断是否为队长账户
	flag = midwares.CheckCaptain(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200514, "非队长账户")
		return
	}

	user, _ := userService.GetUserByID(data.UserID)
	team, _ := teamService.GetTeamByTeamID(uint(user.TeamID))

	//判断团队人数是否符合要求
	if team.Total < 4 {
		utils.JsonErrorResponse(c, 200516, "团队人数不足")
		return
	}

	//判断团队是否已提交过申请
	if team.Status == 1 {
		utils.JsonErrorResponse(c, 200515, "团队已提交报名")
		return
	}

	//可以提交报名
	err = teamService.UpdateStatus(uint(user.TeamID)) //将团队状态改为已提交
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)

}

/* 撤销报名(DEL) */

type DisApplyData struct {
	UserID uint `json:"user_id" binding:"required"`
}

func DisApply(c *gin.Context) {
	var data DisApplyData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
	}
	flag := midwares.CheckLogin(c)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "未登录")
		return
	}
	//判断是否为队长账户
	flag = midwares.CheckCaptain(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200514, "非队长账户")
		return
	}

	user, _ := userService.GetUserByID(data.UserID)

	//判断团队是否已提交过报名
	if teamService.CheckStatus(uint(user.TeamID)) == false {
		utils.JsonErrorResponse(c, 200519, "团队未提交报名")
		return
	}

	//可以撤销报名
	err = teamService.UpdateStatus(uint(user.TeamID)) //将团队状态改为未提交
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)

}

/* 移除团队成员(DEL)（不可删除自己）（必须要在团队未提交报名状态下）*/

type DeleteTeamMemberData struct {
	UserID   uint `json:"user_id" binding:"required"`
	MemberID uint `json:"member_id" binding:"required"` //要删除的成员id

}

func DeleteTeamMember(c *gin.Context) {
	var data DeleteTeamMemberData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	flag := midwares.CheckLogin(c)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "未登录")
		return
	}
	//判断是否为队长账户
	flag = midwares.CheckCaptain(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200514, "非队长账户")
		return
	}

	user, _ := userService.GetUserByID(data.UserID)
	//判断团队是否为已提交报名状态
	flag = teamService.CheckStatus(uint(user.TeamID))
	if flag {
		utils.JsonErrorResponse(c, 200515, "团队已提交报名")
		return
	}
	//判断要删除成员是否在团队中
	member, err := userService.GetUserByID(data.MemberID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	if member.TeamID != user.TeamID {
		utils.JsonErrorResponse(c, 200517, "用户不在团队中")
		return
	}
	//不可删除自己（队长只有转移了队长职位，才可以退出团队）
	if data.MemberID == data.UserID {
		utils.JsonErrorResponse(c, 200518, "不可移除队长")
		return
	}
	//可以移除成员（其实使用的是Update函数）
	err = teamService.UpdateTotal2(uint(user.TeamID)) //Total--
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
	}
	err = userService.UpdateTeamID(data.MemberID, -1)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

/* 添加团队成员(PUT) */

type AddTeamMemberData struct {
	UserID   uint `json:"user_id" binding:"required"`
	MemberID uint `json:"member_id" binding:"required"` //要添加的成员id
}

func AddTeamMember(c *gin.Context) {
	var data AddTeamMemberData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	flag := midwares.CheckLogin(c)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "未登录")
		return
	}
	//判断是否为队长账户
	flag = midwares.CheckCaptain(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200514, "非队长账户")
		return
	}

	user, _ := userService.GetUserByID(data.UserID)
	//判断团队是否为已提交报名状态
	flag = teamService.CheckStatus(uint(user.TeamID))
	if flag {
		utils.JsonErrorResponse(c, 200515, "团队已提交报名")
		return
	}

	team, _ := teamService.GetTeamByTeamID(uint(user.TeamID))
	//团队已满
	if team.Total == 6 {
		utils.JsonErrorResponse(c, 200509, "团队已满")
		return
	}
	//要添加成员用户不存在
	member, err := userService.GetUserByID(data.MemberID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	//要添加成员已加入团队
	if member.TeamID != -1 {
		utils.JsonErrorResponse(c, 200510, "已加入团队")
		return
	}

	//可以添加成员
	err = teamService.UpdateTotal(uint(user.TeamID))
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = userService.UpdateTeamID(data.MemberID, user.TeamID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

/* 解散团队(DEL)（解散团队要输入团队密码）（要在团队未提交报名的状态下）*/

type DeleteTeamData struct {
	UserID   uint   `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func DeleteTeam(c *gin.Context) {
	var data DeleteTeamData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	flag := midwares.CheckLogin(c)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "未登录")
		return
	}
	//判断是否为队长账户
	flag = midwares.CheckCaptain(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200514, "非队长账户")
		return
	}
	user, _ := userService.GetUserByID(data.UserID)
	teamID := uint(user.TeamID)
	//判断团队是否为已提交报名状态
	flag = teamService.CheckStatus(teamID)
	if flag {
		utils.JsonErrorResponse(c, 200515, "团队已提交报名")
		return
	}

	//判断密码是否正确
	team, _ := teamService.GetTeamByTeamID(teamID)
	if teamService.CheckPwd(data.Password, team.Password) == false {
		utils.JsonErrorResponse(c, 200506, "密码不一致")
		return
	}

	//可以解散团队
	err = teamService.Delete(teamID) //删除团队在数据库中的数据
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = userService.UpdateAllTeamID(teamID) //将所有队员的TeamID置成-1
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = userService.UpdateCaptainFlag(data.UserID) //更改队长的isCaptain状态
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//返回成功响应
	utils.JsonSuccessResponse(c, nil)
}

/* 转移队长职位(PUT) */

type TransferCaptainData struct {
	UserID   uint `json:"user_id" binding:"required"`
	MemberID uint `json:"member_id" binding:"required"`
}

func TransferCaptain(c *gin.Context) {
	var data TransferCaptainData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
	}

	flag := midwares.CheckLogin(c)
	if !flag {
		utils.JsonErrorResponse(c, 200507, "未登录")
		return
	}
	//判断是否为队长账户
	flag = midwares.CheckCaptain(data.UserID)
	if !flag {
		utils.JsonErrorResponse(c, 200514, "非队长账户")
		return
	}
	//要添加成员用户不存在
	member, err := userService.GetUserByID(data.MemberID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200502, "用户不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	user, _ := userService.GetUserByID(data.UserID)

	//判断团队是否为已提交报名状态
	flag = teamService.CheckStatus(uint(user.TeamID))
	if flag {
		utils.JsonErrorResponse(c, 200515, "团队已提交报名")
		return
	}

	//要转移队长职位成员不在团队中
	if member.TeamID != user.TeamID {
		utils.JsonErrorResponse(c, 200517, "用户不在团队中")
		return
	}

	//可以转移队长职位
	err = userService.UpdateCaptainFlag(data.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = userService.UpdateCaptainFlag(data.MemberID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
