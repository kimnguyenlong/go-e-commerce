package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var customerSchema = bson.M{
	"bsonType":             "object",
	"title":                "Customer Schema",
	"required":             []string{"email", "password", "full_name", "address"},
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"email": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
			"pattern":   "[a-z0-9]+@[a-z0-9]+",
		},
		"password": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
		},
		"full_name": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
		},
		"phone": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
			"pattern":   "[0-9]+",
		},
		"address": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
		},
	},
}

var customer *Customer = nil

type Customer struct {
	*Base
}

func NewCustomer(dbCon *mongo.Database) (*Customer, error) {
	if customer != nil {
		return customer, nil
	}
	base, err := NewBase(dbCon, "customers", customerSchema)
	if err != nil {
		return nil, err
	}
	customer = &Customer{
		Base: base,
	}
	return customer, nil
}
