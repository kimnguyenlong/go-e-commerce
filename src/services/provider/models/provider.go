package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var providerSchema = bson.M{
	"bsonType":             "object",
	"title":                "Provider Schema",
	"required":             []string{"email", "password", "full_name", "address", "company_name", "company_address", "url"},
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
		"company_name": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
		},
		"company_phone": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
			"pattern":   "[0-9]+",
		},
		"company_email": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
			"pattern":   "[a-z0-9]+@[a-z0-9]+",
		},
		"company_address": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
		},
		"url": bson.M{
			"bsonType":  "string",
			"maxLength": 256,
			"pattern":   "[a-z0-9]+",
		},
	},
}

var provider *Provider = nil

type Provider struct {
	*Base
}

func NewProvider(dbCon *mongo.Database) (*Provider, error) {
	if provider != nil {
		return provider, nil
	}
	base, err := NewBase(dbCon, "providers", providerSchema)
	if err != nil {
		return nil, err
	}
	provider = &Provider{
		Base: base,
	}
	return provider, nil
}
