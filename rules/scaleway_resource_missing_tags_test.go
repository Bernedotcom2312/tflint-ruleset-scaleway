package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayResourceMissingTags(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "no tags",
			Content: `
resource "scaleway_instance_server" "web" {
  type = "DEV1-S"
}`,
			Config: `
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayResourceMissingTagsRule(),
					Message: "The resource is missing the following tags: \"env\", \"owner\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 42},
					},
				},
			},
		},
		{
			Name: "missing one tag",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  tags = ["env=prod", "team=payments"]
}`,
			Config: `
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayResourceMissingTagsRule(),
					Message: "The resource is missing the following tags: \"owner\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 39},
					},
				},
			},
		},
		{
			Name: "key:value and bare tags",
			Content: `
resource "scaleway_vpc_private_network" "main" {
  tags = ["env:prod", "owner"]
}`,
			Config: `
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "map tags on buckets",
			Content: `
resource "scaleway_object_bucket" "logs" {
  name = "logs"
  tags = {
    env = "prod"
  }
}`,
			Config: `
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayResourceMissingTagsRule(),
					Message: "The resource is missing the following tags: \"owner\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 10},
						End:      hcl.Pos{Line: 6, Column: 4},
					},
				},
			},
		},
		{
			Name: "tag prefix is not a match",
			Content: `
resource "scaleway_lb" "main" {
  tags = ["environment=prod", "owner=alice"]
}`,
			Config: `
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayResourceMissingTagsRule(),
					Message: "The resource is missing the following tags: \"env\".",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 45},
					},
				},
			},
		},
		{
			Name: "excluded and non taggable resources",
			Content: `
resource "scaleway_k8s_pool" "default" {
  node_type = "PRO2-S"
}

resource "scaleway_rdb_acl" "main" {
  instance_id = "fr-par/11111111-1111-1111-1111-111111111111"
}

resource "aws_instance" "other" {
  instance_type = "t3.micro"
}`,
			Config: `
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
  exclude = ["scaleway_k8s_pool"]
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "no configuration",
			Content: `
resource "scaleway_instance_server" "web" {
  type = "DEV1-S"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown tags are ignored",
			Content: `
variable "tags" {}

resource "scaleway_instance_server" "web" {
  tags = var.tags
}`,
			Config: `
rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner"]
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayResourceMissingTagsRule()

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content, ".tflint.hcl": tc.Config})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("Unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, tc.Expected, runner.Issues)
		})
	}
}
