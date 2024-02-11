package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func process(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request.Body
	reqJson, err := json.Marshal(req)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error",
			StatusCode: 404,
		}, err
	}
	fmt.Println(reqJson)
	return events.APIGatewayProxyResponse{
		Body:       "Hello World",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(process)
}
