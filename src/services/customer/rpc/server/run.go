package server

import (
	"ecommerce/customer/models"
	"ecommerce/customer/rpc/server/customer"
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

	customerModel, err := models.NewCustomer()
	if err != nil {
		return err
	}

	customerServer := &CustomerServer{
		customerModel: customerModel,
	}

	grpcServer := grpc.NewServer()

	customer.RegisterCustomerServer(grpcServer, customerServer)

	go func() {
		log.Printf("[gRPC]: gRPC Server is listening on %s \n", port)
		grpcServer.Serve(lis)
	}()

	return nil
}
