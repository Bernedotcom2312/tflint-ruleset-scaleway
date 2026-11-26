package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayRedisClusterInvalidNodeType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid RED1-XXL",
			Content: `
resource "scaleway_redis_cluster" "main" {
  node_type = "RED1-XXL"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRedisClusterInvalidNodeTypeRule(),
					Message: "\"RED1-XXL\" is an invalid Managed Redis node type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 15},
						End:      hcl.Pos{Line: 3, Column: 25},
					},
				},
			},
		},
		{
			Name: "invalid db-dev-s",
			Content: `
resource "scaleway_redis_cluster" "main" {
  node_type = "db-dev-s"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRedisClusterInvalidNodeTypeRule(),
					Message: "\"db-dev-s\" is an invalid Managed Redis node type.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 15},
						End:      hcl.Pos{Line: 3, Column: 25},
					},
				},
			},
		},
		{
			Name: "valid RED1-micro",
			Content: `
resource "scaleway_redis_cluster" "main" {
  node_type = "RED1-micro"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid red1-xs",
			Content: `
resource "scaleway_redis_cluster" "main" {
  node_type = "red1-xs"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayRedisClusterInvalidNodeTypeRule()

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
