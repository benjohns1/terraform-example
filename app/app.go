package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest handles an API Gateway proxy request
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Capture the 'name' query parameter from the request
	var name string
	var ok bool
	if name, ok = request.QueryStringParameters["name"]; !ok {
		name = "my pretties"
	}

	// Build the response
	message := struct {
		Message string `json:"message"`
		SampleEnvVar string `json:"sampleEnvVar"`
	} {
		Message: fmt.Sprintf("Hello, %s!", name),
		SampleEnvVar: os.Getenv("SampleEnvVar"),
	}

	// Encode the response data into a JSON string
	messageJSON, _ := json.Marshal(message)

	// Return response to the API gateway
	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           map[string]string{"Content-Type": "application/json"},
		Body:              string(messageJSON),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
