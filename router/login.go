package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"website/db"
	"website/res"
)

func UserLogin(ctx *gin.Context) {
	// fmt.Println(ctx.Request.Body)
	var u UserLoginParam
	if e := ctx.ShouldBind(&u); e != nil {
		ctx.JSON(400, res.ParamErr)
		return
	}

	user := User{}
	r := db.Db.Where(map[string]interface{}{"email": u.Email}).Find(&user)
	if r.Error != nil {
		ctx.JSON(400, res.ParamErr.Error("User Not Found"))
		return
	}
	//fmt.Println("RowsAffected:", r.RowsAffected)

	// todo: MD5
	if user.Pass != u.Pass {
		ctx.JSON(400, res.UserPasswordErr)
		return
	}

	j, err := json.Marshal(&user)
	if err != nil {
		ctx.JSON(400, res.MarshalJsonErr)
		return
	}
	ctx.SetCookie("token", string(j), 3600, "/", "my.com", false, true)

	ctx.JSON(200, res.Ok.Json(user))
}

func UserLogout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "my.com", false, true)
	ctx.JSON(200, res.Ok.Json(""))
}
