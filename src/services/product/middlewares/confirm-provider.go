package middlewares

import (
	"ecommerce/product/customerror"
	"ecommerce/product/rpc/client"
	"ecommerce/product/rpc/client/provider"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfirmProvider(ctx *gin.Context) {
	pid := ctx.GetString("uid")
	if pid == "" {
		panic("Empty uid in ConfirmProvider middleware")
	}
	providerInfo := &provider.ProviderInfo{
		PID: pid,
	}
	providerClient, err := client.GetProviderClient()
	if err != nil {
		panic(err)
	}
	isExisting, err := providerClient.RunIsExistingProvider(providerInfo)
	if err != nil {
		panic(err)
	}
	if !isExisting {
		panic(customerror.NewAPIError(http.StatusUnauthorized, "Access denied"))
	}
	ctx.Next()
}
