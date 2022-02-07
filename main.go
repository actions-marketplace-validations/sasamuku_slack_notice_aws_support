package main

import (
	"github.com/sasamuku/slack_notice_aws_support/aws"
	"github.com/sasamuku/slack_notice_aws_support/slack"
)

func main() {
	aftertime := "2021-11-06T01:27:05.739Z"
	beforetime := "2022-01-06T01:27:05.739Z"
	language := "ja"
	webhookUrl := ""

	input := aws.NewDescribeCasesInput(aftertime, beforetime, language)
	cases := aws.GetCases(input)

	slack.Notify(cases, webhookUrl)
}
