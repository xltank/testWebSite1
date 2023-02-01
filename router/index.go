package router

import (
	"websiteGin/midware"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	g := engine.Group("/api")
	g.POST("/signup", SignUp)
	g.POST("/login", UserLogin)
	g.POST("/logout", UserLogout)

	PubInitRouter(engine)

	engine.Use(midware.Auth())

	UserInitRouter(engine)
	GroupInitRouter(engine)
}
