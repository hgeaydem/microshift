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
		"cli":                       "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:b456a5bd94290f1a875c17419993e30a661d8f46659a00d28551681210ce7777",
		"coredns":                   "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:260f68e2b587372f6b36206c82d1a1b2d52edc09f5831643d64734586919253b",
		"haproxy_router":            "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:0f4a27f848795eefe6c2e8d53224eecc2831fcb377f1602ce6cff2ef40e8fca7",
		"kube_rbac_proxy":           "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:cb3c0d2b5a6420a432d0df9fc5cabc44e5529acffb489d13bd6c7570fd4e3adc",
		"openssl":                   "registry.access.redhat.com/ubi8/openssl@sha256:3f781a07e59d164eba065dba7d8e7661ab2494b21199c379b65b0ff514a1b8d0",
		"ovn_kubernetes_microshift": "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:f87a24869bcb305dd726cc1f45115415894a9b0ebdc0c5902a038abe051765d6",
		"pod":                       "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:a8ef2351440927bfc107456a0eb1a2c0a2bc5de49765be7dff0fdfbde60d030f",
		"service_ca_operator":       "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:10be1a369305d1453fecf100a3489bdf2000d1f3c09c020999e04b2ec7e43698",
		"odf_topolvm":               "quay.io/rhceph-dev/odf4-odf-topolvm-rhel8@sha256:2855918d1849c99a835eb03c53ce07170c238111fd15d2fe50cd45611fcd1ceb",
                "ose_csi_ext_provisioner":   "quay.io/rhceph-dev/openshift-ose-csi-external-provisioner@sha256:c3b2417f8fcb8883275f0e613037f83133ccc3f91311a30688e4be520544ea4a",
                "ose_csi_ext_resizer":       "quay.io/rhceph-dev/openshift-ose-csi-external-resizer@sha256:213f43d61b3a214a4a433c7132537be082a108d55005f2ba0777c2ea97489799",
                "topolvm-csi-snapshotter":   "quay.io/rhceph-dev/openshift-ose-csi-external-snapshotter@sha256:734c095670d21b77f18c84670d6c9a7742be1d9151dca0da20f41858ede65ed8",
                "ose_csi_livenessprobe":     "quay.io/rhceph-dev/openshift-ose-csi-livenessprobe@sha256:b05559aa038708ab448cfdfed2ca880726aed6cc30371fea4d6a42c972c0c728",
                "ose_csi_node_registrar":    "quay.io/rhceph-dev/openshift-ose-csi-node-driver-registrar@sha256:fb0f5e531847db94dcadc61446b9a892f6f92ddf282e192abf2fdef6c6af78f2",
	}
}
