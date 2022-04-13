package main

import (
	"context"
	"ecommerce/provider/db"
	"ecommerce/provider/middlewares"
	"ecommerce/provider/routes"
	"log"
	"os"

	"ecommerce/provider/rpc/server"

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

	router := gin.Default()

	router.Use(middlewares.CrossOriginResource())

	router.Use(middlewares.ErrorHandler)

	rootGroup := router.Group("/api")

	err = routes.ConfigRouteProviders(rootGroup)
	if err != nil {
		log.Fatalf("Cannot config route providers: %s", err)
	}

	err = server.Run()
	if err != nil {
		log.Fatalf("Cannot run grpc server: %s", err)
	}

	err = router.Run(os.Getenv("HTTP_PORT"))
	if err != nil {
		log.Fatalf("Cannot run server: %s", err)
	}
}
