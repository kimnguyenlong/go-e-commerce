package models

import "go.mongodb.org/mongo-driver/bson"

type Order struct {
	*Base
}

var orderSchema = bson.M{
	"bsonType":             "object",
	"title":                "Order schema",
	"required":             []string{"customerId", "items", "status", "created", "updated"},
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"customerId": bson.M{
			"bsonType": "string",
		},
		"items": bson.M{
			"bsonType": "array",
			"items": bson.M{
				"bsonType": "object",
				"properties": bson.M{
					"productId": bson.M{
						"bsonType": "string",
					},
					"quantity": bson.M{
						"bsonType": "int",
						"minimum":  1,
					},
				},
			},
		},
		"status": bson.M{
			"bsonType": "string",
		},
		"created": bson.M{
			"bsonType": "long",
		},
		"updated": bson.M{
			"bsonType": "long",
		},
	},
}

var order *Order = nil

func NewOrder() (*Order, error) {
	if order != nil {
		return order, nil
	}
	base, err := NewBase("orders", orderSchema)
	if err != nil {
		return nil, err
	}
	return &Order{
		Base: base,
	}, nil
}
