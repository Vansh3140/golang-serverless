package user

import (
	"encoding/json"
	"errors"
	"github.com/Vansh3140/golang-serverless/pkg/validators"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Error messages for common issues
var (
	ErrorFailedToFetchRecord     = "failed to fetch record from dynamodb"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record fetched from dynamodb"
	ErrorInvalidUserData         = "invalid user data"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "couldn't marshal the item"
	ErrorCouldNotDeleteItem      = "couldn't delete the item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item"
	ErrorUserAlreadyExists       = "user already exists"
	ErrorUserDoesNotExist        = "user doesn't exist"
)

// User represents a user entity in the system
type User struct {
	Email     string `json:"email"`     // User's email address
	FirstName string `json:"firstname"` // User's first name
	LastName  string `json:"lastname"`  // User's last name
}

// FetchUser retrieves a user by email from DynamoDB.
//
// Parameters:
// - email: The email of the user to fetch.
// - tableName: The name of the DynamoDB table.
// - dynaClient: The DynamoDB client interface.
//
// Returns:
// - A pointer to the User struct containing user details.
// - An error if the user cannot be fetched or unmarshaled.
func FetchUser(email string, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	// Fetch the item from DynamoDB
	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	// Unmarshal the result into a User struct
	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

// FetchUsers retrieves all users from DynamoDB.
//
// Parameters:
// - tableName: The name of the DynamoDB table.
// - dynaClient: The DynamoDB client interface.
//
// Returns:
// - A pointer to a slice of User structs.
// - An error if the users cannot be fetched or unmarshaled.
func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Scan the table for all items
	result, err := dynaClient.Scan(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	// Unmarshal the result into a slice of User structs
	items := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, items)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}

	return items, nil
}

// CreateUser creates a new user in DynamoDB.
//
// Parameters:
// - req: APIGatewayProxyRequest containing the user data.
// - tableName: The name of the DynamoDB table.
// - dynaClient: The DynamoDB client interface.
//
// Returns:
// - A pointer to the newly created User struct.
// - An error if user creation fails.
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var newUser User

	// Unmarshal the request body into a User struct
	if err := json.Unmarshal([]byte(req.Body), &newUser); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	// Validate the user's email
	if !validators.IsEmailValid(newUser.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	// Check if the user already exists
	currUser, _ := FetchUser(newUser.Email, tableName, dynaClient)
	if currUser != nil && len(currUser.Email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	// Marshal the new user into a DynamoDB item
	result, err := dynamodbattribute.MarshalMap(newUser)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	// Insert the new user into the DynamoDB table
	input := &dynamodb.PutItemInput{
		Item:      result,
		TableName: aws.String(tableName),
	}

	_, putErr := dynaClient.PutItem(input)
	if putErr != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}

	return &newUser, nil
}

// UpdateUser updates an existing user in DynamoDB.
//
// Parameters:
// - req: APIGatewayProxyRequest containing the updated user data.
// - tableName: The name of the DynamoDB table.
// - dynaClient: The DynamoDB client interface.
//
// Returns:
// - A pointer to the updated User struct.
// - An error if the update fails.
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	var newUser User

	// Unmarshal the request body into a User struct
	if err := json.Unmarshal([]byte(req.Body), &newUser); err != nil {
		return nil, errors.New(ErrorInvalidEmail)
	}

	// Check if the user exists
	currUser, _ := FetchUser(newUser.Email, tableName, dynaClient)
	if currUser != nil && len(currUser.Email) == 0 {
		return nil, errors.New(ErrorUserDoesNotExist)
	}

	// Marshal the updated user into a DynamoDB item
	result, err := dynamodbattribute.MarshalMap(newUser)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	// Update the user in the DynamoDB table
	input := &dynamodb.PutItemInput{
		Item:      result,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}

	return &newUser, nil
}

// DeleteUser deletes a user from DynamoDB by email.
//
// Parameters:
// - req: APIGatewayProxyRequest containing the user's email in query parameters.
// - tableName: The name of the DynamoDB table.
// - dynaClient: The DynamoDB client interface.
//
// Returns:
// - An error if the user could not be deleted.
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) error {
	email := req.QueryStringParameters["email"]

	// Prepare the delete item input
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	// Delete the item from DynamoDB
	_, err := dynaClient.DeleteItem(input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
