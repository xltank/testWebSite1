package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website/res"
)

func SendOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusBadRequest, res.Ok.Json(data))
}

func SendParamError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, res.ParamErr.Error(msg))
}

func SendServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, res.ServerErr.Error(msg))
}
