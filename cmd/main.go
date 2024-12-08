package main

import (
	"github.com/Vansh3140/golang-serverless/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
)

// Global DynamoDB client interface
var (
	dynaClient dynamodbiface.DynamoDBAPI
)

// main function initializes the AWS session, DynamoDB client, and starts the Lambda function handler.
func main() {
	// Get AWS region from the environment variable
	region := os.Getenv("AWS_REGION")

	// Create a new AWS session
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)}, // AWS region for the session
	)
	if err != nil {
		// Exit if the session cannot be created
		return
	}

	// Initialize the DynamoDB client using the session
	dynaClient = dynamodb.New(awsSession)

	// Start the Lambda function and set the handler
	lambda.Start(handler)
}

// tableName stores the DynamoDB table name from the environment variable
var tableName = os.Getenv("TABLE_NAME")

// handler processes incoming API Gateway requests and routes them to appropriate handler functions.
// It supports CRUD operations for user management.
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Route the request based on HTTP method
	switch req.HTTPMethod {
	case "GET":
		// Handle GET requests to fetch user data
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		// Handle POST requests to create a new user
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PUT":
		// Handle PUT requests to update existing user data
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		// Handle DELETE requests to remove a user
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		// Handle unsupported HTTP methods
		return handlers.UnhandledMethod()
	}
}
