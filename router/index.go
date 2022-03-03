package router

import (
	"github.com/gin-gonic/gin"
	"website/midware"
)

func InitRouter(engine *gin.Engine) {

	g := engine.Group("/api")
	g.POST("/signup", SignUp)
	g.POST("/login", UserLogin)

	PubInitRouter(engine)

	engine.Use(midware.Auth())

	g.POST("/logout", UserLogout)
	UserInitRouter(engine)
	GroupInitRouter(engine)
}
