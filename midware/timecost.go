package midware

import (
	"fmt"
	"time"
	"websiteGin/utils"

	"github.com/gin-gonic/gin"
)

func TimeCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		d := time.Now().Sub(t)
		fmt.Println("<---", c.Request.URL, utils.ToFixed(float64(d)/1000000, 2), "ms")
	}
}
