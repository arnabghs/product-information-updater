package queueUtils

//go:generate mockgen -source=sns.go -destination=./mocks/mock_sns.go

import (
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSSession interface {
	Publish(input *sns.PublishInput) (*sns.PublishOutput, error)
}

type snsSession struct {
	sns *sns.SNS
}

func NewSNSSession(sns *sns.SNS) SNSSession {
	return &snsSession{
		sns: sns,
	}
}

func (s *snsSession) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	return s.sns.Publish(input)
}
