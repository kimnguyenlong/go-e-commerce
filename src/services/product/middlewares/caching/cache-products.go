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

const PRODUCTS_TTL = time.Hour * 1

func CacheProducts(ctx *gin.Context) {
	products, exist := ctx.Get("products")
	if !exist {
		log.Println("Internal error: Error when caching products: key products did not exist in context")
		return
	}
	_, ok := products.([]*entities.Product)
	if !ok {
		log.Println("Internal error: Error when caching products: cannot cast products to []*entity.Product")
		return
	}
	rdb, err := db.GetRedisClient()
	if err != nil {
		log.Printf("Internal error: Error when caching products: %s", err)
		return
	}
	productsMap := map[string]string{}
	for _, p := range products.([]*entities.Product) {
		productJson, err := json.Marshal(p)
		if err != nil {
			continue
		}
		productsMap[p.ID] = string(productJson)
	}
	err = rdb.HSet(context.Background(), "products", productsMap).Err()
	if err != nil {
		log.Printf("Internal error: Error when caching products: %s", err)
		return
	}
	err = rdb.Expire(context.Background(), "products", PRODUCTS_TTL).Err()
	if err != nil {
		log.Printf("Internal error: Error when caching products: %s", err)
		return
	}
}
