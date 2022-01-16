package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Group struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description"`
}

func GroupInitRouter(engine *gin.Engine) {
	r := engine.Group("/group")
	r.GET("/list", GroupList)
}

func GroupList(ctx *gin.Context) {
	time.Sleep(123 * time.Millisecond)
	ctx.JSON(200, gin.H{
		"rtn": 0,
		"data": gin.H{
			"list": []User{
				User{},
				User{},
			},
		},
	})
}
