package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/support"
)

type Case struct {
	Subject    string `json:"subject"`
	Status     string `json:"status"`
	SubmitteBy string `json:"submitteBy"`
	Url        string `json:"url"`
}

func NewDescribeCasesInput(aftertime, beforetime, language string) *support.DescribeCasesInput {
	return &support.DescribeCasesInput{
		AfterTime:            &aftertime,
		BeforeTime:           &beforetime,
		Language:             &language,
		IncludeResolvedCases: true,
	}
}

func GetCases(input *support.DescribeCasesInput) *[]Case {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}

	client := support.NewFromConfig(cfg)

	output, err := client.DescribeCases(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}

	cases := ArrangeCases(output)
	return cases
}

func ArrangeCases(output *support.DescribeCasesOutput) *[]Case {
	cases := []Case{}

	outputCases := output.Cases
	for _, c := range outputCases {
		eachCase := Case{
			Subject:    *c.Subject,
			Status:     *c.Status,
			SubmitteBy: *c.SubmittedBy,
			Url:        "",
		}
		cases = append(cases, eachCase)
	}
	return &cases
}
