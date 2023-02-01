package router

import (
	. "websiteGin/db"
	. "websiteGin/model"
	"websiteGin/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GroupInitRouter(engine *gin.Engine) {
	r := engine.Group("/api/group")
	r.GET("/list", list)
	r.POST("/", createMany)
}

func list(c *gin.Context) {
	var q GroupQueryParam
	err := c.ShouldBindQuery(q)
	if err != nil {
		res.SendOK(c, "Invalid Params")
		return
	}

	var groups []Group
	Db.Offset(q.Offset).Limit(q.Limit).Find(&groups)
	res.SendOK(c, groups)
}

func createMany(c *gin.Context) {
	var gs []Group
	err := c.ShouldBindJSON(&gs)
	if err != nil {
		res.SendParamError(c, 0, err.Error())
		return
	}

	Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&gs)
	res.SendOK(c, gs)
}
