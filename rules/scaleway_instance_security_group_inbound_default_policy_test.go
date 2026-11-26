package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayInstanceSecurityGroupInboundDefaultPolicy(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "policy not set",
			Content: `
resource "scaleway_instance_security_group" "main" {
  name = "web"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupInboundDefaultPolicyRule(),
					Message: "inbound_default_policy is not set and defaults to \"accept\". Set it to \"drop\" and allow traffic with explicit inbound rules.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 51},
					},
				},
			},
		},
		{
			Name: "policy is accept",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "accept"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupInboundDefaultPolicyRule(),
					Message: "inbound_default_policy is \"accept\". Set it to \"drop\" and allow traffic with explicit inbound rules.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 28},
						End:      hcl.Pos{Line: 3, Column: 36},
					},
				},
			},
		},
		{
			Name: "policy is drop",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy  = "drop"
  outbound_default_policy = "accept"

  inbound_rule {
    action = "accept"
    port   = 443
  }
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayInstanceSecurityGroupInboundDefaultPolicyRule()

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
