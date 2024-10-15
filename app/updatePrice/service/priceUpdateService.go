package priceUpdateService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
	"product-information-updater/app/updatePrice/models"
	"product-information-updater/app/updatePrice/repository"
)

type Service interface {
	Process(request updatePriceModel.RequestBody) error
	saveToDb(request updatePriceModel.RequestBody) error
	publishToSNS(message updatePriceModel.RequestBody) error
}

type service struct {
	productRepo priceUpdateRepository.ProductUpdateInfoRepo
	snsTopicARN string
	snsSession  *sns.SNS
}

func NewPriceUpdateService(productRepo priceUpdateRepository.ProductUpdateInfoRepo, snsTopicARN string, snsSession *sns.SNS) Service {
	return &service{
		productRepo: productRepo,
		snsTopicARN: snsTopicARN,
		snsSession:  snsSession,
	}
}

func (s service) Process(request updatePriceModel.RequestBody) error {
	// map to DB schema
	document := request

	saveErr := s.saveToDb(document)
	if saveErr != nil {
		log.Printf("Failed to save to DB: %v", saveErr)
		return saveErr
	}

	// use a go routine for async tasks
	publishErr := s.publishToSNS(document)
	if publishErr != nil {
		log.Printf("Failed to publish to SNS: %v", publishErr)
		return publishErr
	}

	return nil
}

func (s service) saveToDb(request updatePriceModel.RequestBody) error {
	// mapper
	paylaod := updatePriceModel.ProductEvent{
		ID:      request.ID,
		Message: request.Message,
	}

	return s.productRepo.Save(paylaod)
}

func (s service) publishToSNS(message updatePriceModel.RequestBody) error {
	msg := fmt.Sprintf("ID: %s, Message: %s", message.ID, message.Message)
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
