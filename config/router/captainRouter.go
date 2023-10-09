package router

import (
	"TRS/app/controllers/captainController"
	"TRS/app/midwares"
	"github.com/gin-gonic/gin"
)

func captainRouterInit(r *gin.RouterGroup) {
	captain := r.Group("/captain")
	{
		captain.PUT("/update", midwares.CheckLogin, captainController.UpdateTeamInfo)
		captain.PUT("/apply", midwares.CheckLogin, captainController.Apply)
		captain.DELETE("/dis_apply", midwares.CheckLogin, captainController.DisApply)
		captain.DELETE("/remove", midwares.CheckLogin, captainController.DeleteTeamMember)
		captain.PUT("/add", midwares.CheckLogin, captainController.AddTeamMember)
		captain.DELETE("/delete", midwares.CheckLogin, captainController.DeleteTeam)
		captain.PUT("/transfer", midwares.CheckLogin, captainController.TransferCaptain)
	}
}
