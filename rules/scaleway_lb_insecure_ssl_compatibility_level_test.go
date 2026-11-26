package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayLBInsecureSSLCompatibilityLevel(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "old",
			Content: `
resource "scaleway_lb" "main" {
  type                    = "LB-S"
  ssl_compatibility_level = "ssl_compatibility_level_old"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayLBInsecureSSLCompatibilityLevelRule(),
					Message: "ssl_compatibility_level_old accepts TLS versions older than 1.2. Use ssl_compatibility_level_intermediate (the default) or ssl_compatibility_level_modern.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 29},
						End:      hcl.Pos{Line: 4, Column: 58},
					},
				},
			},
		},
		{
			Name: "intermediate",
			Content: `
resource "scaleway_lb" "main" {
  type                    = "LB-S"
  ssl_compatibility_level = "ssl_compatibility_level_intermediate"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "modern",
			Content: `
resource "scaleway_lb" "main" {
  ssl_compatibility_level = "ssl_compatibility_level_modern"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default",
			Content: `
resource "scaleway_lb" "main" {
  type = "LB-S"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayLBInsecureSSLCompatibilityLevelRule()

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
