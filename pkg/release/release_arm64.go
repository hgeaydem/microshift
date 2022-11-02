/*
Copyright © 2021 MicroShift Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package release

// For the amd64 architecture we use the existing and tested and
// published OCP or other component upstream images

func init() {
	Image = map[string]string{
		"cli":                       "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:6aed4b901ff2456c803ff159b50021bd86c6473c0c65f76c39269fb330640d1b",
		"coredns":                   "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:be116df9f46266e604f1028fbde3742e3d63fef7c6421fb5a2a3673ce4f971e7",
		"haproxy_router":            "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:4f5f51a8ba723590e550bc2024eaea6554bb4b749ec4356229060d3bd92fc9eb",
		"kube_rbac_proxy":           "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:51ce83de9a43dfe8731c532f4ed0402c70da8df1e056292190c2cb43a45232a7",
		"openssl":                   "registry.access.redhat.com/ubi8/openssl@sha256:3f781a07e59d164eba065dba7d8e7661ab2494b21199c379b65b0ff514a1b8d0",
		"ovn_kubernetes_microshift": "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:d667c28d9059afbf3c20fc69bf0a215803e497df2e92998d5d5b13a2296b74f7",
		"pod":                       "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f310397d0a61f0ffefbb5753cb72f271d483eac82f4dd676d592fe2ea8487a30",
		"service_ca_operator":       "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f41255221e5dafd5688684bc0446fedf4a7526b3a02e05ca787bd166449088f8",
		"odf_topolvm":               "quay.io/rhceph-dev/odf4-odf-topolvm-rhel8@sha256:2855918d1849c99a835eb03c53ce07170c238111fd15d2fe50cd45611fcd1ceb",
		"ose_csi_ext_provisioner":   "quay.io/rhceph-dev/openshift-ose-csi-external-provisioner@sha256:c3b2417f8fcb8883275f0e613037f83133ccc3f91311a30688e4be520544ea4a",
		"ose_csi_ext_resizer":       "quay.io/rhceph-dev/openshift-ose-csi-external-resizer@sha256:213f43d61b3a214a4a433c7132537be082a108d55005f2ba0777c2ea97489799",
		"topolvm-csi-snapshotter":   "quay.io/rhceph-dev/openshift-ose-csi-external-snapshotter@sha256:734c095670d21b77f18c84670d6c9a7742be1d9151dca0da20f41858ede65ed8",
		"ose_csi_livenessprobe":     "quay.io/rhceph-dev/openshift-ose-csi-livenessprobe@sha256:b05559aa038708ab448cfdfed2ca880726aed6cc30371fea4d6a42c972c0c728",
		"ose_csi_node_registrar":    "quay.io/rhceph-dev/openshift-ose-csi-node-driver-registrar@sha256:fb0f5e531847db94dcadc61446b9a892f6f92ddf282e192abf2fdef6c6af78f2",
	}
}
