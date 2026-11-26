package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// ScalewayIAMPolicyNotAttachedToGroupRule checks whether "scaleway_iam_policy"
// is attached directly to a user or an application instead of a group.
type ScalewayIAMPolicyNotAttachedToGroupRule struct {
	tflint.DefaultRule

	resourceType string
	// principals maps the principal attributes to their name, with and without article.
	principals map[string][2]string
}

// NewScalewayIAMPolicyNotAttachedToGroupRule returns a new rule.
func NewScalewayIAMPolicyNotAttachedToGroupRule() *ScalewayIAMPolicyNotAttachedToGroupRule {
	return &ScalewayIAMPolicyNotAttachedToGroupRule{
		resourceType: "scaleway_iam_policy",
		principals: map[string][2]string{
			"user_id":        {"a user", "user"},
			"application_id": {"an application", "application"},
		},
	}
}

// Name returns the rule name.
func (r *ScalewayIAMPolicyNotAttachedToGroupRule) Name() string {
	return "scaleway_iam_policy_not_attached_to_group"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayIAMPolicyNotAttachedToGroupRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayIAMPolicyNotAttachedToGroupRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayIAMPolicyNotAttachedToGroupRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_iam_policy" is attached to a user or an application.
func (r *ScalewayIAMPolicyNotAttachedToGroupRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: "user_id"}, {Name: "application_id"}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, name := range []string{"user_id", "application_id"} {
			attribute, exists := resource.Body.Attributes[name]
			if !exists {
				continue
			}
			if err := runner.EmitIssue(
				r,
				fmt.Sprintf("The policy is attached directly to %s. Attach it to an IAM group (group_id) and add the %s to the group.", r.principals[name][0], r.principals[name][1]),
				attribute.Range,
			); err != nil {
				return err
			}
		}
	}

	return nil
}
