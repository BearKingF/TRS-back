package clientController

import (
	"TRS/app/models"
	"TRS/app/services/sessionService"
	"TRS/app/services/teamService"
	"TRS/app/services/userService"
	"TRS/app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* 编辑个人信息(PUT) */

type UpdateUserinfoData struct {
	// ID              uint   `json:"id" binding:"required"` //必需
	Username        string `json:"username"`
	Sex             string `json:"sex"`
	PhoneNum        string `json:"phone_num"`
	Email           string `json:"email"`
	Major           string `json:"major"`
	OldPassword     string `json:"old_password"` //如果要修改密码，需要先输入原密码
	NewPassword     string `json:"new_password"` //再两次输入一致的新密码
	ConfirmPassword string `json:"confirm_password"`
}

func UpdateUserInfo(c *gin.Context) {
	var data UpdateUserinfoData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	//不传入UserID的参数，而是用GetUserSession来获取当前登录的用户
	user, err := sessionService.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "未登录")
	}
	//user, _ := userService.GetUserByID(data.ID)
	//判断修改的手机号是否已被注册
	if data.PhoneNum != "" {
		err = userService.CheckUserExistByPhoneNum(data.PhoneNum)
		if err == nil {
			utils.JsonErrorResponse(c, 200504, "手机号已注册")
			return
		} else if err != gorm.ErrRecordNotFound {
			utils.JsonInternalServerErrorResponse(c)
			return
		}

	}
	//判断修改的邮箱是否已被注册
	if data.Email != "" {
		err = userService.CheckUserExistByEmail(data.Email)
		if err == nil {
			utils.JsonErrorResponse(c, 200505, "邮箱已注册")
			return
		} else if err != gorm.ErrRecordNotFound {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	//判断若要修改密码，输入的原密码是否正确及两次输入的新密码是否一致
	var flag bool
	if data.OldPassword != "" || data.NewPassword != "" || data.ConfirmPassword != "" {
		if data.OldPassword != "" && data.NewPassword != "" && data.ConfirmPassword != "" {
			flag = userService.CheckPwd(data.OldPassword, user.Password) && userService.CheckPwd(data.NewPassword, data.ConfirmPassword)
			if !flag {
				utils.JsonErrorResponse(c, 200506, "密码不一致")
				return
			}
		} else {
			utils.JsonErrorResponse(c, 200506, "密码不一致")
			return
		}
	}

	// 更新用户信息
	err = userService.UpdateUserInfo(models.User{
		//ID:       data.ID,
		ID:       user.ID,
		Username: data.Username,
		Sex:      data.Sex,
		PhoneNum: data.PhoneNum,
		Email:    data.Email,
		Major:    data.Major,
		Password: data.NewPassword,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

/* 创建团队(POST) */

type CreateTeamData struct {
	//UserID   uint   `json:"user_id" binding:"required"`
	TeamName string `json:"team_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateTeam(c *gin.Context) {
	var data CreateTeamData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	user, err := sessionService.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "未登录")
	}

	// user, _ := userService.GetUserByID(data.UserID)
	if user.TeamID != -1 {
		utils.JsonErrorResponse(c, 200510, "已加入团队") //限制每人只可加入一个团队
		return
	}

	//判断团队名是否已被注册
	err = teamService.CheckTeamExistByTeamName(data.TeamName)
	if err == nil { //团队名已存在
		utils.JsonErrorResponse(c, 200512, "团队名已注册")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//可以创建团队
	teamID, err := teamService.CreateTeam(models.Team{
		TeamName:  data.TeamName,
		CaptainID: user.ID,
		//CaptainID: data.UserID,
		Password: data.Password,
		Total:    1,
		Status:   2, //状态：未提交报名
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	_ = userService.UpdateTeamID(user.ID, int(teamID))
	_ = userService.UpdateCaptainFlag(user.ID) //团队创建者成为队长
	//返回成功响应
	utils.JsonSuccessResponse(c, nil)
}

/* 加入团队(POST) */

type JoinTeamData struct {
	//UserID   uint   `json:"user_id" binding:"required"`
	TeamID   uint   `json:"team_id" binding:"required"`  //团队ID
	Password string `json:"password" binding:"required"` //团队密码
}

func JoinTeam(c *gin.Context) {
	var data JoinTeamData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	user, err := sessionService.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "未登录")
	}

	err = teamService.CheckTeamExistByTeamID(data.TeamID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200508, "团队不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	//团队存在
	team, _ := teamService.GetTeamByTeamID(data.TeamID)
	if data.Password != team.Password {
		utils.JsonErrorResponse(c, 200506, "密码不一致")
		return
	}
	//密码正确
	if team.Total == 6 {
		utils.JsonErrorResponse(c, 200509, "团队已满")
		return
	}
	//若团队已提交报名，不可再加入该团队
	if team.Status == 1 {
		utils.JsonErrorResponse(c, 200515, "团队已提交报名")
		return
	}
	// user, _ := userService.GetUserByID(data.UserID)
	if user.TeamID != -1 {
		utils.JsonErrorResponse(c, 200510, "已加入团队") //限制每人只可加入一个团队
		return
	}

	//可以加入该团队
	_ = userService.UpdateTeamID(user.ID, int(data.TeamID))
	_ = teamService.UpdateTotal(data.TeamID)
	//返回成功响应
	utils.JsonSuccessResponse(c, nil)

}

/* 获取团队信息(GET) */

type GetTeamInfoData struct {
	// UserID uint `form:"user_id" binding:"required"` //form???
	TeamID uint `form:"team_id" binding:"required"`
}

func GetTeamInfo(c *gin.Context) {
	var data GetTeamInfoData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}

	user, err := sessionService.GetUserSession(c)
	if err != nil {
		utils.JsonErrorResponse(c, 200507, "未登录")
	}

	// user, _ := userService.GetUserByID(data.UserID)
	if user.TeamID != int(data.TeamID) {
		utils.JsonErrorResponse(c, 200513, "权限错误")
		return
	}

	//可以获取团队信息
	team, _ := teamService.GetTeamByTeamID(data.TeamID)
	captain, _ := userService.GetUserByID(team.CaptainID)
	var teamMemberList []models.User //创建切片，存储所有队员对象
	//团队中必定有成员
	teamMemberList, _ = teamService.GetTeamMember(data.TeamID)

	utils.JsonSuccessResponse(c, gin.H{
		"team_name":        team.TeamName,
		"captain":          captain.Username,
		"team_member_list": teamMemberList,
		"team_num":         team.Total,
		"team_status":      team.Status, //团队状态 (new)
	})
}
