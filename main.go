package main

import (
	"TRS/config/configStart"
	"TRS/config/router"
	"TRS/config/session"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configStart.Init()

	r := gin.Default()
	session.Init(r)
	router.Init(r)
	//r.Use(sessions.Sessions("mysession", store))

	err := r.Run()
	//err := r.Run(":" + config.Config.GetString("server.port")) //原来是r.Run()
	if err != nil {
		log.Fatal("Server start error: ", err)
	}
}
