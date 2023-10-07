package session

import (
	"TRS/config/redis"
	"github.com/gin-contrib/sessions"
	sessionRedis "github.com/gin-contrib/sessions/redis" //什么意思？？
	"github.com/gin-gonic/gin"
)

func setRedis(r *gin.Engine, name string) {
	Info := redis.RedisInfo
	store, _ := sessionRedis.NewStore(10, "tcp", Info.Host+":"+Info.Port, "", []byte("secret"))
	//[]byte("secret")是？？要改吗？
	r.Use(sessions.Sessions(name, store))
}
