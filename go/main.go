package main

import (
	"context"
	"go-handler/groups"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Groups string `json:"groups"`
}

type Response struct {
	Sizes string `json:"sizes"`
}

func HandleGroups(ctx context.Context, event Event) (Response, error) {
	var response Response
	response.Sizes = groups.DeterminePossibleSizes(event.Groups)
	return response, nil
}

func main() {
	lambda.Start(HandleGroups)
}
