package midware

import (
	"github.com/gin-gonic/gin"
	"log"
	. "website/db"
	. "website/model"
	"website/res"
	"website/utils"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookie, err := ctx.Cookie("token")

		log.Println("Auth, coolie: " + cookie)
		_, claims, err := utils.ParseToken(cookie)

		if err != nil {
			log.Println("get cookie error:", err)
			ctx.AbortWithStatusJSON(401, res.TokenParseErr)
			return // `return` not works. To return before other handlers, use Abortxxx().
		}
		log.Println("UserId from token: ", claims.UserId)
		//todo ? 查询用户信息
		var user User
		r := Db.First(&user, claims.UserId)
		if r.Error != nil {
			utils.SendParamError(ctx, r.Error.Error())
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
