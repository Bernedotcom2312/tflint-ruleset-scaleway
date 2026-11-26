package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayInvalidRegion(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid region on a resource",
			Content: `
resource "scaleway_object_bucket" "main" {
  name   = "my-bucket"
  region = "eu-west-1"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInvalidRegionRule(),
					Message: "\"eu-west-1\" is an invalid Scaleway region.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 12},
						End:      hcl.Pos{Line: 4, Column: 23},
					},
				},
			},
		},
		{
			Name: "zone used as a region",
			Content: `
resource "scaleway_rdb_instance" "main" {
  region = "fr-par-1"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInvalidRegionRule(),
					Message: "\"fr-par-1\" is an invalid Scaleway region.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 22},
					},
				},
			},
		},
		{
			Name: "invalid region on the provider",
			Content: `
provider "scaleway" {
  region = "de-fra"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInvalidRegionRule(),
					Message: "\"de-fra\" is an invalid Scaleway region.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 12},
						End:      hcl.Pos{Line: 3, Column: 20},
					},
				},
			},
		},
		{
			Name: "valid regions",
			Content: `
provider "scaleway" {
  region = "fr-par"
}

data "scaleway_vpc" "default" {
  region = "pl-waw"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "other providers are ignored",
			Content: `
provider "aws" {
  region = "eu-west-3"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "region" {}

provider "scaleway" {
  region = var.region
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayInvalidRegionRule()

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
