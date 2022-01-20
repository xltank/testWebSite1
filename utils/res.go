package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"website/res"
)

func SendParamError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusBadRequest, res.ParamErr.Error(msg))
}

func SendOK(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusBadRequest, res.Ok.Json(data))
}
