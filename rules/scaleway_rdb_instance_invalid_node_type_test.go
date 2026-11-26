package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayRdbInstanceInvalidNodeType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid db-dev-xxs",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type = "db-dev-xxs"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceInvalidNodeTypeRule(),
					Message: "\"db-dev-xxs\" is an invalid Managed Database node type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 15},
						End:      hcl.Pos{Line: 3, Column: 27},
					},
				},
			},
		},
		{
			Name: "invalid DB-GP-XXL",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type = "DB-GP-XXL"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstanceInvalidNodeTypeRule(),
					Message: "\"DB-GP-XXL\" is an invalid Managed Database node type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 15},
						End:      hcl.Pos{Line: 3, Column: 26},
					},
				},
			},
		},
		{
			Name: "valid db-dev-s",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type = "db-dev-s"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid DB-PRO2-XXS",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type = "DB-PRO2-XXS"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid db-play2-pico",
			Content: `
resource "scaleway_rdb_instance" "main" {
  node_type = "db-play2-pico"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "value" {}

resource "scaleway_rdb_instance" "main" {
  node_type = var.value
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayRdbInstanceInvalidNodeTypeRule()

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
