package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayObjectBucketInvalidName(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid ab",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "ab"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayObjectBucketInvalidNameRule(),
					Message: "\"ab\" is an invalid bucket name: it must be between 3 and 63 characters long.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 14},
					},
				},
			},
		},
		{
			Name: "invalid My_Bucket",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "My_Bucket"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayObjectBucketInvalidNameRule(),
					Message: "\"My_Bucket\" is an invalid bucket name: it can only contain lowercase letters, numbers, dots (.) and hyphens (-).",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 21},
					},
				},
			},
		},
		{
			Name: "invalid -my-bucket",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "-my-bucket"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayObjectBucketInvalidNameRule(),
					Message: "\"-my-bucket\" is an invalid bucket name: it must begin and end with a lowercase letter or a number.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 22},
					},
				},
			},
		},
		{
			Name: "invalid 192.168.10.1",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "192.168.10.1"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayObjectBucketInvalidNameRule(),
					Message: "\"192.168.10.1\" is an invalid bucket name: it must not be formatted as an IP address.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 24},
					},
				},
			},
		},
		{
			Name: "too long",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayObjectBucketInvalidNameRule(),
					Message: "\"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\" is an invalid bucket name: it must be between 3 and 63 characters long.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 10},
						End:      hcl.Pos{Line: 3, Column: 76},
					},
				},
			},
		},
		{
			Name: "valid my-bucket",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "my-bucket"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid logs.example.com",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "logs.example.com"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid abc",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "abc"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid aaaaaaaaaaaaaaaaaaaa",
			Content: `
resource "scaleway_object_bucket" "main" {
  name = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "suffix" {}

resource "scaleway_object_bucket" "main" {
  name = "logs-${var.suffix}"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayObjectBucketInvalidNameRule()

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
