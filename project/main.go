// Package project holds metadata shared by the ruleset.
package project

import "fmt"

// Version is the ruleset version.
const Version string = "0.1.0"

// ReferenceLink returns the documentation link of a rule.
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v%s/docs/rules/%s.md", Version, name)
}
