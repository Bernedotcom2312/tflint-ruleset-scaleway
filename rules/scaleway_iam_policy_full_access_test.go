package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayIAMPolicyFullAccess(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "full access",
			Content: `
resource "scaleway_iam_policy" "main" {
  group_id = scaleway_iam_group.admins.id

  rule {
    project_ids          = [scaleway_account_project.main.id]
    permission_set_names = ["AllProductsFullAccess"]
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayIAMPolicyFullAccessRule(),
					Message: "AllProductsFullAccess grants every permission on every product. Use product-specific permission sets, such as InstancesFullAccess.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 7, Column: 28},
						End:      hcl.Pos{Line: 7, Column: 53},
					},
				},
			},
		},
		{
			Name: "full access in a second rule",
			Content: `
resource "scaleway_iam_policy" "main" {
  group_id = scaleway_iam_group.admins.id

  rule {
    project_ids          = [scaleway_account_project.main.id]
    permission_set_names = ["InstancesReadOnly"]
  }

  rule {
    organization_id      = "11111111-1111-1111-1111-111111111111"
    permission_set_names = ["IAMReadOnly", "AllProductsFullAccess"]
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayIAMPolicyFullAccessRule(),
					Message: "AllProductsFullAccess grants every permission on every product. Use product-specific permission sets, such as InstancesFullAccess.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 12, Column: 28},
						End:      hcl.Pos{Line: 12, Column: 68},
					},
				},
			},
		},
		{
			Name: "specific permission sets",
			Content: `
resource "scaleway_iam_policy" "main" {
  group_id = scaleway_iam_group.devs.id

  rule {
    project_ids          = [scaleway_account_project.main.id]
    permission_set_names = ["InstancesFullAccess", "ObjectStorageReadOnly"]
  }
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayIAMPolicyFullAccessRule()

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
