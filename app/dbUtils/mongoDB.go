package dbUtils

//go:generate mockgen -source=mongoDB.go -destination=./mocks/mock_mongoDB.go --package=mocks

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCollection interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
}

type mongoCollection struct {
	coll *mongo.Collection
}

func NewMongoColl(coll *mongo.Collection) MongoCollection {
	return &mongoCollection{
		coll: coll,
	}
}

func (m *mongoCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	id, err := m.coll.InsertOne(ctx, document)
	return id, err
}
