// Package rules contains the TFLint rules for terraform-provider-scaleway.
package rules

import "github.com/terraform-linters/tflint-plugin-sdk/tflint"

// Rules is the list of all rules provided by this ruleset.
var Rules = []tflint.Rule{
	NewScalewayInstanceServerInvalidTypeRule(),
	NewScalewayK8sPoolInvalidNodeTypeRule(),
	NewScalewayRdbInstanceInvalidNodeTypeRule(),
	NewScalewayRedisClusterInvalidNodeTypeRule(),
	NewScalewayLBInvalidTypeRule(),
	NewScalewayK8sClusterInvalidCNIRule(),
	NewScalewayRdbInstanceInvalidEngineRule(),
	NewScalewayInvalidZoneRule(),
	NewScalewayInvalidRegionRule(),
	NewScalewayObjectBucketInvalidNameRule(),
	NewScalewayInstanceSecurityGroupInboundDefaultPolicyRule(),
	NewScalewayRdbInstanceBackupDisabledRule(),
	NewScalewayInstanceSecurityGroupAdminPortExposedRule(),
	NewScalewayRdbInstancePublicEndpointRule(),
	NewScalewayRedisClusterPublicEndpointRule(),
	NewScalewayRdbInstanceEncryptionAtRestDisabledRule(),
	NewScalewayLBInsecureSSLCompatibilityLevelRule(),
	NewScalewayIAMAPIKeyMissingExpirationRule(),
	NewScalewayObjectBucketVersioningDisabledRule(),
}
