package routes

import (
	"ecommerce/customer/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigRouteCustomers(rootGroup *gin.RouterGroup, dbCon *mongo.Database) error {
	customerController, err := controllers.NewCustomerController(dbCon)
	if err != nil {
		return err
	}
	customersGroup := rootGroup.Group("/customers")
	{
		customersGroup.POST("/register", customerController.Register())
		customersGroup.POST("/login", customerController.Login())
	}
	return nil
}
