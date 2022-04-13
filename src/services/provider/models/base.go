package models

import (
	"context"
	"ecommerce/provider/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Base struct {
	Collection *mongo.Collection
}

func NewBase(name string, schema interface{}) (*Base, error) {
	db, err := db.GetDBCon()
	if err != nil {
		return nil, err
	}
	validator := bson.M{
		"$jsonSchema": schema,
	}

	// update schema if the collection is existing
	updateValidatorCmd := bson.D{
		{Key: "collMod", Value: name},
		{Key: "validator", Value: validator},
	}
	err = db.RunCommand(context.Background(), updateValidatorCmd).Err()
	if err == nil {
		return &Base{
			Collection: db.Collection(name),
		}, nil
	}

	// create a new collection if the collection doesn't exist
	opts := options.CreateCollection().SetValidator(validator)
	err = db.CreateCollection(context.Background(), name, opts)
	if err != nil {
		return nil, err
	}
	return &Base{
		Collection: db.Collection(name),
	}, nil
}
