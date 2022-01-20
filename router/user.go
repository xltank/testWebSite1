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
	. "website/model"
	. "website/utils"
)

func UserInitRouter(engine *gin.Engine) {
	r := engine.Group("/user")
	r.Use(midware.Auth())
	r.GET("/", UserList)
	r.POST("/", UserCreateMany)
	r.PUT("/", UserUpsertOne)
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
	var total int64
	var r *gorm.DB
	if q.Keyword != "" {
		r = db.Db.Where("name like ?", kw).Or("email like ?", kw).Or("department like ?", kw).Limit(q.Limit).Offset(q.Offset).Find(&users).Count(&total)
	} else {
		r = db.Db.Limit(q.Limit).Offset(q.Offset).Find(&users).Count(&total)
	}

	if r.Error != nil {
		ctx.JSON(400, error.NewParamError(r.Error))
		return
	}

	ReturnOK(ctx, gin.H{
		"list":   users,
		"offset": q.Offset,
		"limit":  q.Limit,
		"total":  total,
	})
}

func UserCreateMany(ctx *gin.Context) {
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

func UserUpsertOne(ctx *gin.Context) {
	var user User
	e := ctx.BindJSON(&user)
	if e != nil {
		ctx.JSON(400, error.NewParamError(e))
		return
	}

	r := db.Db.Save(&user)
	if r.Error != nil {
		ctx.JSON(400, error.NewServerError(r.Error))
		return
	}

	ctx.JSON(200, gin.H{
		"rtn":  0,
		"data": user,
	})
}

func UserDelete(ctx *gin.Context) {

}
