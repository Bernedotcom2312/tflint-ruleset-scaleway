package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayLBInvalidType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid LB-XXL",
			Content: `
resource "scaleway_lb" "main" {
  type = "LB-XXL"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayLBInvalidTypeRule(),
					Message: "\"LB-XXL\" is an invalid Load Balancer commercial type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 18},
					},
				},
			},
		},
		{
			Name: "invalid LB-GP-S-MULTICLOUD",
			Content: `
resource "scaleway_lb" "main" {
  type = "LB-GP-S-MULTICLOUD"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayLBInvalidTypeRule(),
					Message: "\"LB-GP-S-MULTICLOUD\" is an invalid Load Balancer commercial type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 30},
					},
				},
			},
		},
		{
			Name: "valid LB-S",
			Content: `
resource "scaleway_lb" "main" {
  type = "LB-S"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid lb-gp-xl",
			Content: `
resource "scaleway_lb" "main" {
  type = "lb-gp-xl"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "value" {}

resource "scaleway_lb" "main" {
  type = var.value
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayLBInvalidTypeRule()

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
