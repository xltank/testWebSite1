package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"website/config"
	"website/db"
	error2 "website/error"
	"website/midware"
	"website/router"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Println("Init config error: ", err)
		return
	}

	r := gin.Default()

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, error2.NewServerError(err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	//r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//	return fmt.Sprintf("%s - [%s] %s %s \"%s\" %d %s \"%s\" %s\n",
	//		param.TimeStamp.Format(time.RFC3339),
	//		param.ClientIP,
	//		param.Method,
	//		param.Request.Proto,
	//		param.Path,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))

	r.Use(midware.TimeCost())

	db.InitMysql()

	router.InitRouter(r)

	r.Run()
}
