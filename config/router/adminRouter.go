package router

import (
	"TRS/app/controllers/adminController"
	"github.com/gin-gonic/gin"
)

func adminRouterInit(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	{
		admin.GET("/get", adminController.GetAllTeamInfo)
	}
}
