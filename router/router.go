package router

import (
	"github.com/gin-gonic/gin"
	"website/model"
)

func InitRouter(r *gin.Engine) {

	r.POST("/login", model.UserLogin)

	model.PubInitRouter(r)
	model.UserInitRouter(r)
	model.GroupInitRouter(r)
}
