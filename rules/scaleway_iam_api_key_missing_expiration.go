package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayIAMAPIKeyMissingExpirationRule checks whether "scaleway_iam_api_key" has no expiration date.
type ScalewayIAMAPIKeyMissingExpirationRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayIAMAPIKeyMissingExpirationRule returns a new rule.
func NewScalewayIAMAPIKeyMissingExpirationRule() *ScalewayIAMAPIKeyMissingExpirationRule {
	return &ScalewayIAMAPIKeyMissingExpirationRule{
		resourceType:  "scaleway_iam_api_key",
		attributeName: "expires_at",
	}
}

// Name returns the rule name.
func (r *ScalewayIAMAPIKeyMissingExpirationRule) Name() string {
	return "scaleway_iam_api_key_missing_expiration"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayIAMAPIKeyMissingExpirationRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayIAMAPIKeyMissingExpirationRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayIAMAPIKeyMissingExpirationRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_iam_api_key" has no expiration date.
func (r *ScalewayIAMAPIKeyMissingExpirationRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		if _, exists := resource.Body.Attributes[r.attributeName]; exists {
			continue
		}
		if err := runner.EmitIssue(r, "The API key never expires. Set expires_at and rotate the key before that date.", resource.DefRange); err != nil {
			return err
		}
	}

	return nil
}
