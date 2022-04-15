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

func GetCachedReviews(ctx *gin.Context) {
	pid, err := primitive.ObjectIDFromHex(ctx.Param("pid"))
	if err != nil {
		panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid productId"))
	}
	rdb, err := db.GetRedisClient()
	if err != nil {
		panic(err)
	}
	redisReviews, err := rdb.HGet(context.Background(), "reviews", pid.Hex()).Result()
	if err == redis.Nil {
		ctx.Next()
		return
	}
	if err != nil {
		panic(err)
	}
	var reviews []*entities.Review
	err = json.Unmarshal([]byte(redisReviews), &reviews)
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  reviews,
	})
	ctx.Abort()
}
