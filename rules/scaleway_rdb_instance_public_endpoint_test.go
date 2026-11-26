package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayRdbInstancePublicEndpoint(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "no endpoint block",
			Content: `
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstancePublicEndpointRule(),
					Message: "The Database Instance has no private_network block, so it gets a public endpoint by default. Attach it to a Private Network.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 40},
					},
				},
			},
		},
		{
			Name: "explicit public endpoint",
			Content: `
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"

  private_network {
    pn_id = scaleway_vpc_private_network.main.id
  }

  load_balancer {}
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayRdbInstancePublicEndpointRule(),
					Message: "The load_balancer block exposes the Database Instance on a public endpoint. Use a private_network endpoint only.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 11, Column: 3},
						End:      hcl.Pos{Line: 11, Column: 16},
					},
				},
			},
		},
		{
			Name: "private endpoint only",
			Content: `
resource "scaleway_rdb_instance" "main" {
  name      = "main"
  node_type = "DB-DEV-S"
  engine    = "PostgreSQL-16"

  private_network {
    pn_id       = scaleway_vpc_private_network.main.id
    enable_ipam = true
  }
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayRdbInstancePublicEndpointRule()

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
