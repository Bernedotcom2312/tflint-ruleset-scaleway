package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayInstanceServerInvalidTypeRule checks whether "scaleway_instance_server" has an invalid commercial type.
type ScalewayInstanceServerInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayInstanceServerInvalidTypeRule returns a new rule.
func NewScalewayInstanceServerInvalidTypeRule() *ScalewayInstanceServerInvalidTypeRule {
	return &ScalewayInstanceServerInvalidTypeRule{
		resourceType:  "scaleway_instance_server",
		attributeName: "type",
	}
}

// Name returns the rule name.
func (r *ScalewayInstanceServerInvalidTypeRule) Name() string {
	return "scaleway_instance_server_invalid_type"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayInstanceServerInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayInstanceServerInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayInstanceServerInvalidTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_instance_server" has an invalid commercial type.
func (r *ScalewayInstanceServerInvalidTypeRule) Check(runner tflint.Runner) error {
	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, models.IsValidInstanceServerType, func(value string) string {
		return fmt.Sprintf(`"%s" is an invalid Instance commercial type.`, value)
	})
}
