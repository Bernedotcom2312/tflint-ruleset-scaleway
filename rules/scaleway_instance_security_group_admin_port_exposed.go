package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// defaultAdminPorts are the administration ports checked when the rule has
// no "ports" configuration: SSH and RDP.
var defaultAdminPorts = []int{22, 3389}

// ScalewayInstanceSecurityGroupAdminPortExposedRule checks whether a
// security group accepts administration ports from the whole internet.
type ScalewayInstanceSecurityGroupAdminPortExposedRule struct {
	tflint.DefaultRule

	resourceTypes []string
}

type adminPortExposedRuleConfig struct {
	Ports []int `hclext:"ports,optional"`
}

// NewScalewayInstanceSecurityGroupAdminPortExposedRule returns a new rule.
func NewScalewayInstanceSecurityGroupAdminPortExposedRule() *ScalewayInstanceSecurityGroupAdminPortExposedRule {
	return &ScalewayInstanceSecurityGroupAdminPortExposedRule{
		resourceTypes: []string{"scaleway_instance_security_group", "scaleway_instance_security_group_rules"},
	}
}

// Name returns the rule name.
func (r *ScalewayInstanceSecurityGroupAdminPortExposedRule) Name() string {
	return "scaleway_instance_security_group_admin_port_exposed"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayInstanceSecurityGroupAdminPortExposedRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayInstanceSecurityGroupAdminPortExposedRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayInstanceSecurityGroupAdminPortExposedRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether inbound rules accept administration ports from anywhere.
func (r *ScalewayInstanceSecurityGroupAdminPortExposedRule) Check(runner tflint.Runner) error {
	config := adminPortExposedRuleConfig{}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}
	ports := config.Ports
	if len(ports) == 0 {
		ports = defaultAdminPorts
	}

	schema := &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "inbound_rule",
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{
						{Name: "action"},
						{Name: "protocol"},
						{Name: "port"},
						{Name: "port_range"},
						{Name: "ip"},
						{Name: "ip_range"},
					},
				},
			},
		},
	}

	for _, resourceType := range r.resourceTypes {
		resources, err := runner.GetResourceContent(resourceType, schema, nil)
		if err != nil {
			return err
		}

		for _, resource := range resources.Blocks {
			for _, rule := range resource.Body.Blocks {
				port, source, exposed, err := r.exposedPort(runner, rule.Body.Attributes, ports)
				if err != nil {
					return err
				}
				if !exposed {
					continue
				}

				if err := runner.EmitIssue(
					r,
					fmt.Sprintf("Inbound rule accepts administration port %d from %s. Reach Instances through the Public Gateway SSH bastion or a VPN instead.", port, source),
					rule.DefRange,
				); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// exposedPort reports whether an inbound rule accepts one of ports from the
// whole internet, mirroring the provider defaults: protocol defaults to TCP,
// a missing ip_range means 0.0.0.0/0 and a missing port means all ports.
// Rules whose values cannot be determined statically are ignored.
func (r *ScalewayInstanceSecurityGroupAdminPortExposedRule) exposedPort(runner tflint.Runner, attrs hclext.Attributes, ports []int) (int, string, bool, error) {
	action, known, err := r.stringAttribute(runner, attrs, "action", "")
	if err != nil || !known || action != "accept" {
		return 0, "", false, err
	}

	protocol, known, err := r.stringAttribute(runner, attrs, "protocol", "TCP")
	if err != nil || !known {
		return 0, "", false, err
	}
	if protocol = strings.ToUpper(protocol); protocol != "TCP" && protocol != "ANY" {
		return 0, "", false, nil
	}

	if _, exists := attrs["ip"]; exists {
		// A single address (/32) is never the whole internet.
		return 0, "", false, nil
	}
	source, known, err := r.stringAttribute(runner, attrs, "ip_range", "0.0.0.0/0")
	if err != nil || !known || (source != "0.0.0.0/0" && source != "::/0") {
		return 0, "", false, err
	}

	from, to := 0, 0
	if attr, exists := attrs["port_range"]; exists {
		portRange, known, err := evaluateString(runner, attr)
		if err != nil || !known {
			return 0, "", false, err
		}
		if _, err := fmt.Sscanf(portRange, "%d-%d", &from, &to); err != nil {
			return 0, "", false, nil
		}
	} else if attr, exists := attrs["port"]; exists {
		port, known, err := evaluateInt(runner, attr)
		if err != nil || !known {
			return 0, "", false, err
		}
		from, to = port, port
	}

	for _, port := range ports {
		// Port 0 (or no port at all) means every port.
		if (from == 0 && to == 0) || (port >= from && port <= to) {
			return port, source, true, nil
		}
	}
	return 0, "", false, nil
}

func (r *ScalewayInstanceSecurityGroupAdminPortExposedRule) stringAttribute(runner tflint.Runner, attrs hclext.Attributes, name, defaultValue string) (string, bool, error) {
	attr, exists := attrs[name]
	if !exists {
		return defaultValue, defaultValue != "", nil
	}
	return evaluateString(runner, attr)
}
