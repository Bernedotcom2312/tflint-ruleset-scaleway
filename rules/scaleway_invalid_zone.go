package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/rules/models"
)

// ScalewayInvalidZoneRule checks whether Scaleway blocks use an unknown zone.
type ScalewayInvalidZoneRule struct {
	localityRule
}

// NewScalewayInvalidZoneRule returns a new rule.
func NewScalewayInvalidZoneRule() *ScalewayInvalidZoneRule {
	return &ScalewayInvalidZoneRule{
		localityRule: localityRule{
			name:          "scaleway_invalid_zone",
			attributeName: "zone",
			valid:         models.IsValidZone,
		},
	}
}

// Check checks whether Scaleway blocks use an unknown zone.
func (r *ScalewayInvalidZoneRule) Check(runner tflint.Runner) error {
	return r.check(runner, r)
}
