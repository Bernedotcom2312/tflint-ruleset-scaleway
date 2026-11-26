package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayInstanceServerInvalidType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid DEV1-XXL",
			Content: `
resource "scaleway_instance_server" "main" {
  type = "DEV1-XXL"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceServerInvalidTypeRule(),
					Message: "\"DEV1-XXL\" is an invalid Instance commercial type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 20},
					},
				},
			},
		},
		{
			Name: "invalid t2.micro",
			Content: `
resource "scaleway_instance_server" "main" {
  type = "t2.micro"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceServerInvalidTypeRule(),
					Message: "\"t2.micro\" is an invalid Instance commercial type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 20},
					},
				},
			},
		},
		{
			Name: "valid DEV1-S",
			Content: `
resource "scaleway_instance_server" "main" {
  type = "DEV1-S"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid pro2-xxs",
			Content: `
resource "scaleway_instance_server" "main" {
  type = "pro2-xxs"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid POP2-HC-2C-4G",
			Content: `
resource "scaleway_instance_server" "main" {
  type = "POP2-HC-2C-4G"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "value" {}

resource "scaleway_instance_server" "main" {
  type = var.value
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayInstanceServerInvalidTypeRule()

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
