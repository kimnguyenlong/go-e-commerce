package routes

import (
	"ecommerce/cart/controllers"
	"ecommerce/cart/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRouteCarts(rootGroup *gin.RouterGroup) error {
	cartController, err := controllers.NewCartController()
	if err != nil {
		return err
	}
	cartsGroup := rootGroup.Group("/carts").Use(middlewares.Authenticate, middlewares.ConfirmCustomer)
	{
		cartsGroup.POST("/", cartController.SetCart())
		cartsGroup.GET("/:id", cartController.GetCart())
	}
	return nil
}
