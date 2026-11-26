package rules

import (
	"fmt"
	"sort"
	"strings"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayResourceMissingTagsRule checks whether taggable Scaleway resources
// carry the tags required by the configuration.
type ScalewayResourceMissingTagsRule struct {
	tflint.DefaultRule
}

type resourceMissingTagsRuleConfig struct {
	Tags    []string `hclext:"tags"`
	Exclude []string `hclext:"exclude,optional"`
}

// NewScalewayResourceMissingTagsRule returns a new rule.
func NewScalewayResourceMissingTagsRule() *ScalewayResourceMissingTagsRule {
	return &ScalewayResourceMissingTagsRule{}
}

// Name returns the rule name.
func (r *ScalewayResourceMissingTagsRule) Name() string {
	return "scaleway_resource_missing_tags"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayResourceMissingTagsRule) Enabled() bool {
	return false
}

// Severity returns the rule severity.
func (r *ScalewayResourceMissingTagsRule) Severity() tflint.Severity {
	return tflint.NOTICE
}

// Link returns the rule reference link.
func (r *ScalewayResourceMissingTagsRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether taggable Scaleway resources carry the required tags.
func (r *ScalewayResourceMissingTagsRule) Check(runner tflint.Runner) error {
	config := resourceMissingTagsRuleConfig{}
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}
	if len(config.Tags) == 0 {
		return nil
	}

	excluded := map[string]bool{}
	for _, resourceType := range config.Exclude {
		excluded[resourceType] = true
	}

	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "resource",
				LabelNames: []string{"type", "name"},
				Body:       &hclext.BodySchema{Attributes: []hclext.AttributeSchema{{Name: "tags"}}},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range content.Blocks {
		resourceType := resource.Labels[0]
		if !models.IsTaggableResource(resourceType) || excluded[resourceType] {
			continue
		}

		attribute, exists := resource.Body.Attributes["tags"]
		if !exists {
			if err := r.emit(runner, config.Tags, resource.DefRange); err != nil {
				return err
			}
			continue
		}

		var callback any = func(tags []string) error {
			return r.emitMissing(runner, config.Tags, hasListTag(tags), attribute)
		}
		if models.HasMapTags(resourceType) {
			callback = func(tags map[string]string) error {
				return r.emitMissing(runner, config.Tags, hasMapTag(tags), attribute)
			}
		}
		if err := runner.EvaluateExpr(attribute.Expr, callback, nil); err != nil {
			return err
		}
	}

	return nil
}

func (r *ScalewayResourceMissingTagsRule) emitMissing(runner tflint.Runner, required []string, has func(string) bool, attribute *hclext.Attribute) error {
	var missing []string
	for _, key := range required {
		if !has(key) {
			missing = append(missing, key)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return r.emit(runner, missing, attribute.Expr.Range())
}

func (r *ScalewayResourceMissingTagsRule) emit(runner tflint.Runner, missing []string, issueRange hcl.Range) error {
	sorted := append([]string{}, missing...)
	sort.Strings(sorted)

	quoted := make([]string, 0, len(sorted))
	for _, key := range sorted {
		quoted = append(quoted, fmt.Sprintf("%q", key))
	}

	return runner.EmitIssue(
		r,
		fmt.Sprintf("The resource is missing the following tags: %s.", strings.Join(quoted, ", ")),
		issueRange,
	)
}

// hasListTag matches a key against Scaleway list tags, where a key is either
// a bare tag ("prod") or prefixes a value ("env=prod" or "env:prod").
func hasListTag(tags []string) func(string) bool {
	return func(key string) bool {
		for _, tag := range tags {
			if tag == key || strings.HasPrefix(tag, key+"=") || strings.HasPrefix(tag, key+":") {
				return true
			}
		}
		return false
	}
}

// hasMapTag matches a key against map tags (Object Storage buckets).
func hasMapTag(tags map[string]string) func(string) bool {
	return func(key string) bool {
		_, ok := tags[key]
		return ok
	}
}
