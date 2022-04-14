package client

import (
	"context"
	"ecommerce/order/rpc/client/cart"

	"os"

	"google.golang.org/grpc"
)

type CartClient struct {
	client cart.CartClient
}

var cartClient *CartClient = nil

func GetCartClient() (*CartClient, error) {
	if cartClient != nil {
		return cartClient, nil
	}
	conn, err := grpc.Dial(os.Getenv("CART_GRPC_SERVER_ADDR"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := cart.NewCartClient(conn)
	return &CartClient{
		client: client,
	}, nil
}

func (c CartClient) RunGetSingleCart(cartInfo *cart.CartInfo) (*cart.CartData, error) {
	return c.client.GetSingleCart(context.Background(), cartInfo)
}

func (c CartClient) RunClearCart(cartInfo *cart.CartInfo) (*cart.CartData, error) {
	return c.client.ClearCart(context.Background(), cartInfo)
}
