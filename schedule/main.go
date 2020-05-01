package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/briantbates/go-lambda/schedule/helpers"
)

// Response type
type Response events.APIGatewayProxyResponse

// Request type
type Request struct {
	Name string `json:"name"`
}

// ErrorResponse struct
type ErrorResponse struct {
	Message string `json:"message"`
}

// Handler func
func Handler(_ context.Context, request events.APIGatewayProxyRequest) (Response, error) {

	err := helpers.ParseAndCheckBody(request.Body)

	if err != nil {

		message, _ := json.Marshal(map[string]interface{}{
			"message": "provided body is invalid",
		})

		return Response{StatusCode: 422, Body: string(message)}, nil
	}

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	placeholderResponse, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(placeholderResponse))

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(request.Body),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
