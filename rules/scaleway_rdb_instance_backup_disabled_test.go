package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayRdbInstanceBackupDisabled(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "backups disabled",
			Content: `
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"

  disable_backup = true
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceBackupDisabledRule(),
					Message: "Automated backups are disabled for this Database Instance.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 7, Column: 20},
						End:      hcl.Pos{Line: 7, Column: 24},
					},
				},
			},
		},
		{
			Name: "backups enabled",
			Content: `
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"

  disable_backup = false
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "default",
			Content: `
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayRdbInstanceBackupDisabledRule()

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
