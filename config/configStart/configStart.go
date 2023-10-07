package configStart

import (
	"TRS/config/config"
	"TRS/config/database"
	"TRS/config/redis"
)

func Init() {
	config.Init()
	database.Init()
	redis.Init()
}
