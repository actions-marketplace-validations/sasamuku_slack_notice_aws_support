package slack

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/sasamuku/slack_notice_aws_support/aws"
)

type Payload struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func Notify(cases []*aws.Case, webhookUrl string) {
	text := `Subject: {{ .Subject }}
Status: {{ .Status }}
SubmitteBy: {{ .SubmitteBy }}
TimeCreated: {{ .TimeCreated }}
Url: {{ .Url }}
`
	tpl, err := template.New("").Parse(text)
	if err != nil {
		log.Fatal(err)
	}

	var consolidatedCases string
	for _, c := range cases {
		var buf bytes.Buffer
		if err := tpl.Execute(&buf, &c); err != nil {
			log.Fatal(err)
		}
		consolidatedCases = consolidatedCases + buf.String() + "---\n"
	}

	payload, err := json.Marshal(Payload{
		Username: "AWS Support Case Notice",
		Text:     consolidatedCases,
	})
	if err != nil {
		log.Fatal(err)
	}

	unescapePayload, err := url.QueryUnescape(string(payload))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.PostForm(webhookUrl, url.Values{"payload": {unescapePayload}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
}
