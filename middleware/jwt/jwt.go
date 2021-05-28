package jwt

import (
	"net/http"
	"time"

	"gin-jwt-demo/pkg/code"
	"gin-jwt-demo/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 根据实际情况取TOKEN, 这里从request header取
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": code.ERR_AUTH_NULL,
				"msg":  code.GetMsg(code.ERR_AUTH_NULL),
			})
			return
		}

		claims, err := jwt.ParseToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": code.ERR_AUTH_INVALID,
				"msg":  code.GetMsg(code.ERR_AUTH_INVALID),
			})
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": code.ERR_AUTH_EXPIRED,
				"msg":  code.GetMsg(code.ERR_AUTH_EXPIRED),
			})
			return
		}

		// 此处已经通过了, 可以把Claims中的有效信息拿出来放入上下文使用
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}
