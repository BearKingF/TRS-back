package router

import (
	"TRS/app/controllers/clientController"
	"github.com/gin-gonic/gin"
)

func clientRouterInit(r *gin.RouterGroup) {
	client := r.Group("/client")
	{
		client.PUT("/update", clientController.UpdateUserInfo) //编辑个人信息
		client.POST("/create", clientController.CreateTeam)    //创建团队
		client.POST("/join", clientController.JoinTeam)        //加入团队
		client.GET("/get", clientController.GetTeamInfo)       //获取团队信息

	}
}
