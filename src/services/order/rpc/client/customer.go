package client

import (
	"context"
	"ecommerce/order/rpc/client/customer"

	"os"

	"google.golang.org/grpc"
)

type CustomerClient struct {
	client customer.CustomerClient
}

var customerClient *CustomerClient = nil

func GetCustomerClient() (*CustomerClient, error) {
	if customerClient != nil {
		return customerClient, nil
	}
	conn, err := grpc.Dial(os.Getenv("CUSTOMER_GRPC_SERVER_ADDR"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := customer.NewCustomerClient(conn)
	return &CustomerClient{
		client: client,
	}, nil
}

func (pC CustomerClient) RunIsExistingCustomer(customerInfo *customer.CustomerInfo) (bool, error) {
	result, err := pC.client.IsExistingCustomer(context.Background(), customerInfo)
	if err != nil {
		return false, err
	}
	return result.IsExisting, nil
}
