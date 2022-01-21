package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
	. "website/db"
	. "website/model"
	"website/res"
	. "website/utils"
)

func GroupInitRouter(engine *gin.Engine) {
	r := engine.Group("/group")
	r.GET("/list", list)
	r.POST("/", createMany)
}

func list(ctx *gin.Context) {
	var q GroupQueryParam
	err := ctx.ShouldBindQuery(q)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.ParamErr.Error("Invalid Params"))
		return
	}

	var groups []Group
	Db.Offset(q.Offset).Limit(q.Limit).Find(&groups)
	ctx.JSON(200, res.Ok.Json(groups))
}

func createMany(ctx *gin.Context) {
	var gs []Group
	err := ctx.ShouldBindJSON(&gs)
	if err != nil {
		SendParamError(ctx, err.Error())
		return
	}

	Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&gs)
	SendOK(ctx, gs)
}
