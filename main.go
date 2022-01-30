package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/support"
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	client := support.NewFromConfig(cfg)

	output, err := client.DescribeCases(context.TODO(), &support.DescribeCasesInput{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
