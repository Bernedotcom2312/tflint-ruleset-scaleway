package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayRedisClusterPublicEndpoint(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no private network",
			Content: `
resource "scaleway_redis_cluster" "main" {
  name      = "main"
  version   = "7.2.11"
  node_type = "RED1-micro"

  acl {
    ip = "0.0.0.0/0"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRedisClusterPublicEndpointRule(),
					Message: "The Redis cluster has no private_network block, so it is exposed on a public endpoint. Attach it to a Private Network.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 41},
					},
				},
			},
		},
		{
			Name: "private network",
			Content: `
resource "scaleway_redis_cluster" "main" {
  name      = "main"
  version   = "7.2.11"
  node_type = "RED1-micro"

  private_network {
    id = scaleway_vpc_private_network.main.id
  }
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayRedisClusterPublicEndpointRule()

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
