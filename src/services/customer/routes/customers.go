package routes

import (
	"ecommerce/customer/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRouteCustomers(rootGroup *gin.RouterGroup) error {
	customersController, err := controllers.NewCustomersController()
	if err != nil {
		return err
	}
	customersGroup := rootGroup.Group("/customers")
	{
		customersGroup.POST("/register", customersController.Register())
		customersGroup.POST("/login", customersController.Login())
	}
	return nil
}
