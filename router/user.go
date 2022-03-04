package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strings"
	. "website/db"
	. "website/model"
	"website/res"
)

var UserFields = []string{"email", "id", "name", "department", "role", "createdAt", "updatedAt"}

func UserInitRouter(engine *gin.Engine) {
	r := engine.Group("/api/user")
	r.GET("/", UserList)
	r.POST("/", UserCreateMany)
	r.PUT("/", UserUpsertOne)
	r.POST("/:uid/group/:gid/role/:role", UserAddToGroup)
}

func UserList(ctx *gin.Context) {
	// test codes
	v, ok := ctx.Get("user")
	if !ok {
		res.SendServerError(ctx, 0, "Not found ctx.user")
		return
	}
	var user User
	if user, ok = v.(User); !ok {
		res.SendServerError(ctx, 0, "Invalid ctx.user")
		return
	}
	log.Println("ctx.user:", user)

	var q UserQueryParam
	err := ctx.ShouldBindQuery(&q)
	if err != nil {
		res.SendParamError(ctx, 0, "")
		return
	}
	fmt.Println("UserQueryParam:", q)

	q.Keyword = strings.TrimSpace(q.Keyword)
	kw := "%" + q.Keyword + "%"
	var users []User
	var total int64
	var r *gorm.DB
	if q.Keyword != "" {
		r = Db.Model(&user).Preload("Groups").Where("name like ?", kw).Or("email like ?", kw).Or("department like ?", kw).Count(&total).Order("createdAt desc").Offset(q.Offset).Limit(q.Limit).Find(&users).Select(UserFields)
	} else {
		r = Db.Model(&user).Preload("Groups").Count(&total).Order("createdAt desc").Offset(q.Offset).Limit(q.Limit).Find(&users).Select(UserFields)
	}

	if r.Error != nil {
		res.SendParamError(ctx, 0, r.Error.Error())
		return
	}

	res.SendOK(ctx, gin.H{
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
		res.SendParamError(ctx, 0, e.Error())
		return
	}
	log.Println("UserCreate, ", users)
	r := Db.Create(&users)
	if r.Error != nil {
		res.SendParamError(ctx, 0, r.Error.Error())
		return
	}

	res.SendOK(ctx, users)
}

func UserUpsertOne(ctx *gin.Context) {
	var user User
	e := ctx.BindJSON(&user)
	if e != nil {
		res.SendParamError(ctx, 0, e.Error())
		return
	}
	user.Pass = ""

	r := Db.Save(&user)
	if r.Error != nil {
		res.SendParamError(ctx, 0, r.Error.Error())
		return
	}

	res.SendOK(ctx, user)

}

func UserAddToGroup(ctx *gin.Context) {
	var ug UserGroup
	err := ctx.ShouldBindUri(&ug)
	if err != nil {
		res.SendParamError(ctx, 0, err.Error())
		return
	}
	log.Println("ug:", ug)

	// Option 1
	//没办法写入 role 字段
	//err = Db.Model(&User{Model: Model{ID: ug.UserId}}).Where("id = ?", ug.UserId).Association("Groups").Append(&Group{Model: Model{ID: ug.GroupId}})

	// Option 2
	// 需要给user_group表创建unique联合索引，而且：1，不需要指定 Columns；2，不需要在UserGroup model里声明 gorm 标注；3，不需要 SetupJoinTable；
	// 文档：INSERT ... ON DUPLICATE KEY UPDATE is a MariaDB/MySQL extension to the INSERT statement that,
	// 		if it finds a duplicate unique or primary key, will instead perform an UPDATE.
	// 语句：INSERT INTO `group` (`createdAt`,`updatedAt`,`name`,`desc`) VALUES ('2022-01-21 12:18:31.518','2022-01-21 12:18:31.518','group3','desc2'),('2022-01-21 12:18:31.518','2022-01-21 12:18:31.518','group4','desc3') ON DUPLICATE KEY UPDATE `id`=`id`
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
		res.SendParamError(ctx, 0, err.Error())
		return
	}
	res.SendOK(ctx, ug)
}
