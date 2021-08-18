package sessions

import (
	"github.com/SYSU-ECNC/workspace-be/internal/pkg/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store, _ = redis.NewStore(10, "tcp", config.Get("redis_addr"), "", []byte("secret"))

func Middleware() gin.HandlerFunc {
	return sessions.Sessions("ecnc_workspace", store)
}

func Store(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
