package caching

import (
	"context"
	"ecommerce/product/db"
	"ecommerce/product/entities"
	"encoding/json"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const REVIEWS_TTL = time.Hour * 1

func CacheReviews(ctx *gin.Context) {
	productId, exist := ctx.Get("productId")
	if !exist {
		log.Println("Internal error: Error when caching reviews: key productId did not exist in context")
		return
	}
	reviews, exist := ctx.Get("reviews")
	if !exist {
		log.Println("Internal error: Error when caching reviews: key reviews did not exist in context")
		return
	}
	_, ok := reviews.([]*entities.Review)
	if !ok {
		log.Println("Internal error: Error when caching reviews: cannot cast to []*entity.Review")
		return
	}
	rdb, err := db.GetRedisClient()
	if err != nil {
		log.Printf("Internal error: Error when caching reviews: %s", err)
		return
	}
	reviewsJSON, err := json.Marshal(reviews.([]*entities.Review))
	if err != nil {
		log.Printf("Internal error: Error when caching reviews: %s", err)
		return
	}
	err = rdb.HSet(context.Background(), "reviews", productId, string(reviewsJSON)).Err()
	if err != nil {
		log.Printf("Internal error: Error when caching reviews: %s", err)
		return
	}
	err = rdb.Expire(context.Background(), "reviews", REVIEWS_TTL).Err()
	if err != nil {
		log.Printf("Internal error: Error when caching reviews: %s", err)
		return
	}
}
