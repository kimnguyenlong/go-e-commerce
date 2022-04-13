package client

import (
	"context"
	"ecommerce/product/rpc/client/provider"

	"os"

	"google.golang.org/grpc"
)

type ProviderClient struct {
	client provider.ProviderClient
}

var providerClient *ProviderClient = nil

func GetProviderClient() (*ProviderClient, error) {
	if providerClient != nil {
		return providerClient, nil
	}
	conn, err := grpc.Dial(os.Getenv("PROVIDER_GRPC_SERVER_ADDR"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := provider.NewProviderClient(conn)
	return &ProviderClient{
		client: client,
	}, nil
}

func (pC ProviderClient) RunIsExistingProvider(providerInfo *provider.ProviderInfo) (bool, error) {
	result, err := pC.client.IsExistingProvider(context.Background(), providerInfo)
	if err != nil {
		return false, err
	}
	return result.IsExisting, nil
}
