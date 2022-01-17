package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"strings"
	"website/db"
	error "website/error"
	"website/midware"
)

type User struct {
	Model
	Name       string `json:"name,omitempty" binding:"required"`
	Email      string `json:"email,omitempty" binding:"required"`
	Pass       string `json:"pass,omitempty" binding:"required"`
	Role       string `json:"role,omitempty"` // max role: sa > admin > editor > member
	Department string `json:"department,omitempty"`
}

type UserQueryParam struct {
	Keyword string `form:"keyword"`
	Offset  int    `form:"offset"`
	Limit   int    `form:"limit"`
}

type UserLoginParam struct {
	Email string `binding:"required"`
	Pass  string `binding:"required"`
}

func UserInitRouter(engine *gin.Engine) {
	r := engine.Group("/user")
	r.Use(midware.Auth())
	r.GET("/", UserList)
	r.POST("/", UserCreateBulk)
}

func UserList(ctx *gin.Context) {
	var q UserQueryParam
	err := ctx.ShouldBindQuery(&q)
	if err != nil {
		ctx.JSON(400, error.NewParamError(err))
		return
	}
	fmt.Println(q)

	q.Keyword = strings.TrimSpace(q.Keyword)
	kw := "%" + q.Keyword + "%"
	var users []User
	var r *gorm.DB
	if q.Keyword != "" {
		r = db.Db.Where("name like ?", kw).Or("email like ?", kw).Or("department like ?", kw).Limit(q.Limit).Offset(q.Offset).Find(&users)
	} else {
		r = db.Db.Limit(q.Limit).Offset(q.Offset).Find(&users)
	}

	if r.Error != nil {
		ctx.JSON(400, error.NewParamError(r.Error))
		return
	}

	ctx.JSON(200, gin.H{
		"rtn": 0,
		"data": gin.H{
			"list":   users,
			"offset": q.Offset,
			"limit":  q.Limit,
		},
	})
}

func UserCreateBulk(ctx *gin.Context) {
	var users []User
	e := ctx.BindJSON(&users)
	if e != nil {
		ctx.JSON(400, error.NewParamError(e))
		return
	}
	log.Println("UserCreate, ", users)
	r := db.Db.Create(&users)
	if r.Error != nil {
		ctx.JSON(400, error.NewServerError(r.Error))
		return
	}

	ctx.JSON(200, gin.H{
		"rtn":  0,
		"data": users,
	})
}

func UserUpdate(ctx *gin.Context) {
	var user User
	e := ctx.BindJSON(&user)
	if e != nil {
		ctx.JSON(400, error.NewParamError(e))
		return
	}

	//r := db.Db.up
}

func UserDelete(ctx *gin.Context) {

}
