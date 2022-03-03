package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"website/config"
	"website/db"
	"website/midware"
	"website/router"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Println("Init config error: ", err)
		return
	}

	r := gin.New()

	r.Static("/public", "./public")

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		//log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusInternalServerError, gin.H{"rtn": 1000})
	}))

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s \"%s\" %d %s \"%s\" %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Request.Proto,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.Use(midware.TimeCost())

	r.Use(midware.CORSMiddleware())

	db.InitMysql()

	router.InitRouter(r)

	r.Run()
}
