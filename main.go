package main

import (
	"fmt"

	"github.com/sasamuku/aws_support_to_github_issue/aws"
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
