package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayRdbInstanceEncryptionAtRestDisabled(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "not set",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type = "DB-DEV-S"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceEncryptionAtRestDisabledRule(),
					Message: "Encryption at rest is not enabled for this Database Instance. Set encryption_at_rest = true.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 40},
					},
				},
			},
		},
		{
			Name: "disabled",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type          = "DB-DEV-S"
  encryption_at_rest = false
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceEncryptionAtRestDisabledRule(),
					Message: "Encryption at rest is disabled for this Database Instance. Set encryption_at_rest = true.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 24},
						End:      hcl.Pos{Line: 4, Column: 29},
					},
				},
			},
		},
		{
			Name: "enabled",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type          = "DB-DEV-S"
  encryption_at_rest = true
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "encryption" {}

resource "scaleway_rdb_instance" "main" {
  encryption_at_rest = var.encryption
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayRdbInstanceEncryptionAtRestDisabledRule()

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
