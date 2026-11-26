package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayK8sPoolInvalidNodeType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid t2.micro",
			Content: `
resource "scaleway_k8s_pool" "main" {
  node_type = "t2.micro"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayK8sPoolInvalidNodeTypeRule(),
					Message: "\"t2.micro\" is an invalid Kubernetes pool node type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 15},
						End:      hcl.Pos{Line: 3, Column: 25},
					},
				},
			},
		},
		{
			Name: "invalid PRO3-S",
			Content: `
resource "scaleway_k8s_pool" "main" {
  node_type = "PRO3-S"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayK8sPoolInvalidNodeTypeRule(),
					Message: "\"PRO3-S\" is an invalid Kubernetes pool node type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 15},
						End:      hcl.Pos{Line: 3, Column: 23},
					},
				},
			},
		},
		{
			Name: "valid PRO2-S",
			Content: `
resource "scaleway_k8s_pool" "main" {
  node_type = "PRO2-S"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid pro2_xxs",
			Content: `
resource "scaleway_k8s_pool" "main" {
  node_type = "pro2_xxs"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid external",
			Content: `
resource "scaleway_k8s_pool" "main" {
  node_type = "external"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "value" {}

resource "scaleway_k8s_pool" "main" {
  node_type = var.value
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayK8sPoolInvalidNodeTypeRule()

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
