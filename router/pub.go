package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"website/res"
)

func PubInitRouter(engine *gin.Engine) {
	r := engine.Group("/api/pub")
	r.GET("/ping", PubGetServerTime)
	r.GET("/panic", func(ctx *gin.Context) {
		panic("/pub/panic")
	})

	r.GET("/file/main", func(c *gin.Context) {
		c.File("main.go")
	})
}

func PubGetServerTime(ctx *gin.Context) {
	fmt.Println("Pub Get Server Time")
	res.SendOK(ctx, gin.H{
		"rtn": 0,
		"data": gin.H{
			"now": time.Now(),
		},
	})
}
