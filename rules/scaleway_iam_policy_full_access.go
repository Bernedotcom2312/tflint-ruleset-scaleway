package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// fullAccessPermissionSet grants every permission on every product.
const fullAccessPermissionSet = "AllProductsFullAccess"

// ScalewayIAMPolicyFullAccessRule checks whether "scaleway_iam_policy"
// grants the AllProductsFullAccess permission set.
type ScalewayIAMPolicyFullAccessRule struct {
	tflint.DefaultRule

	resourceType string
}

// NewScalewayIAMPolicyFullAccessRule returns a new rule.
func NewScalewayIAMPolicyFullAccessRule() *ScalewayIAMPolicyFullAccessRule {
	return &ScalewayIAMPolicyFullAccessRule{resourceType: "scaleway_iam_policy"}
}

// Name returns the rule name.
func (r *ScalewayIAMPolicyFullAccessRule) Name() string {
	return "scaleway_iam_policy_full_access"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayIAMPolicyFullAccessRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayIAMPolicyFullAccessRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *ScalewayIAMPolicyFullAccessRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_iam_policy" grants AllProductsFullAccess.
func (r *ScalewayIAMPolicyFullAccessRule) Check(runner tflint.Runner) error {
	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "rule",
				Body: &hclext.BodySchema{Attributes: []hclext.AttributeSchema{{Name: "permission_set_names"}}},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		for _, rule := range resource.Body.Blocks {
			attribute, exists := rule.Body.Attributes["permission_set_names"]
			if !exists {
				continue
			}

			err := runner.EvaluateExpr(attribute.Expr, func(names []string) error {
				for _, name := range names {
					if name == fullAccessPermissionSet {
						return runner.EmitIssue(
							r,
							"AllProductsFullAccess grants every permission on every product. Use product-specific permission sets, such as InstancesFullAccess.",
							attribute.Expr.Range(),
						)
					}
				}
				return nil
			}, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
