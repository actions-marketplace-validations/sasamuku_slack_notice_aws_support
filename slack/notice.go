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

type SlackNotice struct {
	WebhookUrl string
	Payload    *Payload
}

type Payload struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func ConvertToNoticeFormat(cases []*aws.Case) string {
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

	var formattedCases string
	for _, c := range cases {
		var buf bytes.Buffer
		if err := tpl.Execute(&buf, &c); err != nil {
			log.Fatal(err)
		}
		formattedCases = formattedCases + buf.String() + "---\n"
	}
	return formattedCases
}

func NewPayload(username, text string) *Payload {
	return &Payload{
		Username: username,
		Text:     text,
	}
}

func (p *Payload) toUnescapedJson() string {
	jsonPayload, err := json.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}
	unescapePayload, err := url.QueryUnescape(string(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}
	return unescapePayload
}

func NewSlackNotice(webhookUrl string, payload *Payload) *SlackNotice {
	return &SlackNotice{
		WebhookUrl: webhookUrl,
		Payload:    payload,
	}
}

func (s *SlackNotice) Run() (statusCode int) {
	resp, err := http.PostForm(s.WebhookUrl, url.Values{"payload": {s.Payload.toUnescapedJson()}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode
}
