package controllers

import (
	"context"
	"ecommerce/cart/customerror"
	"ecommerce/cart/entities"
	"ecommerce/cart/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartController struct {
	cartModel *models.Cart
}

func NewCartController() (*CartController, error) {
	cartModel, err := models.NewCart()
	if err != nil {
		return nil, err
	}
	return &CartController{
		cartModel: cartModel,
	}, nil
}

func (c CartController) SetCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cart *entities.Cart
		err := ctx.BindJSON(&cart)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		customerId := ctx.GetString("uid")
		cart.CustomerId = customerId
		filter := bson.M{"customerId": customerId}
		opts := options.FindOneAndReplace().SetReturnDocument(options.After).SetUpsert(true)
		err = c.cartModel.Collection.FindOneAndReplace(context.Background(), filter, cart, opts).Decode(&cart)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  cart,
		})
	}
}

func (c CartController) GetCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cartId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid cartId"))
		}
		customerId := ctx.GetString("uid")
		filter := bson.M{"_id": cartId, "customerId": customerId}
		var cart *entities.Cart
		err = c.cartModel.Collection.FindOne(context.Background(), filter).Decode(&cart)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("No cart with id %s", cartId.Hex())))
			}
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  cart,
		})
	}
}
