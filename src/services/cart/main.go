package main

import (
	"context"
	"ecommerce/cart/db"
	"ecommerce/cart/middlewares"
	"ecommerce/cart/rpc/server"
	"log"
	"os"

	"ecommerce/cart/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Cannot load the .env file: %s\n", err)
	// }

	dbCon, err := db.GetDBCon()
	if err != nil {
		log.Fatalf("Cannot connect to the Database: %s\n", err)
	}
	defer dbCon.Client().Disconnect(context.Background())

	router := gin.Default()

	router.Use(middlewares.CrossOriginResource())

	router.Use(middlewares.ErrorHandler)

	rootGroup := router.Group("/api")

	err = routes.ConfigRouteCarts(rootGroup)
	if err != nil {
		log.Fatalf("Cannot config route carts: %s", err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalf("Cannot run the gRPC server: %s", err)
	}

	err = router.Run(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatalf("Cannot run server: %s", err)
	}
}
