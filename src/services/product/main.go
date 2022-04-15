package main

import (
	"context"
	"ecommerce/product/db"
	"ecommerce/product/middlewares"
	"ecommerce/product/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load the .env file: %s\n", err)
	}

	dbCon, err := db.GetDBCon()
	if err != nil {
		log.Fatalf("Cannot connect to the Database: %s\n", err)
	}
	defer dbCon.Client().Disconnect(context.Background())

	rdb, err := db.GetRedisClient()
	if err != nil {
		log.Fatalf("Cannot connect to the Redis: %s\n", err)
	}
	defer rdb.Close()

	router := gin.Default()

	router.Use(middlewares.CrossOriginResource())

	router.Use(middlewares.ErrorHandler)

	rootGroup := router.Group("/api")

	err = routes.ConfigRouteProducts(rootGroup)
	if err != nil {
		log.Fatalf("Cannot config route products: %s\n", err)
	}

	err = router.Run(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatalf("Cannot run server: %s\n", err)
	}
}
