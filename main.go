package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sasamuku/slack_notice_aws_support/aws"
	"github.com/sasamuku/slack_notice_aws_support/slack"
)

func main() {
	now := time.Now().UTC()
	lastYear := now.AddDate(-1, 0, 0)
	beforetime := now.Format(time.RFC3339)
	aftertime := lastYear.Format(time.RFC3339)

	include_resolved_cases := os.Getenv("INPUT_INCLUDE_RESOLVED_CASES")
	language := os.Getenv("INPUT_LANGUAGE")
	webhookUrl := os.Getenv("INPUT_WEBHOOK_URL")
	include_resolved_cases_b, err := strconv.ParseBool(include_resolved_cases)
	if err != nil {
		log.Fatal(err)
	}

	input := aws.NewDescribeCasesInput(aftertime, beforetime, language, include_resolved_cases_b)
	caseList := aws.GetCaseList(input)

	payload := slack.NewPayload("AWS Support Case Notice", slack.ConvertToNoticeFormat(caseList))
	notice := slack.NewSlackNotice(webhookUrl, payload)
	notice.Run()
}
