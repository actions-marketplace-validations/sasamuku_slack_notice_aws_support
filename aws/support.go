package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
)

type Case struct {
	Subject     string `json:"subject"`
	Status      string `json:"status"`
	SubmittedBy string `json:"SubmittedBy"`
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

func GetCaseList(input *support.DescribeCasesInput) []*Case {
	client := loadConfig()
	output := outputCases(client, input)
	caseDetails := extractCaseDetails(output)
	caseList := makeCaseList(caseDetails)
	return caseList
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

func extractCaseDetails(o *support.DescribeCasesOutput) []types.CaseDetails {
	caseDetails := o.Cases
	return caseDetails
}

func makeCaseList(cd []types.CaseDetails) []*Case {
	var caseList []*Case
	for _, c := range cd {
		eachCase := Case{
			Subject:     *c.Subject,
			Status:      *c.Status,
			SubmittedBy: *c.SubmittedBy,
			TimeCreated: *c.TimeCreated,
			Url:         "https://console.aws.amazon.com/support/home#/case/?displayId=" + *c.DisplayId + "%26language=" + *c.Language,
		}
		caseList = append(caseList, &eachCase)
	}
	return caseList
}
