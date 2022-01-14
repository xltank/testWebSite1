package main

import (
	"github.com/gin-gonic/gin"
	"website/midware"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(midware.TimeCost())

	InitRouter(r)

	r.Run()
}
