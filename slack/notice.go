package slack

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/sasamuku/slack_notice_aws_support/aws"
)

type Payload struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func Notify(cases *[]aws.Case, webhookUrl string) {
	jsonCases, err := json.Marshal(cases)
	if err != nil {
		log.Fatal(err)
	}
	payload, err := json.Marshal(Payload{
		Username: "AWS Support Case Notice",
		Text:     string(jsonCases),
	})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.PostForm(webhookUrl, url.Values{"payload": {string(payload)}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
