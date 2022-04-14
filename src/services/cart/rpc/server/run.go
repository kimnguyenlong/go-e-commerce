package server

import (
	"ecommerce/cart/models"
	"ecommerce/cart/rpc/server/cart"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func Run() error {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	cartModel, err := models.NewCart()
	if err != nil {
		return err
	}

	cartServer := &CartServer{
		cartModel: cartModel,
	}

	grpcServer := grpc.NewServer()

	cart.RegisterCartServer(grpcServer, cartServer)

	go func() {
		log.Printf("[gRPC]: gRPC Server is listening on %s \n", port)
		grpcServer.Serve(lis)
	}()

	return nil
}
