package middlewares

import (
	"ecommerce/cart/customerror"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			code := http.StatusInternalServerError
			message := "Something went wrong, please try again later!"
			switch err.(type) {
			case customerror.APIError:
				apiErr := err.(customerror.APIError)
				code = apiErr.Code
				message = apiErr.Message
			default:
				log.Printf("Internal error: %s\n", err)
			}
			ctx.JSON(code, gin.H{
				"error":   true,
				"message": message,
			})
			ctx.Abort()
		}
	}()
	ctx.Next()
}
