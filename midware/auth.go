package midware

import (
	"github.com/gin-gonic/gin"
	"log"
	"website/res"
	"website/utils"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookie, err := ctx.Cookie("token")

		log.Println("Auth, coolie: " + cookie)
		_, claims, err := utils.ParseToken(cookie)

		log.Println("User Info: ", claims)

		if err != nil {
			log.Println("get cookie error:", err)
			ctx.AbortWithStatusJSON(401, res.TokenParseErr)
			return // `return` not works. To return before other handlers, use Abortxxx().
		}
		//todo ? 查询用户信息
		ctx.Next()
	}
}
