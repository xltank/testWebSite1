package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.POST("/login", UserLogin)
	r.POST("/logout", UserLogout)

	PubInitRouter(r)
	UserInitRouter(r)
	GroupInitRouter(r)
}
