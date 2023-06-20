// Copyright (c) 2023 Volvo Car Corporation
// SPDX-License-Identifier: Apache-2.0

// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package metricsserver

import (
	ku "github.com/volvo-cars/lingon/pkg/kubeutil"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SA = &corev1.ServiceAccount{
	ObjectMeta: metav1.ObjectMeta{
		Labels:    BaseLabels(),
		Name:      appName,
		Namespace: namespace,
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ServiceAccount",
	},
}

var AuthReaderRB = &rbacv1.RoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels:    BaseLabels(),
		Name:      "metrics-server-auth-reader",
		Namespace: ku.NSKubeSystem,
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Role",
		Name:     "extension-apiserver-authentication-reader", // predefined ?
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      SA.Name,
			Namespace: namespace,
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "RoleBinding",
	},
}

var SystemAggregatedReaderCR = &rbacv1.ClusterRole{
	ObjectMeta: metav1.ObjectMeta{
		Labels: ku.MergeLabels(
			BaseLabels(), map[string]string{
				ku.LabelRbacAggregateToAdmin: "true",
				ku.LabelRbacAggregateToEdit:  "true",
				ku.LabelRbacAggregateToView:  "true",
			},
		),
		Name: "system:metrics-server-aggregated-reader",
	},
	Rules: []rbacv1.PolicyRule{
		{
			APIGroups: []string{"metrics.k8s.io"},
			Resources: []string{"pods", "nodes"},
			Verbs:     []string{"get", "list", "watch"},
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRole",
	},
}

var SystemCR = &rbacv1.ClusterRole{
	ObjectMeta: metav1.ObjectMeta{
		Labels: BaseLabels(),
		Name:   "system:" + appName,
	},
	Rules: []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"nodes/metrics"},
			Verbs:     []string{"get"},
		}, {
			APIGroups: []string{""},
			Resources: []string{"pods", "nodes", "namespaces", "configmaps"},
			Verbs:     []string{"get", "list", "watch"},
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRole",
	},
}

var SystemAuthDelegatorCRB = &rbacv1.ClusterRoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: BaseLabels(),
		Name:   "metrics-server:system:auth-delegator",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "ClusterRole",
		Name:     "system:auth-delegator",
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      SA.Name,
			Namespace: namespace,
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRoleBinding",
	},
}

var SystemCRB = &rbacv1.ClusterRoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: BaseLabels(),
		Name:   "system:" + appName,
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "ClusterRole",
		Name:     SystemCR.Name,
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      SA.Name,
			Namespace: namespace,
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "ClusterRoleBinding",
	},
}