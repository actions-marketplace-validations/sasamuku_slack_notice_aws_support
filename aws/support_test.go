package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/support/types"
)

func Test_Support(t *testing.T) {
	caseId := "case-12345678910-2013-c4c1d2bf33c5cf47"
	ccEmailAddresses := []string{"cc@example.com"}
	displayId := "1234567890"
	language := "ja"
	status := "opened"
	subject := "Test Subject"
	submittedBy := "test@example.com"
	timeCreated := "2021-12-01T12:00:00.000Z"
	url := "https://console.aws.amazon.com/support/home#/case/?displayId=1234567890%26language=ja"

	caseDetails := []types.CaseDetails{
		{
			CaseId:           &caseId,
			CcEmailAddresses: ccEmailAddresses,
			DisplayId:        &displayId,
			Language:         &language,
			Status:           &status,
			Subject:          &subject,
			SubmittedBy:      &submittedBy,
			TimeCreated:      &timeCreated,
		},
	}
	caseList := makeCaseList(caseDetails)
	for _, c := range caseList {
		if c.Subject != subject {
			t.Fatalf("Fail: got = %v, want = %v", c.Subject, subject)
		}
		if c.Status != status {
			t.Fatalf("Fail: got = %v, want = %v", c.Status, status)
		}
		if c.SubmittedBy != submittedBy {
			t.Fatalf("Fail: got = %v, want = %v", c.SubmittedBy, submittedBy)
		}
		if c.TimeCreated != timeCreated {
			t.Fatalf("Fail: got = %v, want = %v", c.TimeCreated, timeCreated)
		}
		if c.Url != url {
			t.Fatalf("Fail: got = %v, want = %v", c.Url, url)
		}
	}
}
