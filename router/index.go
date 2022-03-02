package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	g := engine.Group("/api")
	g.POST("/signup", SignUp)
	g.POST("/login", UserLogin)
	g.POST("/logout", UserLogout)

	PubInitRouter(engine)
	UserInitRouter(engine)
	GroupInitRouter(engine)
}
