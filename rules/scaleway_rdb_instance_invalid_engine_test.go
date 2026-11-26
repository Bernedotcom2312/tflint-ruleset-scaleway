package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayRdbInstanceInvalidEngine(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid postgres-16",
			Content: `
resource "scaleway_rdb_instance" "main" {
  engine = "postgres-16"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceInvalidEngineRule(),
					Message: "\"postgres-16\" is an invalid Managed Database engine. It must look like \"PostgreSQL-16\" or \"MySQL-8\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 25},
					},
				},
			},
		},
		{
			Name: "invalid PostgreSQL16",
			Content: `
resource "scaleway_rdb_instance" "main" {
  engine = "PostgreSQL16"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceInvalidEngineRule(),
					Message: "\"PostgreSQL16\" is an invalid Managed Database engine. It must look like \"PostgreSQL-16\" or \"MySQL-8\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 26},
					},
				},
			},
		},
		{
			Name: "invalid MariaDB-11",
			Content: `
resource "scaleway_rdb_instance" "main" {
  engine = "MariaDB-11"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceInvalidEngineRule(),
					Message: "\"MariaDB-11\" is an invalid Managed Database engine. It must look like \"PostgreSQL-16\" or \"MySQL-8\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 24},
					},
				},
			},
		},
		{
			Name: "valid PostgreSQL-16",
			Content: `
resource "scaleway_rdb_instance" "main" {
  engine = "PostgreSQL-16"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid MySQL-8",
			Content: `
resource "scaleway_rdb_instance" "main" {
  engine = "MySQL-8"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid postgresql-9.6",
			Content: `
resource "scaleway_rdb_instance" "main" {
  engine = "postgresql-9.6"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "value" {}

resource "scaleway_rdb_instance" "main" {
  engine = var.value
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayRdbInstanceInvalidEngineRule()

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
