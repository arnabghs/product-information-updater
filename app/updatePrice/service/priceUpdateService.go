package priceUpdateService

//go:generate mockgen -source=priceUpdateService.go -destination=./mocks/mock_priceUpdateService.go

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
	"log"
	"product-information-updater/app/queueUtils"
	"product-information-updater/app/updatePrice/models"
	"product-information-updater/app/updatePrice/repository"
)

type Service interface {
	Process(ginCtx *gin.Context, productID string, request updatePriceModel.RequestBody) error
	SaveToDb(request updatePriceModel.ProductEvent) error
	PublishToSNS(message updatePriceModel.ProductEvent) error
}

type service struct {
	productRepo priceUpdateRepository.ProductUpdateInfoRepo
	snsTopicARN string
	snsSession  queueUtils.SNSSession
}

func NewPriceUpdateService(productRepo priceUpdateRepository.ProductUpdateInfoRepo, snsTopicARN string, snsSession queueUtils.SNSSession) Service {
	return &service{
		productRepo: productRepo,
		snsTopicARN: snsTopicARN,
		snsSession:  snsSession,
	}
}

func (s service) Process(ginCtx *gin.Context, productID string, request updatePriceModel.RequestBody) error {
	// use this context for creating logger, tracing spans etc

	// mapper
	paylaod := updatePriceModel.ProductEvent{
		ID:        request.ID,
		Message:   request.Message,
		ProductID: productID,
	}

	saveErr := s.SaveToDb(paylaod)
	if saveErr != nil {
		log.Printf("Failed to save to DB: %v", saveErr)
		return saveErr
	}

	publishErr := s.PublishToSNS(paylaod)
	if publishErr != nil {
		log.Printf("Failed to publish to SNS: %v", publishErr)
		return publishErr
	}

	return nil
}

func (s service) SaveToDb(payload updatePriceModel.ProductEvent) error {
	return s.productRepo.Save(payload)
}

func (s service) PublishToSNS(message updatePriceModel.ProductEvent) error {
	msg := fmt.Sprintf("ID: %s, Message: %s, ProductID: %s", message.ID, message.Message, message.ProductID)
	publishInput := &sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: aws.String(s.snsTopicARN),
	}

	_, err := s.snsSession.Publish(publishInput)
	if err != nil {
		log.Printf("Failed to publish message to SNS: %v", err)
		return err
	}

	return nil
}
