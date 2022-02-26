package slack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sasamuku/slack_notice_aws_support/aws"
)

func Test_Notice(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		fmt.Fprintln(w, body)
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	type wants struct {
		statusCode int
	}

	tests := map[string]struct {
		cases      []*aws.Case
		webhookUrl string
		wants      wants
	}{
		"ok": {
			cases: []*aws.Case{
				{
					Subject:     "Test",
					Status:      "opened",
					SubmitteBy:  "user@example.com",
					TimeCreated: "2021-03-01T01:43:57.974Z",
					Url:         "https://console.aws.amazon.com/support/home#/case/?displayId=12345%26language=ja",
				},
			},
			webhookUrl: ts.URL + "/",
			wants:      wants{statusCode: 200},
		},
		"ok_empty_cases": {
			cases:      []*aws.Case{},
			webhookUrl: ts.URL + "/",
			wants:      wants{statusCode: 200},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			payload := NewPayload("AWS Support Case Notice", ConvertToNoticeFormat(tt.cases))
			notice := NewSlackNotice(tt.webhookUrl, payload)
			if statusCode := notice.Run(); statusCode != tt.wants.statusCode {
				t.Fatalf("run() status: got = %v, want = %v", statusCode, tt.wants.statusCode)
			}
		})
	}
}
