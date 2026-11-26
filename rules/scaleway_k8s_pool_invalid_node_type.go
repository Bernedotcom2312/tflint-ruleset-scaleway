package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayK8sPoolInvalidNodeTypeRule checks whether "scaleway_k8s_pool" has an invalid node type.
type ScalewayK8sPoolInvalidNodeTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayK8sPoolInvalidNodeTypeRule returns a new rule.
func NewScalewayK8sPoolInvalidNodeTypeRule() *ScalewayK8sPoolInvalidNodeTypeRule {
	return &ScalewayK8sPoolInvalidNodeTypeRule{
		resourceType:  "scaleway_k8s_pool",
		attributeName: "node_type",
	}
}

// Name returns the rule name.
func (r *ScalewayK8sPoolInvalidNodeTypeRule) Name() string {
	return "scaleway_k8s_pool_invalid_node_type"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayK8sPoolInvalidNodeTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayK8sPoolInvalidNodeTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayK8sPoolInvalidNodeTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_k8s_pool" has an invalid node type.
func (r *ScalewayK8sPoolInvalidNodeTypeRule) Check(runner tflint.Runner) error {
	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, models.IsValidK8sNodeType, func(value string) string {
		return fmt.Sprintf(`"%s" is an invalid Kubernetes pool node type.`, value)
	})
}
