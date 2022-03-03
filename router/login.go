package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid"
	. "website/db"
	. "website/model"
	"website/res"
	"website/utils"
)

const passSalt = "jpijsdfvamsdvoasjdf"

func init() {
}

func SignUp(ctx *gin.Context) {
	var u UserLoginParam
	if e := ctx.ShouldBindJSON(&u); e != nil {
		res.SendParamError(ctx, 0, e.Error())
		return
	}
	fmt.Println(u)

	p1 := utils.Decode(u.Pass)

	// TODO: check email & password pattern

	// gen salt
	salt1 := gonanoid.MustID(utils.Rand(10, 20))
	salt2 := gonanoid.MustID(utils.Rand(10, 20))
	fmt.Println(`---`, salt1, salt2)

	res.SendOK(ctx, UserLoginParam{Email: u.Email, Pass: p1})
}

func UserLogin(ctx *gin.Context) {
	// fmt.Println(ctx.Request.Body)
	var u UserLoginParam
	if e := ctx.ShouldBindJSON(&u); e != nil {
		res.SendParamError(ctx, 0, e.Error())
		return
	}

	user := User{}
	r := Db.Where(map[string]interface{}{"email": u.Email}).Find(&user)
	if r.Error != nil {
		res.SendParamError(ctx, 0, "User Not Found")
		return
	}
	//fmt.Println("RowsAffected:", r.RowsAffected)

	// todo: MD5
	if user.Pass != u.Pass {
		res.SendParamError(ctx, 0, "Passwords do not match")
		return
	}

	j, err := json.Marshal(&user)
	if err != nil {
		res.SendParamError(ctx, 0, "Login info error")
		return
	}
	fmt.Println("User:", j)
	token, err := utils.GetToken(2) // todo ? id
	if err != nil {
		res.SendParamError(ctx, 0, "")
	}
	ctx.SetCookie("token", token, 24*3600, "/", "my.com", false, true)

	res.SendOK(ctx, user)
}

func UserLogout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "my.com", false, true)
	res.SendOK(ctx, "")
}
