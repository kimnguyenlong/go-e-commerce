package routes

import (
	"ecommerce/product/controllers"
	"ecommerce/product/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRouteProducts(rootGroup *gin.RouterGroup) error {
	productsController, err := controllers.NewProductsController()
	reviewsController, err := controllers.NewReviewsController()
	if err != nil {
		return err
	}
	productsGroup := rootGroup.Group("/products")
	{
		productsGroup.POST("/", middlewares.Authenticate, middlewares.ConfirmProvider, productsController.CreateProduct())
		productsGroup.GET("/", productsController.GetProducts())

	}

	singleProductGroup := productsGroup.Group("/:pid")
	{
		singleProductGroup.GET("/", productsController.GetSingleProduct())
		singleProductGroup.DELETE("/", middlewares.Authenticate, middlewares.ConfirmProvider, productsController.DeleteProduct())
		singleProductGroup.PATCH("/", middlewares.Authenticate, middlewares.ConfirmProvider, productsController.UpdateProduct())
	}

	reviewsGroup := singleProductGroup.Group("/reviews")
	{
		reviewsGroup.POST("/", middlewares.Authenticate, middlewares.ConfirmCustomer, reviewsController.CreateReview())
		reviewsGroup.GET("/", reviewsController.GetReviews())
	}

	singleReviewGroup := reviewsGroup.Group("/:rid")
	{
		singleReviewGroup.DELETE("/", middlewares.Authenticate, middlewares.ConfirmCustomer, reviewsController.DeleteReview())
		singleReviewGroup.PATCH("/", middlewares.Authenticate, middlewares.ConfirmCustomer, reviewsController.UpdateReview())
	}
	return nil
}
