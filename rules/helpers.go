package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// checkStringAttribute evaluates attributeName on every resourceType
// resource of the module and emits an issue for each value rejected by valid.
// Unknown, null and non-string values are skipped.
func checkStringAttribute(
	runner tflint.Runner,
	rule tflint.Rule,
	resourceType, attributeName string,
	valid func(string) bool,
	message func(string) string,
) error {
	resources, err := runner.GetResourceContent(resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: attributeName}},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(value string) error {
			if valid(value) {
				return nil
			}
			return runner.EmitIssue(rule, message(value), attribute.Expr.Range())
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}
