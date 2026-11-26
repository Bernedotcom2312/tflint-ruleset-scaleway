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

// evaluateString evaluates attribute as a string. known is false when the
// value cannot be determined statically (unknown, null, sensitive...).
func evaluateString(runner tflint.Runner, attribute *hclext.Attribute) (value string, known bool, err error) {
	err = runner.EvaluateExpr(attribute.Expr, func(v string) error {
		value, known = v, true
		return nil
	}, nil)
	return value, known, err
}

// evaluateInt evaluates attribute as an integer. known is false when the
// value cannot be determined statically.
func evaluateInt(runner tflint.Runner, attribute *hclext.Attribute) (value int, known bool, err error) {
	err = runner.EvaluateExpr(attribute.Expr, func(v int) error {
		value, known = v, true
		return nil
	}, nil)
	return value, known, err
}
