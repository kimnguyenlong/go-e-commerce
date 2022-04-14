package models

import "go.mongodb.org/mongo-driver/bson"

type Cart struct {
	*Base
}

var cartSchema = bson.M{
	"bsonType":             "object",
	"title":                "Cart schema",
	"required":             []string{"customerId", "items"},
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
	},
}

var cart *Cart = nil

func NewCart() (*Cart, error) {
	if cart != nil {
		return cart, nil
	}
	base, err := NewBase("carts", cartSchema)
	if err != nil {
		return nil, err
	}
	return &Cart{
		Base: base,
	}, nil
}
