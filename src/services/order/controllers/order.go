package controllers

import (
	"context"
	"ecommerce/order/customerror"
	"ecommerce/order/entities"
	"ecommerce/order/models"
	"ecommerce/order/rpc/client"
	"ecommerce/order/rpc/client/cart"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderController struct {
	orderModel *models.Order
}

func NewOrderController() (*OrderController, error) {
	orderModel, err := models.NewOrder()
	if err != nil {
		return nil, err
	}
	return &OrderController{
		orderModel: orderModel,
	}, nil
}

func (c OrderController) CreateOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId := ctx.GetString("uid")
		cartInfo := &cart.CartInfo{
			CustomerId: customerId,
		}
		cartClient, err := client.GetCartClient()
		if err != nil {
			panic(err)
		}
		cart, err := cartClient.RunGetSingleCart(cartInfo)
		if err != nil {
			if e, ok := status.FromError(err); ok {
				if e.Code() == codes.NotFound {
					panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("User with id %s dont have a cart yet", customerId)))
				}
			}
			panic(err)
		}
		items := make([]*entities.OrderItem, len(cart.Items))
		for i, item := range cart.Items {
			items[i] = &entities.OrderItem{
				ProductId: item.ProductId,
				Quantity:  int(item.Quantity),
			}
		}
		order := &entities.Order{
			CustomerId: customerId,
			Created:    time.Now().Unix(),
			Updated:    time.Now().Unix(),
			Status:     "PENDING",
			Items:      items,
		}
		result, err := c.orderModel.Collection.InsertOne(context.Background(), order)
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, err.Error()))
		}
		order.ID = result.InsertedID.(primitive.ObjectID).Hex()
		ctx.JSON(http.StatusCreated, gin.H{
			"error": false,
			"data":  order,
		})

		_, err = cartClient.RunClearCart(cartInfo)
		if err != nil {
			log.Printf("Internal error: %s\n", err)
		}
	}
}

func (c OrderController) GetOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customerId := ctx.GetString("uid")
		var orders []*entities.Order
		filter := bson.M{"customerId": customerId}
		curs, err := c.orderModel.Collection.Find(context.Background(), filter)
		if err != nil {
			panic(err)
		}
		defer curs.Close(context.Background())
		err = curs.All(context.Background(), &orders)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  orders,
		})
	}
}

func (c OrderController) GetSingleController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orderId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
		if err != nil {
			panic(customerror.NewAPIError(http.StatusBadRequest, "Invalid orderId"))
		}
		customerId := ctx.GetString("uid")
		var order *entities.Order
		filter := bson.M{
			"_id":        orderId,
			"customerId": customerId,
		}
		err = c.orderModel.Collection.FindOne(context.Background(), filter).Decode(&order)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				panic(customerror.NewAPIError(http.StatusNotFound, fmt.Sprintf("No order with id %s", orderId.Hex())))
			}
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error": false,
			"data":  order,
		})
	}
}
