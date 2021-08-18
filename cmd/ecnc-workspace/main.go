package main

import (
	"github.com/SYSU-ECNC/workspace-be/internal/pkg/auth"
	"github.com/SYSU-ECNC/workspace-be/internal/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(sessions.Middleware())

	r.GET("/sso/authorize", auth.Authorize)
	r.GET("/sso/callback", auth.Callback)

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Store(c)

		c.JSON(200, gin.H{"user": session.Get("user")})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
