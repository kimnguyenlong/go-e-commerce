package server

import (
	context "context"
	"ecommerce/provider/models"
	"ecommerce/provider/rpc/server/provider"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProviderServer struct {
	provider.UnimplementedProviderServer
	providerModel *models.Provider
}

func (this *ProviderServer) IsExistingProvider(ctx context.Context, pI *provider.ProviderInfo) (*provider.Result, error) {
	pID, err := primitive.ObjectIDFromHex(pI.GetPID())

	if err != nil {
		return &provider.Result{
			IsExisting: false,
		}, nil
	}

	result := this.providerModel.Collection.FindOne(context.Background(), bson.M{"_id": pID})
	if result.Err() != nil {
		return &provider.Result{
			IsExisting: false,
		}, nil
	}
	return &provider.Result{
		IsExisting: true,
	}, nil
}
