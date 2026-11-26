package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayRdbInstancePublicEndpointRule checks whether "scaleway_rdb_instance"
// is exposed on a public endpoint.
type ScalewayRdbInstancePublicEndpointRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewScalewayRdbInstancePublicEndpointRule returns a new rule.
func NewScalewayRdbInstancePublicEndpointRule() *ScalewayRdbInstancePublicEndpointRule {
	return &ScalewayRdbInstancePublicEndpointRule{resourceType: "scaleway_rdb_instance"}
}

// Name returns the rule name.
func (r *ScalewayRdbInstancePublicEndpointRule) Name() string {
	return "scaleway_rdb_instance_public_endpoint"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayRdbInstancePublicEndpointRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayRdbInstancePublicEndpointRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayRdbInstancePublicEndpointRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_rdb_instance" is exposed on a public endpoint.
//
// The provider creates a public endpoint when the load_balancer block is set,
// and also by default when neither private_network nor load_balancer is set.
func (r *ScalewayRdbInstancePublicEndpointRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "private_network", Body: &hclext.BodySchema{}},
			{Type: "load_balancer", Body: &hclext.BodySchema{}},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if lb := resource.Body.Blocks.OfType("load_balancer"); len(lb) > 0 {
			if err := runner.EmitIssue(
				r,
				"The load_balancer block exposes the Database Instance on a public endpoint. Use a private_network endpoint only.",
				lb[0].DefRange,
			); err != nil {
				return err
			}
			continue
		}

		if len(resource.Body.Blocks.OfType("private_network")) == 0 {
			if err := runner.EmitIssue(
				r,
				"The Database Instance has no private_network block, so it gets a public endpoint by default. Attach it to a Private Network.",
				resource.DefRange,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
