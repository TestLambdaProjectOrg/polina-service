package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var appEnv = os.Getenv("APP_ENV")

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, param MyEvent) (string, error) {

	return fmt.Sprintf("Hello to Polina. You are in %s", appEnv), nil
}

func main() {
	lambda.Start(HandleRequest)
}
