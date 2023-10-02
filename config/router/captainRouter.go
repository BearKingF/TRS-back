package router

import (
	"TRS/app/controllers/captainController"
	"github.com/gin-gonic/gin"
)

func captainRouterInit(r *gin.RouterGroup) {
	captain := r.Group("/captain")
	{
		captain.PUT("/update", captainController.UpdateTeamInfo)
		captain.PUT("/apply", captainController.Apply)
		captain.DELETE("/dis_apply", captainController.DisApply)
		captain.DELETE("/remove", captainController.DeleteTeamMember)
		captain.PUT("/add", captainController.AddTeamMember)
		captain.DELETE("/delete", captainController.DeleteTeam)
		captain.PUT("/transfer", captainController.TransferCaptain)
	}
}
