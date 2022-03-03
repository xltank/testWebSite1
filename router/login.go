package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
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

	originalPass := utils.Decode(u.Pass)

	// TODO: check email & password pattern

	// gen salt
	//salt1 := gonanoid.MustID(utils.Rand(10, 20))
	//salt2 := gonanoid.MustID(utils.Rand(10, 20))
	//fmt.Println(`---`, salt1, salt2)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(originalPass), bcrypt.DefaultCost)
	if err != nil {
		res.SendServerError(ctx, 0, err.Error())
		return
	}

	user := User{Email: u.Email, Pass: string(hashedPass)}
	r := Db.Create(&user)
	if r.Error != nil {
		res.SendParamError(ctx, 0, r.Error.Error())
		return
	}

	res.SendOK(ctx, gin.H{
		"email": user.Email,
		"id":    user.ID,
	})
}

func UserLogin(ctx *gin.Context) {
	var u UserLoginParam
	if e := ctx.ShouldBindJSON(&u); e != nil {
		res.SendParamError(ctx, 0, e.Error())
		return
	}

	originalPass := utils.Decode(u.Pass)
	log.Println(`originalPass `, originalPass)

	user := User{}
	r := Db.Where(User{Email: u.Email}).Find(&user)
	if r.Error != nil {
		res.SendParamError(ctx, 0, "User Not Found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(originalPass)); err != nil {
		res.SendParamError(ctx, 0, "Login info error"+err.Error())
		return
	}

	j, err := json.Marshal(&user)
	if err != nil {
		res.SendParamError(ctx, 0, "Login info error")
		return
	}
	fmt.Println("User:", j)

	user.Pass = ""
	token, err := utils.GetToken(user)
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
