// Package rules contains the TFLint rules for terraform-provider-scaleway.
package rules

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

// Rules is the list of all rules provided by this ruleset.
var Rules = []tflint.Rule{
	NewScalewayInstanceServerInvalidTypeRule(),
	NewScalewayK8sPoolInvalidNodeTypeRule(),
	NewScalewayRdbInstanceInvalidNodeTypeRule(),
}
