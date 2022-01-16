package main

import (
	"github.com/gin-gonic/gin"
	"website/midware"
	"website/router"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(midware.TimeCost())

	router.InitRouter(r)

	r.Run()
}
