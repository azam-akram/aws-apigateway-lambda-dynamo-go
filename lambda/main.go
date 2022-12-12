package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	msg := fmt.Sprintf("Hello %s!", name.Name)
	fmt.Println(msg)

	return msg, nil
}

func main() {
	lambda.Start(HandleRequest)
}
