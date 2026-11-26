# scaleway_k8s_cluster_invalid_cni

Disallow an invalid CNI plugin on `scaleway_k8s_cluster`.

## Example

```hcl
resource "scaleway_k8s_cluster" "main" {
  name                        = "main"
  version                     = "1.35"
  delete_additional_resources = true
  cni                         = "flannel" // invalid CNI!
}
```

```
$ tflint
1 issue(s) found:

Error: "flannel" is an invalid Kubernetes CNI plugin. (scaleway_k8s_cluster_invalid_cni)

  on main.tf line 5:
   5:   cni                         = "flannel" // invalid CNI!

Reference: https://github.com/Bernedotcom2312/tflint-ruleset-scaleway/blob/v0.1.0/docs/rules/scaleway_k8s_cluster_invalid_cni.md
```

## Why

The provider only emits a warning for an unknown `cni`, and the cluster creation then fails at apply time. Since `cni` forces a new cluster, a typo can also turn a harmless change into a failed replacement.

## How To Fix

Use one of `cilium`, `cilium_native`, `calico`, `kilo` or `none`. The value is case-sensitive.
