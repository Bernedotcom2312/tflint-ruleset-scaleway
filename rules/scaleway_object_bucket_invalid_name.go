package rules

import (
	"fmt"
	"net"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

var (
	bucketNameCharsRegexp = regexp.MustCompile(`^[a-z0-9.-]+$`)
	bucketNameEdgesRegexp = regexp.MustCompile(`^[a-z0-9](.*[a-z0-9])?$`)
)

// bucketNameError returns why name is not a valid bucket name, or an empty
// string if it is valid. Scaleway Object Storage follows the DNS-compatible
// naming rules of Amazon S3.
func bucketNameError(name string) string {
	switch {
	case len(name) < 3 || len(name) > 63:
		return "it must be between 3 and 63 characters long"
	case !bucketNameCharsRegexp.MatchString(name):
		return "it can only contain lowercase letters, numbers, dots (.) and hyphens (-)"
	case !bucketNameEdgesRegexp.MatchString(name):
		return "it must begin and end with a lowercase letter or a number"
	case net.ParseIP(name) != nil:
		return "it must not be formatted as an IP address"
	default:
		return ""
	}
}

// ScalewayObjectBucketInvalidNameRule checks whether "scaleway_object_bucket" has an invalid name.
type ScalewayObjectBucketInvalidNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayObjectBucketInvalidNameRule returns a new rule.
func NewScalewayObjectBucketInvalidNameRule() *ScalewayObjectBucketInvalidNameRule {
	return &ScalewayObjectBucketInvalidNameRule{
		resourceType:  "scaleway_object_bucket",
		attributeName: "name",
	}
}

// Name returns the rule name.
func (r *ScalewayObjectBucketInvalidNameRule) Name() string {
	return "scaleway_object_bucket_invalid_name"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayObjectBucketInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayObjectBucketInvalidNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayObjectBucketInvalidNameRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_object_bucket" has an invalid name.
func (r *ScalewayObjectBucketInvalidNameRule) Check(runner tflint.Runner) error {
	valid := func(name string) bool { return bucketNameError(name) == "" }

	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, valid, func(name string) string {
		return fmt.Sprintf(`"%s" is an invalid bucket name: %s.`, name, bucketNameError(name))
	})
}
