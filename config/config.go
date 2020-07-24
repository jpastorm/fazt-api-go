package config

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CUser(db *mongo.Database, ctx context.Context) {

	index := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}
	if _, err := db.Collection("user").Indexes().CreateOne(ctx, index); err != nil {
		panic("Could not create index")
	}
}
