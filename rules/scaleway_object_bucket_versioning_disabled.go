package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayObjectBucketVersioningDisabledRule checks whether
// "scaleway_object_bucket" has versioning disabled.
type ScalewayObjectBucketVersioningDisabledRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewScalewayObjectBucketVersioningDisabledRule returns a new rule.
func NewScalewayObjectBucketVersioningDisabledRule() *ScalewayObjectBucketVersioningDisabledRule {
	return &ScalewayObjectBucketVersioningDisabledRule{resourceType: "scaleway_object_bucket"}
}

// Name returns the rule name.
func (r *ScalewayObjectBucketVersioningDisabledRule) Name() string {
	return "scaleway_object_bucket_versioning_disabled"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayObjectBucketVersioningDisabledRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayObjectBucketVersioningDisabledRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayObjectBucketVersioningDisabledRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_object_bucket" has versioning disabled.
func (r *ScalewayObjectBucketVersioningDisabledRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "versioning",
				Body: &hclext.BodySchema{Attributes: []hclext.AttributeSchema{{Name: "enabled"}}},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		versioning := resource.Body.Blocks.OfType("versioning")
		if len(versioning) == 0 {
			if err := runner.EmitIssue(
				r,
				"Versioning is not enabled on this bucket. Add a versioning block with enabled = true.",
				resource.DefRange,
			); err != nil {
				return err
			}
			continue
		}

		attribute, exists := versioning[0].Body.Attributes["enabled"]
		if !exists {
			continue
		}
		err := runner.EvaluateExpr(attribute.Expr, func(enabled bool) error {
			if enabled {
				return nil
			}
			return runner.EmitIssue(
				r,
				"Versioning is disabled on this bucket. Set enabled = true in the versioning block.",
				attribute.Expr.Range(),
			)
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
