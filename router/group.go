package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	. "website/db"
	. "website/model"
	"website/res"
)

func GroupInitRouter(engine *gin.Engine) {
	r := engine.Group("/api/group")
	r.GET("/list", list)
	r.POST("/", createMany)
}

func list(ctx *gin.Context) {
	var q GroupQueryParam
	err := ctx.ShouldBindQuery(q)
	if err != nil {
		res.SendOK(ctx, "Invalid Params")
		return
	}

	var groups []Group
	Db.Offset(q.Offset).Limit(q.Limit).Find(&groups)
	res.SendOK(ctx, groups)
}

func createMany(ctx *gin.Context) {
	var gs []Group
	err := ctx.ShouldBindJSON(&gs)
	if err != nil {
		res.SendParamError(ctx, 0, err.Error())
		return
	}

	Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&gs)
	res.SendOK(ctx, gs)
}
