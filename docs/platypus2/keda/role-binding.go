// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package keda

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OperatorAuthReaderRB = &rbacv1.RoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component":  "operator",
			"app.kubernetes.io/instance":   "keda",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "keda-operator-auth-reader",
			"app.kubernetes.io/part-of":    "keda-operator",
			"app.kubernetes.io/version":    "2.11.1",
			"helm.sh/chart":                "keda-2.11.1",
		},
		Name:      "keda-operator-auth-reader",
		Namespace: "kube-system",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Role",
		Name:     "extension-apiserver-authentication-reader",
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      "keda-operator",
			Namespace: "keda",
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "RoleBinding",
	},
}

var OperatorRB = &rbacv1.RoleBinding{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component":  "operator",
			"app.kubernetes.io/instance":   "keda",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "keda-operator",
			"app.kubernetes.io/part-of":    "keda-operator",
			"app.kubernetes.io/version":    "2.11.1",
			"helm.sh/chart":                "keda-2.11.1",
		},
		Name:      "keda-operator",
		Namespace: "keda",
	},
	RoleRef: rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Role",
		Name:     "keda-operator",
	},
	Subjects: []rbacv1.Subject{
		{
			Kind:      "ServiceAccount",
			Name:      "keda-operator",
			Namespace: "keda",
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "RoleBinding",
	},
}
