package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"website/utils"
)

func TimeCost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Next()

		d := time.Now().Sub(t)
		fmt.Println("<---", ctx.Request.URL, utils.ToFixed(float64(d)/1000000, 2), "ms")
	}
}
