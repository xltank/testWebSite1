package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "website/db"
	. "website/model"
	"website/res"
	. "website/utils"
)

func GroupInitRouter(engine *gin.Engine) {
	r := engine.Group("/group")
	r.GET("/list", list)
	r.POST("/", create)
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

func create(ctx *gin.Context) {
	var g Group
	err := ctx.ShouldBindJSON(&g)
	if err != nil {
		ReturnParamError(ctx, err.Error())
		return
	}

	r := Db.Create(&g)
	if r.Error != nil {
		ReturnParamError(ctx, err.Error())
		return
	}
	ReturnOK(ctx, g)
}
