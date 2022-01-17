package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"website/db"
	"website/error"
)

func UserLogin(ctx *gin.Context) {
	// fmt.Println(ctx.Request.Body)
	var u UserLoginParam
	if e := ctx.ShouldBind(&u); e != nil {
		ctx.JSON(400, error.NewParamError(e.Error()))
		return
	}

	user := User{}
	r := db.Db.Where(map[string]interface{}{"email": u.Email}).Find(&user)
	if r.Error != nil {
		ctx.JSON(400, error.NewLoginError(r.Error.Error()))
		return
	}
	//fmt.Println("RowsAffected:", r.RowsAffected)

	// todo: MD5
	if user.Pass != u.Pass {
		ctx.JSON(400, error.NewLoginError(nil))
		return
	}

	j, err := json.Marshal(&user)
	if err != nil {
		ctx.JSON(400, error.NewParamError(err))
		return
	}
	ctx.SetCookie("token", string(j), 3600, "/", "my.com", false, true)

	ctx.JSON(200, gin.H{
		"rtn":  0,
		"data": user,
	})
}

func UserLogout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "my.com", false, true)
	ctx.JSON(200, gin.H{
		"rtn":  0,
		"data": "",
	})
}
