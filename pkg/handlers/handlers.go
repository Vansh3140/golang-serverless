package handlers

import (
	"github.com/Vansh3140/golang-serverless/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"net/http"
)

// ErrorMethodNotAllowed is the response message for unsupported HTTP methods
var ErrorMethodNotAllowed = "method not allowed"

// ErrorBody represents the structure for error responses
type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"` // Error message in the response body
}

// GetUser handles GET requests to fetch a user by email or all users.
// If the "email" query parameter is provided, it fetches a specific user; otherwise, it fetches all users.
//
// Parameters:
// - req: APIGatewayProxyRequest containing the request data.
// - tableName: DynamoDB table name where user data is stored.
// - dynaClient: DynamoDB client interface.
//
// Returns:
// - APIGatewayProxyResponse with user data or error message.
func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]

	// Fetch a specific user if the "email" query parameter is provided
	if len(email) > 0 {
		result, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}

	// Fetch all users if no "email" query parameter is provided
	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

// CreateUser handles POST requests to create a new user in DynamoDB.
//
// Parameters:
// - req: APIGatewayProxyRequest containing the user data.
// - tableName: DynamoDB table name where the new user will be stored.
// - dynaClient: DynamoDB client interface.
//
// Returns:
// - APIGatewayProxyResponse with the created user data or error message.
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)
}

// UpdateUser handles PUT requests to update existing user data in DynamoDB.
//
// Parameters:
// - req: APIGatewayProxyRequest containing the updated user data.
// - tableName: DynamoDB table name where the user data is stored.
// - dynaClient: DynamoDB client interface.
//
// Returns:
// - APIGatewayProxyResponse with the updated user data or error message.
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

// DeleteUser handles DELETE requests to remove a user from DynamoDB.
//
// Parameters:
// - req: APIGatewayProxyRequest containing the user ID.
// - tableName: DynamoDB table name where the user data is stored.
// - dynaClient: DynamoDB client interface.
//
// Returns:
// - APIGatewayProxyResponse with a success message or error message.
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, "User deleted successfully")
}

// UnhandledMethod handles unsupported HTTP methods and returns a 405 Method Not Allowed response.
//
// Returns:
// - APIGatewayProxyResponse with a "method not allowed" error message.
func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
