package router

import (
	"time"
	"websiteGin/res"

	"github.com/gin-gonic/gin"
)

func PubInitRouter(engine *gin.Engine) {
	r := engine.Group("/api/pub")
	r.GET("/ping", PubGetServerTime)
	r.GET("/panic", func(c *gin.Context) {
		panic("panic info: test panic")
	})

	r.GET("/file/main", func(c *gin.Context) {
		c.File("main.go")
	})
}

func PubGetServerTime(c *gin.Context) {
	// fmt.Println("Pub Get Server Time")
	res.SendOK(c, gin.H{
		"now": time.Now(),
	})
}
