package middleware

import (
	"gin/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var JWTSecretKey = []byte("secret-key")

func ValidateAPIKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const apiKey string = "aw13u210312321"
		//check header ada atau tidak
		key := ctx.GetHeader("x-api-key")

		//validasi api key
		if key != apiKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("Unauthorized", nil))
			return
		}

		ctx.Next()
	}
}

func ValidateUserToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//check header
		header := ctx.GetHeader("Authorization")

		tokenRequest := strings.TrimPrefix(header, "Bearer ")

		token, err := jwt.Parse(tokenRequest, func(t *jwt.Token) (interface{}, error) {
			return JWTSecretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, util.BuildResponse("Unauthorized", nil))
			return
		}

		ctx.Next()
	}
}
