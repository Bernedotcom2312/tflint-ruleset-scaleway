package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayInstanceSecurityGroupInboundDefaultPolicyRule checks whether
// "scaleway_instance_security_group" accepts inbound traffic by default.
type ScalewayInstanceSecurityGroupInboundDefaultPolicyRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayInstanceSecurityGroupInboundDefaultPolicyRule returns a new rule.
func NewScalewayInstanceSecurityGroupInboundDefaultPolicyRule() *ScalewayInstanceSecurityGroupInboundDefaultPolicyRule {
	return &ScalewayInstanceSecurityGroupInboundDefaultPolicyRule{
		resourceType:  "scaleway_instance_security_group",
		attributeName: "inbound_default_policy",
	}
}

// Name returns the rule name.
func (r *ScalewayInstanceSecurityGroupInboundDefaultPolicyRule) Name() string {
	return "scaleway_instance_security_group_inbound_default_policy"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayInstanceSecurityGroupInboundDefaultPolicyRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayInstanceSecurityGroupInboundDefaultPolicyRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayInstanceSecurityGroupInboundDefaultPolicyRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_instance_security_group" accepts inbound traffic by default.
func (r *ScalewayInstanceSecurityGroupInboundDefaultPolicyRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			if err := runner.EmitIssue(
				r,
				`inbound_default_policy is not set and defaults to "accept". Set it to "drop" and allow traffic with explicit inbound rules.`,
				resource.DefRange,
			); err != nil {
				return err
			}
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(policy string) error {
			if policy != "accept" {
				return nil
			}
			return runner.EmitIssue(
				r,
				`inbound_default_policy is "accept". Set it to "drop" and allow traffic with explicit inbound rules.`,
				attribute.Expr.Range(),
			)
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
