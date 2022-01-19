package router

import (
	"github.com/gin-gonic/gin"
	"time"
	"website/res"
)

type Group struct {
	Model
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Desc string `json:"description"`
}

func GroupInitRouter(engine *gin.Engine) {
	r := engine.Group("/group")
	r.GET("/list", GroupList)
}

func GroupList(ctx *gin.Context) {
	time.Sleep(123 * time.Millisecond)
	obj := make(map[string]User, 20)
	obj["list"] = User{}
	ctx.JSON(200, res.Ok.Json(obj))

}
