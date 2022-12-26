package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type Todo struct {
	Id          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Description string `json:"description" dynamodbav:"description"`
	Status      bool   `json:"status" dynamodbav:"status"`
}

func main() {
	lambda.Start(router)
}
