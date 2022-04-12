package main

import (
	"context"
	"ecommerce/provider/db"
	"ecommerce/provider/middlewares"
	"ecommerce/provider/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var port = ":9001"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load the .env file: %s\n", err)
	}

	dbCon, err := db.Connect(os.Getenv("MONGO_CONNECT_URI"))
	if err != nil {
		log.Fatalf("Cannot connect to the Database: %s\n", err)
	}
	defer dbCon.Client().Disconnect(context.Background())

	router := gin.Default()

	router.Use(middlewares.CrossOriginResource())

	router.Use(middlewares.ErrorHandler)

	rootGroup := router.Group("/api")

	err = routes.ConfigRouteProviders(rootGroup, dbCon)
	if err != nil {
		log.Fatalf("Cannot config route providers: %s", err)
	}

	err = router.Run(port)
	if err != nil {
		log.Fatalf("Cannot run server: %s", err)
	}
}
