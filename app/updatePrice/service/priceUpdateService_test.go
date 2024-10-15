package priceUpdateService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	mock_queueUtils "product-information-updater/app/queueUtils/mocks"
	updatePriceModel "product-information-updater/app/updatePrice/models"
	mock_priceUpdateRepository "product-information-updater/app/updatePrice/repository/mocks"
	"testing"
)

//TODO: add tests for error scenarios

func Test_Process(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockRepo := mock_priceUpdateRepository.NewMockProductUpdateInfoRepo(mockCtrl)
	mockSNS := mock_queueUtils.NewMockSNSSession(mockCtrl)
	snsTopicArn := "test-topic"
	prodID := "123"
	request := updatePriceModel.RequestBody{
		ID:      "12",
		Message: "Test-message",
	}
	recorder := httptest.NewRecorder()
	gCtx, _ := gin.CreateTestContext(recorder)

	payload := updatePriceModel.ProductEvent{
		ID:        request.ID,
		Message:   request.Message,
		ProductID: prodID,
	}
	mockRepo.EXPECT().Save(payload).Return(nil)

	publishInput := &sns.PublishInput{
		Message:  aws.String(fmt.Sprintf("ID: %s, Message: %s, ProductID: %s", request.ID, request.Message, prodID)),
		TopicArn: aws.String(snsTopicArn),
	}
	mockSNS.EXPECT().Publish(publishInput).Return(nil, nil)

	svc := NewPriceUpdateService(mockRepo, snsTopicArn, mockSNS)

	err := svc.Process(gCtx, prodID, request)

	assert.NoError(t, err)
}
