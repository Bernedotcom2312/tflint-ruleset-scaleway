// Package models holds the Scaleway catalog data used by the rules.
package models

import "strings"

// set is a lookup table of normalized values.
type set map[string]struct{}

func newSet(values []string, normalize func(string) string) set {
	s := make(set, len(values))
	for _, v := range values {
		s[normalize(v)] = struct{}{}
	}
	return s
}

func (s set) has(v string) bool {
	_, ok := s[v]
	return ok
}

func identity(s string) string { return s }

// normalizeNodeType mirrors the provider's IgnoreCaseAndHyphen diff
// suppression: "PRO2-S", "pro2-s" and "pro2_s" are equivalent.
func normalizeNodeType(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), "-", "_")
}

var (
	instanceServerTypes = newSet(InstanceServerTypes, strings.ToLower)
	k8sNodeTypes        = newSet(append(append([]string{}, InstanceServerTypes...), K8sExternalNodeType), normalizeNodeType)
	rdbNodeTypes        = newSet(RdbNodeTypes, strings.ToLower)
	redisNodeTypes      = newSet(RedisNodeTypes, strings.ToLower)
	lbTypes             = newSet(LBTypes, strings.ToLower)
	k8sCNIs             = newSet(K8sCNIs, identity)
	zones               = newSet(Zones, identity)
	regions             = newSet(Regions, identity)
)

// IsValidInstanceServerType reports whether t is a known Instance commercial type.
// The comparison is case-insensitive, like the provider.
func IsValidInstanceServerType(t string) bool {
	return instanceServerTypes.has(strings.ToLower(t))
}

// IsValidK8sNodeType reports whether t is a known Kubernetes pool node type.
// Case and "-"/"_" differences are ignored, like the provider.
func IsValidK8sNodeType(t string) bool {
	return k8sNodeTypes.has(normalizeNodeType(t))
}

// IsValidRdbNodeType reports whether t is a known Managed Database node type.
func IsValidRdbNodeType(t string) bool {
	return rdbNodeTypes.has(strings.ToLower(t))
}

// IsValidRedisNodeType reports whether t is a known Managed Redis node type.
func IsValidRedisNodeType(t string) bool {
	return redisNodeTypes.has(strings.ToLower(t))
}

// IsValidLBType reports whether t is a known Load Balancer commercial type.
func IsValidLBType(t string) bool {
	return lbTypes.has(strings.ToLower(t))
}

// IsValidK8sCNI reports whether cni is a supported Kubernetes CNI plugin.
func IsValidK8sCNI(cni string) bool {
	return k8sCNIs.has(cni)
}

// IsValidZone reports whether zone is a known Scaleway Availability Zone.
func IsValidZone(zone string) bool {
	return zones.has(zone)
}

// IsValidRegion reports whether region is a known Scaleway region.
func IsValidRegion(region string) bool {
	return regions.has(region)
}
