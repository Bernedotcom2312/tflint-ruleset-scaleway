//go:build e2e

package integration

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"reflect"
	"sort"
	"testing"
)

type result struct {
	Issues []issue           `json:"issues"`
	Errors []json.RawMessage `json:"errors"`
}

type issue struct {
	Rule struct {
		Name     string `json:"name"`
		Severity string `json:"severity"`
	} `json:"rule"`
	Message string `json:"message"`
	Range   struct {
		Filename string `json:"filename"`
		Start    struct {
			Line   int `json:"line"`
			Column int `json:"column"`
		} `json:"start"`
	} `json:"range"`
}

// expectedIssue is the subset of an issue compared by the tests.
type expectedIssue struct {
	Rule     string
	Severity string
	Message  string
	Filename string
	Line     int
	Column   int
}

func TestIntegration(t *testing.T) {
	if _, err := exec.LookPath("tflint"); err != nil {
		t.Fatal("tflint must be installed to run the integration tests")
	}

	cases := []struct {
		Name     string
		Dir      string
		Expected []expectedIssue
	}{
		{
			Name: "rules enabled by default",
			Dir:  "basic",
			Expected: []expectedIssue{
				{"scaleway_invalid_region", "warning", `"fr-par-1" is an invalid Scaleway region.`, "main.tf", 3, 12},
				{"scaleway_instance_server_invalid_type", "error", `"DEV1-XXL" is an invalid Instance commercial type.`, "main.tf", 7, 11},
				{"scaleway_invalid_zone", "warning", `"fr-par-4" is an invalid Scaleway zone.`, "main.tf", 9, 11},
				{"scaleway_k8s_cluster_invalid_cni", "error", `"flannel" is an invalid Kubernetes CNI plugin.`, "main.tf", 15, 33},
				{"scaleway_k8s_pool_invalid_node_type", "error", `"t2.micro" is an invalid Kubernetes pool node type.`, "main.tf", 22, 16},
				{"scaleway_rdb_instance_invalid_node_type", "error", `"db-dev-xxs" is an invalid Managed Database node type.`, "main.tf", 28, 15},
				{"scaleway_rdb_instance_invalid_engine", "error", `"postgres-16" is an invalid Managed Database engine. It must look like "PostgreSQL-16" or "MySQL-8".`, "main.tf", 29, 15},
				{"scaleway_redis_cluster_invalid_node_type", "error", `"RED1-XXL" is an invalid Managed Redis node type.`, "main.tf", 35, 15},
				{"scaleway_lb_invalid_type", "error", `"LB-XXL" is an invalid Load Balancer commercial type.`, "main.tf", 42, 10},
				{"scaleway_object_bucket_invalid_name", "error", `"My_Logs" is an invalid bucket name: it can only contain lowercase letters, numbers, dots (.) and hyphens (-).`, "main.tf", 46, 10},
			},
		},
		{
			Name: "best practice rules enabled in .tflint.hcl",
			Dir:  "best-practices",
			Expected: []expectedIssue{
				{"scaleway_instance_security_group_inbound_default_policy", "warning", `inbound_default_policy is not set and defaults to "accept". Set it to "drop" and allow traffic with explicit inbound rules.`, "main.tf", 1, 1},
				{"scaleway_instance_security_group_inbound_default_policy", "warning", `inbound_default_policy is "accept". Set it to "drop" and allow traffic with explicit inbound rules.`, "main.tf", 7, 28},
				{"scaleway_rdb_instance_backup_disabled", "warning", "Automated backups are disabled for this Database Instance.", "main.tf", 24, 20},
			},
		},
		{
			Name: "landing zone rules enabled in .tflint.hcl",
			Dir:  "landing-zone",
			Expected: []expectedIssue{
				{"scaleway_resource_missing_tags", "info", `The resource is missing the following tags: "owner".`, "main.tf", 12, 28},
				{"scaleway_instance_security_group_admin_port_exposed", "warning", "Inbound rule accepts administration port 22 from 0.0.0.0/0. Reach Instances through the Public Gateway SSH bastion or a VPN instead.", "main.tf", 14, 3},
				{"scaleway_rdb_instance_public_endpoint", "warning", "The Database Instance has no private_network block, so it gets a public endpoint by default. Attach it to a Private Network.", "main.tf", 20, 1},
				{"scaleway_rdb_instance_encryption_at_rest_disabled", "warning", "Encryption at rest is not enabled for this Database Instance. Set encryption_at_rest = true.", "main.tf", 20, 1},
				{"scaleway_redis_cluster_public_endpoint", "warning", "The Redis cluster has no private_network block, so it is exposed on a public endpoint. Attach it to a Private Network.", "main.tf", 40, 1},
				{"scaleway_lb_insecure_ssl_compatibility_level", "warning", "ssl_compatibility_level_old accepts TLS versions older than 1.2. Use ssl_compatibility_level_intermediate (the default) or ssl_compatibility_level_modern.", "main.tf", 52, 29},
				{"scaleway_object_bucket_versioning_disabled", "warning", "Versioning is not enabled on this bucket. Add a versioning block with enabled = true.", "main.tf", 56, 1},
				{"scaleway_iam_api_key_missing_expiration", "warning", "The API key never expires. Set expires_at and rotate the key before that date.", "main.tf", 69, 1},
				{"scaleway_iam_policy_not_attached_to_group", "warning", "The policy is attached directly to an application. Attach it to an IAM group (group_id) and add the application to the group.", "main.tf", 74, 3},
				{"scaleway_iam_policy_full_access", "warning", "AllProductsFullAccess grants every permission on every product. Use product-specific permission sets, such as InstancesFullAccess.", "main.tf", 79, 28},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			cmd := exec.Command("tflint", "--format", "json", "--force")
			cmd.Dir = tc.Dir
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if err := cmd.Run(); err != nil {
				t.Fatalf("tflint failed: %s\nstdout: %s\nstderr: %s", err, stdout.String(), stderr.String())
			}

			var got result
			if err := json.Unmarshal(stdout.Bytes(), &got); err != nil {
				t.Fatalf("failed to parse the tflint output: %s\n%s", err, stdout.String())
			}
			if len(got.Errors) > 0 {
				t.Fatalf("tflint reported errors: %s", stdout.String())
			}

			issues := make([]expectedIssue, 0, len(got.Issues))
			for _, i := range got.Issues {
				issues = append(issues, expectedIssue{
					Rule:     i.Rule.Name,
					Severity: i.Rule.Severity,
					Message:  i.Message,
					Filename: i.Range.Filename,
					Line:     i.Range.Start.Line,
					Column:   i.Range.Start.Column,
				})
			}
			sortIssues(issues)
			sortIssues(tc.Expected)

			if !reflect.DeepEqual(issues, tc.Expected) {
				t.Errorf("unexpected issues:\ngot:  %+v\nwant: %+v", issues, tc.Expected)
			}
		})
	}
}

func sortIssues(issues []expectedIssue) {
	sort.Slice(issues, func(i, j int) bool {
		a, b := issues[i], issues[j]
		if a.Filename != b.Filename {
			return a.Filename < b.Filename
		}
		if a.Line != b.Line {
			return a.Line < b.Line
		}
		if a.Column != b.Column {
			return a.Column < b.Column
		}
		return a.Rule < b.Rule
	})
}
