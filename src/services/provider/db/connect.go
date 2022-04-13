package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbCon *mongo.Database = nil

func GetDBCon() (*mongo.Database, error) {
	if dbCon != nil {
		return dbCon, nil
	}
	dbCon, err := connect(os.Getenv("MONGO_CONNECT_URI"))
	return dbCon, err
}

func connect(uri string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database("e-commerce"), nil
}
