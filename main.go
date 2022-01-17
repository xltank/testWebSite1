package main

import (
	"github.com/gin-gonic/gin"
	"website/db"
	"website/midware"
	"website/router"
)

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.Use(midware.TimeCost())

	db.InitMysql()

	router.InitRouter(r)

	r.Run()
}
