package models

import "testing"

func TestLookups(t *testing.T) {
	cases := []struct {
		name  string
		fn    func(string) bool
		value string
		want  bool
	}{
		{"instance type exact", IsValidInstanceServerType, "DEV1-S", true},
		{"instance type lower case", IsValidInstanceServerType, "pro2-xxs", true},
		{"instance type unknown", IsValidInstanceServerType, "DEV1-XXL", false},
		{"k8s node type exact", IsValidK8sNodeType, "PRO2-S", true},
		{"k8s node type underscore", IsValidK8sNodeType, "pro2_s", true},
		{"k8s node type external", IsValidK8sNodeType, "external", true},
		{"k8s node type unknown", IsValidK8sNodeType, "t2.micro", false},
		{"rdb node type", IsValidRdbNodeType, "DB-DEV-S", true},
		{"rdb node type unknown", IsValidRdbNodeType, "db-dev-xxs", false},
		{"redis node type", IsValidRedisNodeType, "RED1-micro", true},
		{"redis node type unknown", IsValidRedisNodeType, "RED2-S", false},
		{"lb type", IsValidLBType, "lb-gp-m", true},
		{"lb type unknown", IsValidLBType, "LB-XXL", false},
		{"cni", IsValidK8sCNI, "cilium", true},
		{"cni is case sensitive", IsValidK8sCNI, "Cilium", false},
		{"zone", IsValidZone, "fr-par-2", true},
		{"zone unknown", IsValidZone, "fr-par-4", false},
		{"region", IsValidRegion, "nl-ams", true},
		{"region given a zone", IsValidRegion, "nl-ams-1", false},
		{"taggable resource", IsTaggableResource, "scaleway_instance_server", true},
		{"taggable framework resource", IsTaggableResource, "scaleway_block_volume", true},
		{"resource without tags", IsTaggableResource, "scaleway_rdb_acl", false},
		{"map tags", HasMapTags, "scaleway_object_bucket", true},
		{"list tags", HasMapTags, "scaleway_k8s_cluster", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.fn(tc.value); got != tc.want {
				t.Errorf("got %v for %q, want %v", got, tc.value, tc.want)
			}
		})
	}
}
