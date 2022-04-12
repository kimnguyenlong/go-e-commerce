package routes

import (
	"ecommerce/provider/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConfigRouteProviders(rootGroup *gin.RouterGroup, dbCon *mongo.Database) error {
	providerController, err := controllers.NewProviderController(dbCon)
	if err != nil {
		return err
	}
	providersGroup := rootGroup.Group("/providers")
	{
		providersGroup.POST("/register", providerController.Register())
		providersGroup.POST("/login", providerController.Login())
	}
	return nil
}
