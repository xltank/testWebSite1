package main

import (
	"github.com/gin-gonic/gin"
	"website/model"
)

func InitRouter(r *gin.Engine) {

	r.GET("/ping", model.PubGetServerTime)

	r.POST("/login", model.UserLogin)

	userRouter := r.Group("/user")
	userRouter.GET("/list", model.UserList)

	groupRouter := r.Group("/group")
	groupRouter.GET("/list", model.GroupList)

}
