package priceUpdateRepository

//go:generate mockgen -source=productInfoUpdateRepo.go -destination=./mocks/mock_productInfoUpdateRepo.go

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"product-information-updater/app/dbUtils"
	updatePriceModel "product-information-updater/app/updatePrice/models"
	"time"
)

type ProductUpdateInfoRepo interface {
	Save(payload updatePriceModel.ProductEvent) error
}

type productUpdateInfoRepo struct {
	mongoCollection dbUtils.MongoCollection
}

func NewProductUpdateInfoRepo(mongoCollection dbUtils.MongoCollection) ProductUpdateInfoRepo {
	return productUpdateInfoRepo{
		mongoCollection: mongoCollection,
	}
}

func (r productUpdateInfoRepo) Save(payload updatePriceModel.ProductEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	insert := bson.M{
		"$set": bson.M{
			"value": payload,
		},
	}

	_, err := r.mongoCollection.InsertOne(ctx, insert)
	if err != nil {
		log.Printf("Failed to insert document into MongoDB: %v", err)
		return err
	}

	return nil
}
