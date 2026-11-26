package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayInvalidZone(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid zone on a resource",
			Content: `
resource "scaleway_instance_server" "main" {
  type = "DEV1-S"
  zone = "fr-par-4"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInvalidZoneRule(),
					Message: "\"fr-par-4\" is an invalid Scaleway zone.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 10},
						End:      hcl.Pos{Line: 4, Column: 20},
					},
				},
			},
		},
		{
			Name: "region used as a zone",
			Content: `
resource "scaleway_k8s_pool" "main" {
  zone = "fr-par"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInvalidZoneRule(),
					Message: "\"fr-par\" is an invalid Scaleway zone.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 18},
					},
				},
			},
		},
		{
			Name: "invalid zone on a data source",
			Content: `
data "scaleway_instance_image" "ubuntu" {
  zone = "eu-west-1a"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInvalidZoneRule(),
					Message: "\"eu-west-1a\" is an invalid Scaleway zone.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 22},
					},
				},
			},
		},
		{
			Name: "invalid zone on the provider",
			Content: `
provider "scaleway" {
  zone = "nl-ams-4"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInvalidZoneRule(),
					Message: "\"nl-ams-4\" is an invalid Scaleway zone.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 20},
					},
				},
			},
		},
		{
			Name: "valid zones",
			Content: `
provider "scaleway" {
  zone = "fr-par-1"
}

resource "scaleway_instance_server" "main" {
  zone = "it-mil-1"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "other providers are ignored",
			Content: `
provider "google" {
  zone = "europe-west1-b"
}

resource "google_compute_instance" "main" {
  zone = "europe-west1-b"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayInvalidZoneRule()

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
