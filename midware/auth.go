package midware

import (
	"github.com/gin-gonic/gin"
	"log"
	"website/error"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			log.Println("get cookie error:", err)
			ctx.AbortWithStatusJSON(401, error.NewAuthError(err))
			return // to return before other handlers, use Abortxxx(). Otherwise, `return` not works.
		}
		log.Println("Auth, coolie: " + cookie)

		ctx.Next()
	}
}