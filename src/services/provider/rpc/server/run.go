package server

import (
	"ecommerce/provider/models"
	"ecommerce/provider/rpc/server/provider"
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

	providerModel, err := models.NewProvider()
	if err != nil {
		return err
	}

	providerServer := &ProviderServer{
		providerModel: providerModel,
	}

	grpcServer := grpc.NewServer()

	provider.RegisterProviderServer(grpcServer, providerServer)

	go func() {
		log.Printf("gRPC Server is listening on %s \n", port)
		grpcServer.Serve(lis)
	}()

	return nil
}
