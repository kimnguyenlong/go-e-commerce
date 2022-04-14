package routes

import (
	"ecommerce/order/controllers"
	"ecommerce/order/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRouteOrders(rootGroup *gin.RouterGroup) error {
	orderController, err := controllers.NewOrderController()
	if err != nil {
		return err
	}
	cartsGroup := rootGroup.Group("/orders").Use(middlewares.Authenticate, middlewares.ConfirmCustomer)
	{
		cartsGroup.POST("/", orderController.CreateOrder())
		cartsGroup.GET("/", orderController.GetOrders())
		cartsGroup.GET("/:id", orderController.GetSingleController())
	}
	return nil
}
