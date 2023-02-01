package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"websiteGin/config"
	"websiteGin/db"
	"websiteGin/midware"
	"websiteGin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Println("Init config error: ", err)
		return
	}

	r := gin.New()

	r.LoadHTMLGlob("public/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.StaticFile("/favicon.ico", "./public/imgs/favicon.ico")
	r.Static("/public", "./public")

	// gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	// 	// log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	// }

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		// c.JSON(http.StatusInternalServerError, gin.H{"rtn": 1000})
	}))

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s \"%s\" %d %s \"%s\" %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
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

	// r.Run(":" + config.Conf.Port)
	s := &http.Server{
		Addr:    ":" + config.Conf.Port,
		Handler: r,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 2MB
	}
	s.ListenAndServe()
}
