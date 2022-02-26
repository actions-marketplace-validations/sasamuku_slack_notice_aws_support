package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/support"
)

type Case struct {
	Subject     string `json:"subject"`
	Status      string `json:"status"`
	SubmitteBy  string `json:"submitteBy"`
	TimeCreated string `json:"timeCreated"`
	Url         string `json:"url"`
}

func NewDescribeCasesInput(aftertime, beforetime, language string, include bool) *support.DescribeCasesInput {
	return &support.DescribeCasesInput{
		AfterTime:            &aftertime,
		BeforeTime:           &beforetime,
		Language:             &language,
		IncludeResolvedCases: include,
	}
}

func GetCases(input *support.DescribeCasesInput) []*Case {
	client := loadConfig()
	output := outputCases(client, input)
	cases := arrangeCases(output)
	return cases
}

func loadConfig() *support.Client {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	return support.NewFromConfig(cfg)
}

func outputCases(c *support.Client, i *support.DescribeCasesInput) *support.DescribeCasesOutput {
	output, err := c.DescribeCases(context.TODO(), i)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func arrangeCases(output *support.DescribeCasesOutput) []*Case {
	var cases []*Case

	caseDetails := output.Cases
	for _, c := range caseDetails {
		eachCase := Case{
			Subject:     *c.Subject,
			Status:      *c.Status,
			SubmitteBy:  *c.SubmittedBy,
			TimeCreated: *c.TimeCreated,
			Url:         "https://console.aws.amazon.com/support/home#/case/?displayId=" + *c.DisplayId + "%26language=" + *c.Language,
		}
		cases = append(cases, &eachCase)
	}
	return cases
}
