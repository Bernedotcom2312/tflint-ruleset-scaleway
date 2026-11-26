# Scaleway Landing Zone Framework coverage

This page maps the guidelines of the Scaleway Landing Zone Framework (July 2026 edition) to the rules of this ruleset.

TFLint only reads Terraform configurations: guidelines about Organization settings, people or runtime operations cannot be checked statically and are listed as out of scope.

## Enable every Landing Zone rule

```hcl
plugin "scaleway" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/Bernedotcom2312/tflint-ruleset-scaleway"
}

rule "scaleway_instance_security_group_inbound_default_policy" { enabled = true }
rule "scaleway_instance_security_group_admin_port_exposed" { enabled = true }
rule "scaleway_rdb_instance_public_endpoint" { enabled = true }
rule "scaleway_redis_cluster_public_endpoint" { enabled = true }
rule "scaleway_rdb_instance_encryption_at_rest_disabled" { enabled = true }
rule "scaleway_rdb_instance_backup_disabled" { enabled = true }
rule "scaleway_lb_insecure_ssl_compatibility_level" { enabled = true }
rule "scaleway_object_bucket_versioning_disabled" { enabled = true }
rule "scaleway_iam_api_key_missing_expiration" { enabled = true }
rule "scaleway_iam_policy_full_access" { enabled = true }
rule "scaleway_iam_policy_not_attached_to_group" { enabled = true }

rule "scaleway_resource_missing_tags" {
  enabled = true
  tags    = ["env", "owner", "cost-center"]
}
```

## Coverage

✅ covered · 🟡 partially covered · ⬜ not covered yet · ➖ out of scope for a linter

### Networking

| Guideline | Coverage |
| --- | --- |
| Use the regions and zones of the platform (`fr-par`, `nl-ams`, `pl-waw`, `it-mil`) | ✅ [scaleway_invalid_region](rules/scaleway_invalid_region.md), [scaleway_invalid_zone](rules/scaleway_invalid_zone.md) |
| Deploy backend services and databases exclusively within Private Networks, without public IP | 🟡 [scaleway_rdb_instance_public_endpoint](rules/scaleway_rdb_instance_public_endpoint.md), [scaleway_redis_cluster_public_endpoint](rules/scaleway_redis_cluster_public_endpoint.md). Instances with a public IP are not checked yet. |
| Centralize outbound traffic through Public Gateways and inbound traffic through Load Balancers | ⬜ Requires cross-resource analysis. |
| Plan non-overlapping CIDR ranges (IPAM) before peering VPCs | ⬜ Requires cross-resource analysis. |
| Hub-and-spoke, VPC peering, custom routes, Network ACLs | ➖ Architecture choices. |

### Security & compliance

| Guideline | Coverage |
| --- | --- |
| Security groups: default-deny policy for inbound traffic | ✅ [scaleway_instance_security_group_inbound_default_policy](rules/scaleway_instance_security_group_inbound_default_policy.md). Note that the Terraform provider defaults to `accept`, unlike the console. |
| Never expose administration ports to the internet; use a bastion or a Public Gateway | ✅ [scaleway_instance_security_group_admin_port_exposed](rules/scaleway_instance_security_group_admin_port_exposed.md) |
| Restrict Kapsule control plane and Managed Database access to defined CIDRs | ⬜ Requires checking that a `scaleway_k8s_acl` / `scaleway_rdb_acl` targets each cluster or instance. |
| Load Balancers: TLS 1.2 and 1.3 only | ✅ [scaleway_lb_insecure_ssl_compatibility_level](rules/scaleway_lb_insecure_ssl_compatibility_level.md) |
| Load Balancers: redirect HTTP to HTTPS | ⬜ |
| Public applications behind Edge Services with the WAF in blocking mode | ⬜ |
| Enable DNSSEC on domains | ⬜ |
| Managed Databases: encryption at rest | ✅ [scaleway_rdb_instance_encryption_at_rest_disabled](rules/scaleway_rdb_instance_encryption_at_rest_disabled.md) |
| Object Storage: server-side encryption for sensitive buckets | ⬜ Depends on data classification. |
| Block volumes: LUKS encryption | ➖ Operating system level. |
| Secrets in Secret Manager, never in plain text in code | ⬜ Hardcoded passwords are not detected yet. |
| IAM: API keys with strict validity limits | ✅ [scaleway_iam_api_key_missing_expiration](rules/scaleway_iam_api_key_missing_expiration.md) |
| IAM: RBAC through groups, no direct permissions on users or applications | ✅ [scaleway_iam_policy_not_attached_to_group](rules/scaleway_iam_policy_not_attached_to_group.md) |
| IAM: least privilege, specific permission sets rather than `AllProductsFullAccess` | ✅ [scaleway_iam_policy_full_access](rules/scaleway_iam_policy_full_access.md) |
| IAM: scope policies to Projects rather than the Organization | ⬜ Organization-scoped permission sets (billing, IAM) make this noisy. |
| MFA, Owner account protection and recovery, identity federation | ➖ Organization settings. |
| Audit Trail, security monitoring, incident management | ➖ Organization settings and operations. |

### Automation

| Guideline | Coverage |
| --- | --- |
| Validate naming and tagging standards in CI/CD | 🟡 [scaleway_resource_missing_tags](rules/scaleway_resource_missing_tags.md). Naming conventions are not checked yet. |
| Remote state with versioning and locking | 🟡 [scaleway_object_bucket_versioning_disabled](rules/scaleway_object_bucket_versioning_disabled.md) for the state bucket; backend configuration is out of scope. |
| Reference Marketplace images by ID | ⬜ |
| Formatting checks, policy enforcement and security scanning before plan | ✅ This ruleset is meant to run in that step. |

### FinOps & GreenOps

| Guideline | Coverage |
| --- | --- |
| Cost allocation through tags | 🟡 [scaleway_resource_missing_tags](rules/scaleway_resource_missing_tags.md) |
| Right-sizing, budgets and alerts, Savings Plans | ➖ Billing and operations. |

### Disaster recovery

| Guideline | Coverage |
| --- | --- |
| Automated database backups, enabled at creation, at least daily | 🟡 [scaleway_rdb_instance_backup_disabled](rules/scaleway_rdb_instance_backup_disabled.md). The backup frequency is not checked yet. |
| Point-in-time recovery, backups isolated across AZs | ⬜ |
| Snapshot policies on every Block Storage volume | ⬜ |
| Bucket versioning for backup repositories | ✅ [scaleway_object_bucket_versioning_disabled](rules/scaleway_object_bucket_versioning_disabled.md) |
| RTO/RPO tiers, DR strategies, restore tests, Velero | ➖ Architecture and operations. |
