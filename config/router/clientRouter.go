package router

import (
	"TRS/app/controllers/clientController"
	"TRS/app/midwares"
	"github.com/gin-gonic/gin"
)

func clientRouterInit(r *gin.RouterGroup) {

	client := r.Group("/client")
	{
		client.PUT("/update", midwares.CheckLogin, clientController.UpdateUserInfo) //编辑个人信息
		client.POST("/create", midwares.CheckLogin, clientController.CreateTeam)    //创建团队
		client.POST("/join", midwares.CheckLogin, clientController.JoinTeam)        //加入团队
		client.GET("/get", midwares.CheckLogin, clientController.GetTeamInfo)       //获取团队信息
		captainRouterInit(client)

	}
}
