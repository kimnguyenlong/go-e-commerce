package caching

import (
	"context"
	"ecommerce/product/customerror"
	"ecommerce/product/db"
	"ecommerce/product/entities"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCachedProducts(ctx *gin.Context) {
	rdb, err := db.GetRedisClient()
	if err != nil {
		panic(err)
	}
	redisProducts, err := rdb.HGetAll(context.Background(), "products").Result()
	if len(redisProducts) == 0 {
		ctx.Next()
		return
	}
	if err != nil {
		panic(err)
	}
	products := []*entities.Product{}
	for _, val := range redisProducts {
		var product *entities.Product
		err := json.Unmarshal([]byte(val), &product)
		if err != nil {
			continue
		}
		products = append(products, product)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  products,
	})
	ctx.Abort()
}

func GetSingleCachedProduct(ctx *gin.Context) {
	rdb, err := db.GetRedisClient()
	if err != nil {
		panic(err)
	}
	pid, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
	if err != nil {
		panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
	}
	redisProduct, err := rdb.HGet(context.Background(), "products", pid.Hex()).Result()
	if err == redis.Nil {
		ctx.Next()
		return
	}
	if err != nil {
		panic(err)
	}
	var product *entities.Product
	err = json.Unmarshal([]byte(redisProduct), &product)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  product,
	})
	ctx.Abort()
}
