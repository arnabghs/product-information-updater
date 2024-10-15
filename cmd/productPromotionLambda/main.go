package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// This is a scratch file to simulate as a lambda
// Hence all the secrets and values are hardcoded
func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.EuCentral1RegionID),
		Endpoint:    aws.String("http://localhost:4566"),
		DisableSSL:  aws.Bool(true),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", ""),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	svc := sqs.New(sess)
	queueURL := "http://localhost:4566/000000000000/dummy-queue"

	// Polling loop
	for {
		log.Printf("Getting data from queue...")
		result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(5),
			WaitTimeSeconds:     aws.Int64(5),
		})
		if err != nil {
			log.Fatalf("Error receiving messages: %v", err)
		}

		if len(result.Messages) == 0 {
			fmt.Println("No messages received, continuing polling...")
		} else {
			for _, message := range result.Messages {
				fmt.Printf("Message ID: %s\n", *message.MessageId)
				fmt.Printf("Message Body: %s\n", *message.Body)

				// delete the message after reading it
				_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
					QueueUrl:      aws.String(queueURL),
					ReceiptHandle: message.ReceiptHandle,
				})
				if err != nil {
					log.Fatalf("Failed to delete message: %v", err)
				}
				fmt.Println("Message deleted successfully")
			}
		}

		// Optional: Sleep for a short duration before polling again
		time.Sleep(10 * time.Second)
	}
}
