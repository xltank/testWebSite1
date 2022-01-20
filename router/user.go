package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strings"
	. "website/db"
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
	r.POST("/:uid/group/:gid/role/:role", UserAddGroup)
}

func UserList(ctx *gin.Context) {
	var q UserQueryParam
	err := ctx.ShouldBindQuery(&q)
	if err != nil {
		ctx.JSON(400, NewParamError(err))
		return
	}
	fmt.Println(q)

	q.Keyword = strings.TrimSpace(q.Keyword)
	kw := "%" + q.Keyword + "%"
	var users []User
	var total int64
	var r *gorm.DB
	if q.Keyword != "" {
		r = Db.Preload("Groups").Where("name like ?", kw).Or("email like ?", kw).Or("department like ?", kw).Limit(q.Limit).Offset(q.Offset).Find(&users).Count(&total)
	} else {
		r = Db.Preload("Groups").Limit(q.Limit).Offset(q.Offset).Find(&users).Count(&total)
	}

	if r.Error != nil {
		ctx.JSON(400, NewParamError(r.Error))
		return
	}

	SendOK(ctx, gin.H{
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
		ctx.JSON(400, NewParamError(e))
		return
	}
	log.Println("UserCreate, ", users)
	r := Db.Create(&users)
	if r.Error != nil {
		ctx.JSON(400, NewServerError(r.Error))
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
		ctx.JSON(400, NewParamError(e))
		return
	}

	r := Db.Save(&user)
	if r.Error != nil {
		ctx.JSON(400, NewServerError(r.Error))
		return
	}

	ctx.JSON(200, gin.H{
		"rtn":  0,
		"data": user,
	})

}

func UserAddGroup(ctx *gin.Context) {
	var ug UserGroup
	err := ctx.ShouldBindUri(&ug)
	if err != nil {
		SendParamError(ctx, err.Error())
		return
	}
	log.Println("ug:", ug)

	// Option 1
	//没办法写入 role 字段
	//err = Db.Model(&User{Model: Model{ID: ug.UserId}}).Where("id = ?", ug.UserId).Association("Groups").Append(&Group{Model: Model{ID: ug.GroupId}})

	// Option 2
	// 需要给user_group表创建unique联合索引，而且：1，不需要指定 Columns；2，不需要在UserGroup model里声明 gorm 标注；3，不需要 SetupJoinTable；
	Db.Clauses(clause.OnConflict{
		//Columns:   []clause.Column{{Name: "user_id"}, {Name: "group_id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"role": ug.Role}),
	}).Create(&ug)

	// Option 3
	/*err = Db.Transaction(func(tx *gorm.DB) error {
		r := Db.Where("user_id = ? AND group_id = ?", ug.UserId, ug.GroupId).Find(&UserGroup{})
		if r.Error != nil {
			return r.Error
		}
		if r.RowsAffected == 0 {
			Db.Create(&ug)
		} else {
			Db.Where("user_id = ? AND group_id = ?", ug.UserId, ug.GroupId).Updates(&ug)
		}
		return nil
	})*/

	if err != nil {
		SendParamError(ctx, err.Error())
		return
	}
	SendOK(ctx, ug)
}
