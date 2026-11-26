package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayRedisClusterInvalidNodeTypeRule checks whether "scaleway_redis_cluster" has an invalid node type.
type ScalewayRedisClusterInvalidNodeTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayRedisClusterInvalidNodeTypeRule returns a new rule.
func NewScalewayRedisClusterInvalidNodeTypeRule() *ScalewayRedisClusterInvalidNodeTypeRule {
	return &ScalewayRedisClusterInvalidNodeTypeRule{
		resourceType:  "scaleway_redis_cluster",
		attributeName: "node_type",
	}
}

// Name returns the rule name.
func (r *ScalewayRedisClusterInvalidNodeTypeRule) Name() string {
	return "scaleway_redis_cluster_invalid_node_type"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayRedisClusterInvalidNodeTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayRedisClusterInvalidNodeTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayRedisClusterInvalidNodeTypeRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_redis_cluster" has an invalid node type.
func (r *ScalewayRedisClusterInvalidNodeTypeRule) Check(runner tflint.Runner) error {
	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, models.IsValidRedisNodeType, func(value string) string {
		return fmt.Sprintf(`"%s" is an invalid Managed Redis node type.`, value)
	})
}
