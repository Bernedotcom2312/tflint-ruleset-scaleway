package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayInvalidRegionRule checks whether Scaleway blocks use an unknown region.
type ScalewayInvalidRegionRule struct {
	localityRule
}

// NewScalewayInvalidRegionRule returns a new rule.
func NewScalewayInvalidRegionRule() *ScalewayInvalidRegionRule {
	return &ScalewayInvalidRegionRule{
		localityRule: localityRule{
			name:          "scaleway_invalid_region",
			attributeName: "region",
			valid:         models.IsValidRegion,
		},
	}
}

// Check checks whether Scaleway blocks use an unknown region.
func (r *ScalewayInvalidRegionRule) Check(runner tflint.Runner) error {
	return r.check(runner, r)
}
