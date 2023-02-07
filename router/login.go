package router

import (
	"encoding/json"
	"fmt"
	"log"
	. "websiteGin/db"
	. "websiteGin/model"
	"websiteGin/res"
	"websiteGin/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func init() {
}

func SignUp(c *gin.Context) {
	var u UserLoginParam
	if e := c.ShouldBindJSON(&u); e != nil {
		res.SendParamError(c, 0, e.Error())
		return
	}

	originalPass := utils.DecodeRSA(u.Pass)

	log.Println("Singup,", u.Email, originalPass)
	// TODO: check email & password pattern

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(originalPass), bcrypt.DefaultCost)
	if err != nil {
		res.SendServerError(c, 0, err.Error())
		return
	}

	user := User{Email: u.Email, Pass: string(hashedPass)}
	r := Db.Create(&user)
	if r.Error != nil {
		res.SendParamError(c, 0, r.Error.Error())
		return
	}

	res.SendOK(c, gin.H{
		"email": user.Email,
		"id":    user.ID,
	})
}

func UserLogin(c *gin.Context) {
	var u UserLoginParam
	if e := c.ShouldBindJSON(&u); e != nil {
		res.SendParamError(c, 0, e.Error())
		return
	}

	originalPass := utils.DecodeRSA(u.Pass)
	log.Println(`originalPass `, originalPass)

	user := User{}
	r := Db.Where(User{Email: u.Email}).Find(&user)
	if r.Error != nil {
		res.SendParamError(c, 0, "User Not Found")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(originalPass)); err != nil {
		res.SendParamError(c, 0, "Login info error"+err.Error())
		return
	}

	j, err := json.Marshal(&user)
	if err != nil {
		res.SendParamError(c, 0, "Login info error")
		return
	}
	fmt.Println("User:", j)

	user.Pass = ""
	token, err := utils.GetToken(user)
	if err != nil {
		res.SendParamError(c, 0, "")
	}

	c.SetCookie("token", token, 24*3600, "/", "my.com", false, true)
	res.SendOK(c, user)
}

func UserLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "my.com", false, true)
	res.SendOK(c, "")
}
