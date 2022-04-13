package middlewares

import (
	"ecommerce/product/customerror"
	"ecommerce/product/rpc/client"
	"ecommerce/product/rpc/client/customer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfirmCustomer(ctx *gin.Context) {
	cid := ctx.GetString("uid")
	if cid == "" {
		panic("Empty uid in ConfirmCustomer middleware")
	}
	customerInfo := &customer.CustomerInfo{
		Id: cid,
	}
	customerClient, err := client.GetCustomerClient()
	if err != nil {
		panic(err)
	}
	isExisting, err := customerClient.RunIsExistingCustomer(customerInfo)
	if err != nil {
		panic(err)
	}
	if !isExisting {
		panic(customerror.NewAPIError(http.StatusUnauthorized, "Access denied"))
	}
	ctx.Next()
}
