package main

import (
	"fmt"

	"github.com/sasamuku/slack_notice_aws_support/aws"
)

func main() {
	aftertime := "2021-11-06T01:27:05.739Z"
	beforetime := "2022-01-06T01:27:05.739Z"
	language := "ja"

	input := aws.NewDescribeCasesInput(aftertime, beforetime, language)
	cases := aws.GetCases(input)

	for _, c := range *cases {
		fmt.Println(c.Subject)
	}
}
