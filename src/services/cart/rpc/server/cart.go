package server

import (
	"context"
	"ecommerce/cart/models"
	"ecommerce/cart/rpc/server/cart"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CartServer struct {
	cart.UnimplementedCartServer
	cartModel *models.Cart
}

func (s CartServer) GetSingleCart(ctx context.Context, cartInfo *cart.CartInfo) (*cart.CartData, error) {
	log.Println("[gRPC]: Execute GetSingleCart")
	var cart *cart.CartData
	filter := bson.M{"customerId": cartInfo.GetCustomerId()}
	err := s.cartModel.Collection.FindOne(context.Background(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("No cart found with customerId: %s", cartInfo.CustomerId))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return cart, nil
}

func (s CartServer) ClearCart(ctx context.Context, cartInfo *cart.CartInfo) (*cart.CartData, error) {
	log.Println("[gRPC]: Execute ClearCart")
	var cart *cart.CartData
	filter := bson.M{"customerId": cartInfo.GetCustomerId()}
	err := s.cartModel.Collection.FindOneAndDelete(context.Background(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, fmt.Sprintf("No cart found with customerId: %s", cartInfo.CustomerId))
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return cart, nil
}
