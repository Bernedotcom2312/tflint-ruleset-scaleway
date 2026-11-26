package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// localityRule checks a locality attribute ("zone" or "region") on every
// Scaleway resource, data source and provider block of the module.
type localityRule struct {
	tflint.DefaultRule

	name          string
	attributeName string
	valid         func(string) bool
}

// Name returns the rule name.
func (r *localityRule) Name() string {
	return r.name
}

// Enabled returns whether the rule is enabled by default.
func (r *localityRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
//
// The provider itself only warns on unknown localities, so this rule does
// the same: a newly opened locality may not be known by this ruleset yet.
func (r *localityRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link.
func (r *localityRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// check checks whether Scaleway blocks use an unknown locality. Issues are
// emitted on behalf of self, the concrete rule embedding localityRule.
func (r *localityRule) check(runner tflint.Runner, self tflint.Rule) error {
	body := &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{{Name: r.attributeName}},
	}

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{Type: "resource", LabelNames: []string{"type", "name"}, Body: body},
			{Type: "data", LabelNames: []string{"type", "name"}, Body: body},
			{Type: "provider", LabelNames: []string{"name"}, Body: body},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, block := range content.Blocks {
		if !isScalewayBlock(block) {
			continue
		}

		attribute, exists := block.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		err := runner.EvaluateExpr(attribute.Expr, func(value string) error {
			if r.valid(value) {
				return nil
			}
			return runner.EmitIssue(
				self,
				fmt.Sprintf(`"%s" is an invalid Scaleway %s.`, value, r.attributeName),
				attribute.Expr.Range(),
			)
		}, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// isScalewayBlock reports whether a resource, data or provider block belongs
// to the Scaleway provider.
func isScalewayBlock(block *hclext.Block) bool {
	if len(block.Labels) == 0 {
		return false
	}
	if block.Type == "provider" {
		return block.Labels[0] == "scaleway"
	}
	return strings.HasPrefix(block.Labels[0], "scaleway_")
}
