package rules

import (
	"fmt"
	"regexp"

	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/Bernedotcom2312/tflint-ruleset-scaleway/project"
)

// rdbEngineRegexp matches the "<engine>-<version>" format expected by the
// Managed Database API. The provider compares engines case-insensitively.
var rdbEngineRegexp = regexp.MustCompile(`(?i)^(PostgreSQL|MySQL)-[0-9]+(\.[0-9]+)?$`)

func isValidRdbEngine(engine string) bool {
	return rdbEngineRegexp.MatchString(engine)
}

// ScalewayRdbInstanceInvalidEngineRule checks whether "scaleway_rdb_instance" has an invalid engine.
type ScalewayRdbInstanceInvalidEngineRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
}

// NewScalewayRdbInstanceInvalidEngineRule returns a new rule.
func NewScalewayRdbInstanceInvalidEngineRule() *ScalewayRdbInstanceInvalidEngineRule {
	return &ScalewayRdbInstanceInvalidEngineRule{
		resourceType:  "scaleway_rdb_instance",
		attributeName: "engine",
	}
}

// Name returns the rule name.
func (r *ScalewayRdbInstanceInvalidEngineRule) Name() string {
	return "scaleway_rdb_instance_invalid_engine"
}

// Enabled returns whether the rule is enabled by default.
func (r *ScalewayRdbInstanceInvalidEngineRule) Enabled() bool {
	return true
}

// Severity returns the rule severity.
func (r *ScalewayRdbInstanceInvalidEngineRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link.
func (r *ScalewayRdbInstanceInvalidEngineRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether "scaleway_rdb_instance" has an invalid engine.
func (r *ScalewayRdbInstanceInvalidEngineRule) Check(runner tflint.Runner) error {
	return checkStringAttribute(runner, r, r.resourceType, r.attributeName, isValidRdbEngine, func(value string) string {
		return fmt.Sprintf(`"%s" is an invalid Managed Database engine. It must look like "PostgreSQL-16" or "MySQL-8".`, value)
	})
}
