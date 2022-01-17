package model

import (
	"github.com/gin-gonic/gin"
	"website/db"
	error2 "website/error"
	"website/midware"
)

type User struct {
	Model
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Pass       string `json:"pass,omitempty"`
	Role       string `json:"role,omitempty"` // max role: sa > admin > editor > member
	Department string `json:"department,omitempty"`
}

func UserInitRouter(engine *gin.Engine) {
	r := engine.Group("/user")
	r.Use(midware.Auth())
	r.GET("/", UserList)
}

func UserList(ctx *gin.Context) {
	var users []User
	r := db.Db.Find(&users)
	if r.Error != nil {
		ctx.JSON(3001, error2.NewParamError(r.Error))
		return
	}

	ctx.JSON(200, gin.H{
		"rtn": 0,
		"data": gin.H{
			"list": users,
		},
	})
}

func UserCreate(ctx *gin.Context) {

}

func UserUpdate(ctx *gin.Context) {

}

func UserDelete(ctx *gin.Context) {

}
