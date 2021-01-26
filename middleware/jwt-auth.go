package middleware

import (
	"log"
	"net/http"

	"Atlantis-Backend/helper"
	"Atlantis-Backend/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT to authorize jwt in api
func AuthorizeJWT(jwtServ service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			res := helper.ResponseFailed("Failed to process request", "Token invalid", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		token, err := jwtServ.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[userID]: ", claims["userID"])
			log.Println("Claims[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.ResponseFailed("Token invalid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
