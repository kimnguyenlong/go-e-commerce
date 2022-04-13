package models

import "go.mongodb.org/mongo-driver/bson"

type Review struct {
	*Base
}

var reviewSchema = bson.M{
	"bsonType":             "object",
	"title":                "Product schema",
	"required":             []string{"content", "customerId", "productId"},
	"additionalProperties": false,
	"properties": bson.M{
		"_id": bson.M{
			"bsonType": "objectId",
		},
		"content": bson.M{
			"bsonType":  "string",
			"maxLength": 1000,
		},
		"customerId": bson.M{
			"bsonType": "string",
		},
		"productId": bson.M{
			"bsonType": "string",
		},
	},
}

var review *Review = nil

func NewReview() (*Review, error) {
	if review != nil {
		return review, nil
	}
	base, err := NewBase("reviews", reviewSchema)
	if err != nil {
		return nil, err
	}
	return &Review{
		Base: base,
	}, nil
}
