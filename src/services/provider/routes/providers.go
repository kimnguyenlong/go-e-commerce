package routes

import (
	"ecommerce/provider/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRouteProviders(rootGroup *gin.RouterGroup) error {
	providersController, err := controllers.NewProvidersController()
	if err != nil {
		return err
	}
	providersGroup := rootGroup.Group("/providers")
	{
		providersGroup.POST("/register", providersController.Register())
		providersGroup.POST("/login", providersController.Login())
	}
	return nil
}
