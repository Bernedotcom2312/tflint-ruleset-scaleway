package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayLBInvalidTypeRule checks whether "scaleway_lb" has an invalid commercial type.
type ScalewayLBInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayLBInvalidTypeRule returns a new rule.
func NewScalewayLBInvalidTypeRule() *ScalewayLBInvalidTypeRule {
	return &ScalewayLBInvalidTypeRule{
		resourceType:  "scaleway_lb",
		attributeName: "type",
	}
}

// Name returns the rule name.
func (r *ScalewayLBInvalidTypeRule) Name() string {
	return "scaleway_lb_invalid_type"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayLBInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayLBInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayLBInvalidTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_lb" has an invalid commercial type.
func (r *ScalewayLBInvalidTypeRule) Check(runner tflint.Runner) error {
	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, models.IsValidLBType, func(value string) string {
		return fmt.Sprintf(`"%s" is an invalid Load Balancer commercial type.`, value)
	})
}
