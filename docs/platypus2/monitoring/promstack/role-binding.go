// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package promstack

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var KubePromtheusStackGrafanaRB = &rbacv1.RoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "grafana",
			"app.kubernetes.io/version":    "9.5.3",
			"helm.sh/chart":                "grafana-6.57.1",
		},
		Name:      "kube-promtheus-stack-grafana",
		Namespace: "monitoring",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Role",
		Name:     "kube-promtheus-stack-grafana",
	},
	Subjects: []rbacv1.Subject{{
		Kind:      "ServiceAccount",
		Name:      "kube-promtheus-stack-grafana",
		Namespace: "monitoring",
	}},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "RoleBinding",
	},
}

var KubePromtheusStackKubeAdmissionRB = &rbacv1.RoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Annotations: map[string]string{
			"helm.sh/hook":               "pre-install,pre-upgrade,post-install,post-upgrade",
			"helm.sh/hook-delete-policy": "before-hook-creation,hook-succeeded",
		},
		Labels: map[string]string{
			"app":                          "kube-prometheus-stack-admission",
			"app.kubernetes.io/instance":   "kube-promtheus-stack",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/part-of":    "kube-prometheus-stack",
			"app.kubernetes.io/version":    "46.8.0",
			"chart":                        "kube-prometheus-stack-46.8.0",
			"heritage":                     "Helm",
			"release":                      "kube-promtheus-stack",
		},
		Name:      "kube-promtheus-stack-kube-admission",
		Namespace: "monitoring",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Role",
		Name:     "kube-promtheus-stack-kube-admission",
	},
	Subjects: []rbacv1.Subject{{
		Kind:      "ServiceAccount",
		Name:      "kube-promtheus-stack-kube-admission",
		Namespace: "monitoring",
	}},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "RoleBinding",
	},
}