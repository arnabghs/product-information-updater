package priceUpdateRepository

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"product-information-updater/app/dbUtils/mocks"
	updatePriceModel "product-information-updater/app/updatePrice/models"
	"testing"
)

var sampleProductEvent = updatePriceModel.ProductEvent{
	ID:        "123",
	Message:   "Test product update",
	ProductID: "abc12",
}

func Test_Save(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockMongoColl := mocks.NewMockMongoCollection(mockCtrl)

	repo := NewProductUpdateInfoRepo(mockMongoColl)

	testcases := []struct {
		name        string
		dbResponse  interface{}
		dbError     error
		expectError bool
	}{
		{
			name:        "save is successful",
			dbResponse:  &mongo.InsertOneResult{InsertedID: "some-id"},
			dbError:     nil,
			expectError: false,
		},
		{
			name:        "save failed",
			dbResponse:  &mongo.InsertOneResult{},
			dbError:     errors.New("db down"),
			expectError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			insert := bson.M{
				"$set": bson.M{
					"value": sampleProductEvent,
				},
			}
			mockMongoColl.EXPECT().InsertOne(gomock.Any(), insert).Return(tc.dbResponse, tc.dbError)

			err := repo.Save(sampleProductEvent)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
