package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

type Product struct {
	*Base
}

var product *Product = nil

var productSchema = bson.M{
	"bsonType":             "object",
	"title":                "Product schema",
	"required":             []string{"title", "price"},
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"title": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
		},
		"price": bson.M{
			"bsonType": "number",
			"minimum":  0,
		},
		"categories": bson.M{
			"bsonType": "array",
			"items": bson.M{
				"bsonType":  "string",
				"maxLength": 256,
			},
			"minItems":    1,
			"uniqueItems": true,
		},
		"description": bson.M{
			"bsonType":  "string",
			"maxLength": 1000,
		},
		"providerId": bson.M{
			"bsonType": "string",
		},
	},
}

func NewProduct() (*Product, error) {
	if product != nil {
		return product, nil
	}
	base, err := NewBase("products", productSchema)
	if err != nil {
		return nil, err
	}
	product = &Product{
		Base: base,
	}
	return product, nil
}
