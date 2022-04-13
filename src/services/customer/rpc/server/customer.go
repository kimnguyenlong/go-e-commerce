package server

import (
	"context"
	"ecommerce/customer/models"
	"ecommerce/customer/rpc/server/customer"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerServer struct {
	customer.UnimplementedCustomerServer
	customerModel *models.Customer
}

func (s CustomerServer) IsExistingCustomer(ctx context.Context, customerInfo *customer.CustomerInfo) (*customer.Result, error) {
	log.Printf("[gRPC]: execute IsExistingCustomer\n")
	cid, err := primitive.ObjectIDFromHex(customerInfo.GetId())
	if err != nil {
		return &customer.Result{
			IsExisting: false,
		}, nil
	}
	err = s.customerModel.Collection.FindOne(context.Background(), bson.M{"_id": cid}).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &customer.Result{
				IsExisting: false,
			}, nil
		}
		return nil, err
	}
	return &customer.Result{
		IsExisting: true,
	}, nil
}
