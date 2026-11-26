package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayIAMAPIKeyMissingExpiration(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no expiration",
			Content: `
resource "scaleway_iam_api_key" "main" {
  application_id = scaleway_iam_application.ci.id
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayIAMAPIKeyMissingExpirationRule(),
					Message: "The API key never expires. Set expires_at and rotate the key before that date.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 39},
					},
				},
			},
		},
		{
			Name: "expiration set",
			Content: `
resource "scaleway_iam_api_key" "main" {
  application_id = scaleway_iam_application.ci.id
  expires_at     = "2027-06-30T00:00:00Z"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "expiration from a variable",
			Content: `
variable "expires_at" {}

resource "scaleway_iam_api_key" "main" {
  expires_at = var.expires_at
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayIAMAPIKeyMissingExpirationRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
