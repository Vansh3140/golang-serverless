package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

// apiResponse generates a standardized API Gateway Proxy Response.
// It accepts a status code and a response body, formats them into an APIGatewayProxyResponse,
// and sets the "Content-Type" header to "application/json".
//
// Parameters:
// - status: HTTP status code (e.g., 200, 400, 500).
// - body: Response body, which can be any type (usually a struct or map).
//
// Returns:
// - A pointer to an APIGatewayProxyResponse containing the status code, headers, and JSON-encoded body.
// - An error (always nil in this function as errors are not handled here).
func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	// Initialize response with JSON content-type header
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	// Marshal the response body into a JSON string
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)

	// Return the response and nil for error (no error-handling logic here)
	return &resp, nil
}
