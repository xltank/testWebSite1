package midware

import (
	"github.com/gin-gonic/gin"
	"log"
	"website/res"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			log.Println("get cookie error:", err)
			ctx.AbortWithStatusJSON(401, res.AuthErr)
			return // `return` not works. To return before other handlers, use Abortxxx().
		}
		log.Println("Auth, coolie: " + cookie)

		ctx.Next()
	}
}
