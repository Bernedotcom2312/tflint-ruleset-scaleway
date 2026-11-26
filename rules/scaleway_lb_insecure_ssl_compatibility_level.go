package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayLBInsecureSSLCompatibilityLevelRule checks whether "scaleway_lb" accepts TLS versions older than 1.2.
type ScalewayLBInsecureSSLCompatibilityLevelRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayLBInsecureSSLCompatibilityLevelRule returns a new rule.
func NewScalewayLBInsecureSSLCompatibilityLevelRule() *ScalewayLBInsecureSSLCompatibilityLevelRule {
	return &ScalewayLBInsecureSSLCompatibilityLevelRule{
		resourceType:  "scaleway_lb",
		attributeName: "ssl_compatibility_level",
	}
}

// Name returns the rule name.
func (r *ScalewayLBInsecureSSLCompatibilityLevelRule) Name() string {
	return "scaleway_lb_insecure_ssl_compatibility_level"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayLBInsecureSSLCompatibilityLevelRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayLBInsecureSSLCompatibilityLevelRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayLBInsecureSSLCompatibilityLevelRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_lb" accepts TLS versions older than 1.2.
func (r *ScalewayLBInsecureSSLCompatibilityLevelRule) Check(runner tflint.Runner) error {
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

		err := runner.EvaluateExpr(attribute.Expr, func(value string) error {
			if value != "ssl_compatibility_level_old" {
				return nil
			}
			return runner.EmitIssue(r, "ssl_compatibility_level_old accepts TLS versions older than 1.2. Use ssl_compatibility_level_intermediate (the default) or ssl_compatibility_level_modern.", attribute.Expr.Range())
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
