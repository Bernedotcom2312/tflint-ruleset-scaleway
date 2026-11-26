package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayRdbInstanceBackupDisabledRule checks whether "scaleway_rdb_instance"
// disables automated backups.
type ScalewayRdbInstanceBackupDisabledRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayRdbInstanceBackupDisabledRule returns a new rule.
func NewScalewayRdbInstanceBackupDisabledRule() *ScalewayRdbInstanceBackupDisabledRule {
	return &ScalewayRdbInstanceBackupDisabledRule{
		resourceType:  "scaleway_rdb_instance",
		attributeName: "disable_backup",
	}
}

// Name returns the rule name.
func (r *ScalewayRdbInstanceBackupDisabledRule) Name() string {
	return "scaleway_rdb_instance_backup_disabled"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayRdbInstanceBackupDisabledRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayRdbInstanceBackupDisabledRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayRdbInstanceBackupDisabledRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_rdb_instance" disables automated backups.
func (r *ScalewayRdbInstanceBackupDisabledRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(disabled bool) error {
			if !disabled {
				return nil
			}
			return runner.EmitIssue(
				r,
				"Automated backups are disabled for this Database Instance.",
				attribute.Expr.Range(),
			)
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
