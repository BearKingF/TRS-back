package router

import (
	"TRS/app/controllers/adminController"
	"TRS/app/midwares"
	"github.com/gin-gonic/gin"
)

func adminRouterInit(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	{
		admin.GET("/get", midwares.CheckLogin, adminController.GetAllTeamInfo)
	}
}
