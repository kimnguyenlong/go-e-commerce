package middlewares

import (
	"ecommerce/order/customerror"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authenticate(ctx *gin.Context) {
	authorizations := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(authorizations) != 2 || authorizations[0] != "Bearer" {
		panic(customerror.NewAPIError(http.StatusUnauthorized, "Please provide a Bearer token"))
	}
	token, err := jwt.Parse(authorizations[1], func(t *jwt.Token) (interface{}, error) {
		if signMethod := t.Method.Alg(); signMethod != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("Unexpected signing method: %s", signMethod)
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		panic(customerror.NewAPIError(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %s", err)))
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Unix() {
			panic(customerror.NewAPIError(http.StatusUnauthorized, "Token is expired"))
		}
		id := claims["id"].(string)
		ctx.Set("uid", id)
		ctx.Next()
	} else {
		panic(customerror.NewAPIError(http.StatusUnauthorized, "Invalid token"))
	}
}
