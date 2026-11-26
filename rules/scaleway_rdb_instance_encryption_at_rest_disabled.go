package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayRdbInstanceEncryptionAtRestDisabledRule checks whether "scaleway_rdb_instance" has encryption at rest disabled.
type ScalewayRdbInstanceEncryptionAtRestDisabledRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayRdbInstanceEncryptionAtRestDisabledRule returns a new rule.
func NewScalewayRdbInstanceEncryptionAtRestDisabledRule() *ScalewayRdbInstanceEncryptionAtRestDisabledRule {
	return &ScalewayRdbInstanceEncryptionAtRestDisabledRule{
		resourceType:  "scaleway_rdb_instance",
		attributeName: "encryption_at_rest",
	}
}

// Name returns the rule name.
func (r *ScalewayRdbInstanceEncryptionAtRestDisabledRule) Name() string {
	return "scaleway_rdb_instance_encryption_at_rest_disabled"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayRdbInstanceEncryptionAtRestDisabledRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayRdbInstanceEncryptionAtRestDisabledRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayRdbInstanceEncryptionAtRestDisabledRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_rdb_instance" has encryption at rest disabled.
func (r *ScalewayRdbInstanceEncryptionAtRestDisabledRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			if err := runner.EmitIssue(r, "Encryption at rest is not enabled for this Database Instance. Set encryption_at_rest = true.", resource.DefRange); err != nil {
				return err
			}
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(enabled bool) error {
			if enabled {
				return nil
			}
			return runner.EmitIssue(r, "Encryption at rest is disabled for this Database Instance. Set encryption_at_rest = true.", attribute.Expr.Range())
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
