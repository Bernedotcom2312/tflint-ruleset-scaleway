package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayK8sClusterInvalidCNIRule checks whether "scaleway_k8s_cluster" has an invalid CNI plugin.
type ScalewayK8sClusterInvalidCNIRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayK8sClusterInvalidCNIRule returns a new rule.
func NewScalewayK8sClusterInvalidCNIRule() *ScalewayK8sClusterInvalidCNIRule {
	return &ScalewayK8sClusterInvalidCNIRule{
		resourceType:  "scaleway_k8s_cluster",
		attributeName: "cni",
	}
}

// Name returns the rule name.
func (r *ScalewayK8sClusterInvalidCNIRule) Name() string {
	return "scaleway_k8s_cluster_invalid_cni"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayK8sClusterInvalidCNIRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayK8sClusterInvalidCNIRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayK8sClusterInvalidCNIRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_k8s_cluster" has an invalid CNI plugin.
func (r *ScalewayK8sClusterInvalidCNIRule) Check(runner tflint.Runner) error {
	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, models.IsValidK8sCNI, func(value string) string {
		return fmt.Sprintf(`"%s" is an invalid Kubernetes CNI plugin.`, value)
	})
}
