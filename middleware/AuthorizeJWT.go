package middleware

import (
	"fmt"
	"go-jwt/handler"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header found"})

		}
		tokenString := authHeader[len(BearerSchema):]

		if token, err := handler.ValidateToken(tokenString); err != nil {

			fmt.Println("token", tokenString, err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Not Valid Token"})

		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				if token.Valid {
					ctx.Set("UserId", claims["userId"])
					fmt.Println("during authorization", claims["userId"])
				} else {
					ctx.AbortWithStatus((http.StatusUnauthorized))
				}
			}
		}
	}
}