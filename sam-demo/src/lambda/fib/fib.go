package main

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request.QueryStringParameters["n"])
	n, err := strconv.Atoi(request.QueryStringParameters["n"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Input error",
			StatusCode: 500,
		}, err
	}

	message := fmt.Sprintf("Fib sequence number %d: %d", n, Fib(n))
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
