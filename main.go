package main

import (
	"TRS/config/database"
	"TRS/config/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Init()
	r := gin.Default()
	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal("Server start error: ", err)
	}
}
