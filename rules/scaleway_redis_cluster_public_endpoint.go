package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayRedisClusterPublicEndpointRule checks whether "scaleway_redis_cluster"
// is exposed on a public endpoint.
type ScalewayRedisClusterPublicEndpointRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewScalewayRedisClusterPublicEndpointRule returns a new rule.
func NewScalewayRedisClusterPublicEndpointRule() *ScalewayRedisClusterPublicEndpointRule {
	return &ScalewayRedisClusterPublicEndpointRule{resourceType: "scaleway_redis_cluster"}
}

// Name returns the rule name.
func (r *ScalewayRedisClusterPublicEndpointRule) Name() string {
	return "scaleway_redis_cluster_public_endpoint"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayRedisClusterPublicEndpointRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayRedisClusterPublicEndpointRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayRedisClusterPublicEndpointRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_redis_cluster" is exposed on a public endpoint.
//
// The provider exposes the cluster on a public endpoint when no
// private_network block is set.
func (r *ScalewayRedisClusterPublicEndpointRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "private_network", Body: &hclext.BodySchema{}},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if len(resource.Body.Blocks.OfType("private_network")) > 0 {
			continue
		}
		if err := runner.EmitIssue(
			r,
			"The Redis cluster has no private_network block, so it is exposed on a public endpoint. Attach it to a Private Network.",
			resource.DefRange,
		); err != nil {
			return err
		}
	}

	return nil
}
