package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data interface{}) {
	c.JSON(httpStatusCode, gin.H{ //httpStatusCode: 响应状态码
		"code": code, //自定义状态码
		"msg":  msg,  //错误信息
		"data": data, //响应数据
	})
}

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	JsonResponse(c, http.StatusOK, 200, "OK", data)
}

func JsonErrorResponse(c *gin.Context, code int, msg string) {
	JsonResponse(c, http.StatusInternalServerError, code, msg, nil)
}

func JsonInternalServerErrorResponse(c *gin.Context) {
	JsonResponse(c, http.StatusInternalServerError, 200500, "Internal server error", nil)
}
