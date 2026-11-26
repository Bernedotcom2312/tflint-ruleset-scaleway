package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayK8sClusterInvalidCNI(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid flannel",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  cni = "flannel"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayK8sClusterInvalidCNIRule(),
					Message: "\"flannel\" is an invalid Kubernetes CNI plugin.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 9},
						End:      hcl.Pos{Line: 3, Column: 18},
					},
				},
			},
		},
		{
			Name: "invalid Cilium",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  cni = "Cilium"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayK8sClusterInvalidCNIRule(),
					Message: "\"Cilium\" is an invalid Kubernetes CNI plugin.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 9},
						End:      hcl.Pos{Line: 3, Column: 17},
					},
				},
			},
		},
		{
			Name: "valid cilium",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  cni = "cilium"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid cilium_native",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  cni = "cilium_native"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid calico",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  cni = "calico"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid kilo",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  cni = "kilo"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid none",
			Content: `
resource "scaleway_k8s_cluster" "main" {
  cni = "none"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "unknown value is ignored",
			Content: `
variable "value" {}

resource "scaleway_k8s_cluster" "main" {
  cni = var.value
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewScalewayK8sClusterInvalidCNIRule()

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
