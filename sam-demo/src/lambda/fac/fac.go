package main

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Fac(n int) int {
	if n > 0 {
		return n * Fac(n-1)
	}
	return 1
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

	message := fmt.Sprintf("Fib sequence number %d: %d", n, Fac(n))
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
