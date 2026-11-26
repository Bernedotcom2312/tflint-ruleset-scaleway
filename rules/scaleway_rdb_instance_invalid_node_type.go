package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayRdbInstanceInvalidNodeTypeRule checks whether "scaleway_rdb_instance" has an invalid node type.
type ScalewayRdbInstanceInvalidNodeTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayRdbInstanceInvalidNodeTypeRule returns a new rule.
func NewScalewayRdbInstanceInvalidNodeTypeRule() *ScalewayRdbInstanceInvalidNodeTypeRule {
	return &ScalewayRdbInstanceInvalidNodeTypeRule{
		resourceType:  "scaleway_rdb_instance",
		attributeName: "node_type",
	}
}

// Name returns the rule name.
func (r *ScalewayRdbInstanceInvalidNodeTypeRule) Name() string {
	return "scaleway_rdb_instance_invalid_node_type"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayRdbInstanceInvalidNodeTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayRdbInstanceInvalidNodeTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayRdbInstanceInvalidNodeTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_rdb_instance" has an invalid node type.
func (r *ScalewayRdbInstanceInvalidNodeTypeRule) Check(runner tflint.Runner) error {
	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, models.IsValidRdbNodeType, func(value string) string {
		return fmt.Sprintf(`"%s" is an invalid Managed Database node type.`, value)
	})
}
