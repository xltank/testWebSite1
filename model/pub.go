package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

func PubGetServerTime(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"rtn": 0,
		"data": gin.H{
			"now": time.Now(),
		},
	})
}
