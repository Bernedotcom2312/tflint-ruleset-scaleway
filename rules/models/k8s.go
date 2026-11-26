package models

// K8sCNIs lists the CNI plugins supported by Kubernetes Kapsule and Kosmos.
//
// Kept in sync with the k8s.CNI enum from github.com/scaleway/scaleway-sdk-go
// (the "unknown_cni" placeholder is intentionally left out).
var K8sCNIs = []string{
	"cilium",
	"cilium_native",
	"calico",
	"kilo",
	"none",
}

// K8sExternalNodeType is the node type of Kosmos pools whose nodes are
// managed outside of Scaleway.
const K8sExternalNodeType = "external"
