package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayIAMPolicyNotAttachedToGroup(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "attached to a user",
			Content: `
resource "scaleway_iam_policy" "main" {
  user_id = scaleway_iam_user.alice.id
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayIAMPolicyNotAttachedToGroupRule(),
					Message: "The policy is attached directly to a user. Attach it to an IAM group (group_id) and add the user to the group.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 39},
					},
				},
			},
		},
		{
			Name: "attached to an application",
			Content: `
resource "scaleway_iam_policy" "main" {
  application_id = scaleway_iam_application.ci.id
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayIAMPolicyNotAttachedToGroupRule(),
					Message: "The policy is attached directly to an application. Attach it to an IAM group (group_id) and add the application to the group.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 3},
						End:      hcl.Pos{Line: 3, Column: 50},
					},
				},
			},
		},
		{
			Name: "attached to a group",
			Content: `
resource "scaleway_iam_policy" "main" {
  group_id = scaleway_iam_group.devs.id
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayIAMPolicyNotAttachedToGroupRule()

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
