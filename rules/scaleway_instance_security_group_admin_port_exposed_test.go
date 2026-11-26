package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_ScalewayInstanceSecurityGroupAdminPortExposed(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Config   string
		Expected helper.Issues
	}{
		{
			Name: "ssh open to the internet",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action   = "accept"
    port     = 22
    ip_range = "0.0.0.0/0"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupAdminPortExposedRule(),
					Message: "Inbound rule accepts administration port 22 from 0.0.0.0/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 3},
						End:      hcl.Pos{Line: 5, Column: 15},
					},
				},
			},
		},
		{
			Name: "no ip_range defaults to the internet",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action = "accept"
    port   = 3389
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupAdminPortExposedRule(),
					Message: "Inbound rule accepts administration port 3389 from 0.0.0.0/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 3},
						End:      hcl.Pos{Line: 5, Column: 15},
					},
				},
			},
		},
		{
			Name: "ipv6 range covering ssh",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action     = "accept"
    port_range = "20-25"
    ip_range   = "::/0"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupAdminPortExposedRule(),
					Message: "Inbound rule accepts administration port 22 from ::/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 3},
						End:      hcl.Pos{Line: 5, Column: 15},
					},
				},
			},
		},
		{
			Name: "no port means all ports",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action   = "accept"
    protocol = "ANY"
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupAdminPortExposedRule(),
					Message: "Inbound rule accepts administration port 22 from 0.0.0.0/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 3},
						End:      hcl.Pos{Line: 5, Column: 15},
					},
				},
			},
		},
		{
			Name: "security group rules resource",
			Content: `
resource "scaleway_instance_security_group_rules" "main" {
  security_group_id = "fr-par-1/11111111-1111-1111-1111-111111111111"

  inbound_rule {
    action = "accept"
    port   = 22
  }
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupAdminPortExposedRule(),
					Message: "Inbound rule accepts administration port 22 from 0.0.0.0/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 3},
						End:      hcl.Pos{Line: 5, Column: 15},
					},
				},
			},
		},
		{
			Name: "https is fine",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action = "accept"
    port   = 443
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "restricted source",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action   = "accept"
    port     = 22
    ip_range = "10.0.0.0/8"
  }

  inbound_rule {
    action = "accept"
    port   = 22
    ip     = "203.0.113.10"
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "drop rule",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action = "drop"
    port   = 22
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "udp",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action   = "accept"
    protocol = "UDP"
    port     = 22
  }
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "custom ports",
			Content: `
resource "scaleway_instance_security_group" "main" {
  inbound_default_policy = "drop"

  inbound_rule {
    action = "accept"
    port   = 22
  }

  inbound_rule {
    action = "accept"
    port   = 5432
  }
}`,
			Config: `
rule "scaleway_instance_security_group_admin_port_exposed" {
  enabled = true
  ports   = [5432]
}`,
			Expected: helper.Issues{
				{
					Rule:    NewScalewayInstanceSecurityGroupAdminPortExposedRule(),
					Message: "Inbound rule accepts administration port 5432 from 0.0.0.0/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 10, Column: 3},
						End:      hcl.Pos{Line: 10, Column: 15},
					},
				},
			},
		},
	}

	rule := NewScalewayInstanceSecurityGroupAdminPortExposedRule()

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
