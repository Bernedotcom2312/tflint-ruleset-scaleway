package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayObjectBucketVersioningDisabled(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no versioning block",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "backups"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayObjectBucketVersioningDisabledRule(),
					Message: "Versioning is not enabled on this bucket. Add a versioning block with enabled = true.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 41},
					},
				},
			},
		},
		{
			Name: "versioning disabled",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "backups"

  versioning {
    enabled = false
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayObjectBucketVersioningDisabledRule(),
					Message: "Versioning is disabled on this bucket. Set enabled = true in the versioning block.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 15},
						End:      hcl.Pos{Line: 6, Column: 20},
					},
				},
			},
		},
		{
			Name: "versioning enabled",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "backups"

  versioning {
    enabled = true
  }
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayObjectBucketVersioningDisabledRule()

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
