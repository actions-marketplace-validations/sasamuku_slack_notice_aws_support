package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/sasamuku/slack_notice_aws_support/aws"
)

func Test_Notify(t *testing.T) {
	var message string

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		fmt.Fprintln(w, body)

		message, _ = url.QueryUnescape(string(body))
		message = strings.Replace(message, "payload=", "", 1)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	tests := map[string]struct {
		cases      *[]aws.Case
		webhookUrl string
	}{
		"existed": {
			cases: &[]aws.Case{
				{
					Subject:     "Test",
					Status:      "opened",
					SubmitteBy:  "user@example.com",
					TimeCreated: "2021-03-01T01:43:57.974Z",
					Url:         "https://console.aws.amazon.com/support/home#/case/?displayId=12345%26language=ja",
				},
			},
			webhookUrl: ts.URL + "/",
		},
		"empty": {
			cases:      &[]aws.Case{},
			webhookUrl: ts.URL + "/",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			Notify(tt.cases, tt.webhookUrl)
			payload := Payload{}
			json.Unmarshal([]byte(message), &payload)
			fmt.Println(payload.Text)
		})
	}
}
