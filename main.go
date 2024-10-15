package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"net/http"
	"os"
	"product-information-updater/router"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// App holds the application dependencies
type App struct {
	MongoCollection *mongo.Collection
	SNSSession      *sns.SNS
	SNSTopicARN     string
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Proceeding with environment variables.")
	}

	// Initialize the application
	app, err := InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize the application: %v", err)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	appRouter := router.InitializeRouter(app.SNSTopicARN, app.SNSSession, app.MongoCollection)
	err = http.ListenAndServe(":"+port, appRouter)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

// InitializeApp sets up the application with all dependencies
func InitializeApp() (*App, error) {
	// Connect to MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is not set")
	}
	mongoDB := os.Getenv("MONGO_DB")
	if mongoDB == "" {
		return nil, fmt.Errorf("MONGO_DB is not set")
	}
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	if mongoCollection == "" {
		return nil, fmt.Errorf("MONGO_COLLECTION is not set")
	}
	mongoUsername := os.Getenv("MONGO_USERNAME")
	if mongoUsername == "" {
		return nil, fmt.Errorf("MONGO_USERNAME is not set")
	}
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	if mongoPassword == "" {
		return nil, fmt.Errorf("MONGO_PASSWORD is not set")
	}
	mongoAuthSource := os.Getenv("MONGO_AUTH_SOURCE")
	if mongoPassword == "" {
		return nil, fmt.Errorf("MONGO_AUTH_SOURCE is not set")
	}

	var cred options.Credential

	cred.AuthSource = mongoAuthSource
	cred.Username = mongoUsername
	cred.Password = mongoPassword

	clientOptions := options.Client().ApplyURI(mongoURI).SetAuth(cred)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping MongoDB to verify connection
	if err := mongoClient.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")

	// Initialize AWS SNS session
	awsRegion := os.Getenv("AWS_REGION")
	if awsRegion == "" {
		return nil, fmt.Errorf("AWS_REGION is not set")
	}
	awsAccessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	if awsAccessKeyId == "" {
		return nil, fmt.Errorf("AWS_ACCESS_KEY_ID is not set")
	}
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if awsSecretAccessKey == "" {
		return nil, fmt.Errorf("AWS_SECRET_ACCESS_KEY is not set")
	}

	snsEndpoint := "http://localhost:4566" // TODO

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, ""),

		Endpoint:   aws.String(snsEndpoint),
		DisableSSL: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %v", err)
	}

	snsClient := sns.New(sess)

	snsTopicARN := os.Getenv("AWS_SNS_TOPIC_ARN")
	if snsTopicARN == "" {
		return nil, fmt.Errorf("AWS_SNS_TOPIC_ARN is not set")
	}

	// Initialize the App struct
	app := &App{
		MongoCollection: mongoClient.Database(mongoDB).Collection(mongoCollection),
		SNSSession:      snsClient,
		SNSTopicARN:     snsTopicARN,
	}

	return app, nil
}
