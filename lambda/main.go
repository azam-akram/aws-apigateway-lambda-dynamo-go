package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyBook struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

func HandleRequest(ctx context.Context, book MyBook) (string, error) {
	msg := fmt.Sprintf("ID: %s, Title: %s", book.ID, book.Title)
	fmt.Println(msg)

	return msg, nil
}

func main() {
	lambda.Start(HandleRequest)
}
