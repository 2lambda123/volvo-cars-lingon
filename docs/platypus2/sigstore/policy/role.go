// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package policy

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SigstoreCleanupRole = &rbacv1.Role{
	ObjectMeta: metav1.ObjectMeta{
		Annotations: map[string]string{
			"helm.sh/hook":               "post-delete",
			"helm.sh/hook-delete-policy": "hook-succeeded",
			"helm.sh/hook-weight":        "1",
		},
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "sigstore",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "policy-controller",
			"app.kubernetes.io/version":    "0.8.0",
			"control-plane":                "sigstore-policy-controller-cleanup",
			"helm.sh/chart":                "policy-controller-0.6.0",
		},
		Name:      "sigstore-policy-controller-cleanup",
		Namespace: "sigstore",
	},
	Rules: []rbacv1.PolicyRule{
		{
			APIGroups: []string{"coordination.k8s.io"},
			Resources: []string{"leases"},
			Verbs:     []string{"list", "delete"},
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "Role",
	},
}

var SigstoreWebhookRole = &rbacv1.Role{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "sigstore",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "policy-controller",
			"app.kubernetes.io/version":    "0.8.0",
			"control-plane":                "sigstore-policy-controller-webhook",
			"helm.sh/chart":                "policy-controller-0.6.0",
		},
		Name:      "sigstore-policy-controller-webhook",
		Namespace: "sigstore",
	},
	Rules: []rbacv1.PolicyRule{
		{
			APIGroups: []string{""},
			Resources: []string{"configmaps", "secrets"},
			Verbs:     []string{"get", "list", "update", "watch"},
		}, {
			APIGroups: []string{"coordination.k8s.io"},
			Resources: []string{"leases"},
			Verbs: []string{
				"get",
				"list",
				"create",
				"update",
				"delete",
				"patch",
				"watch",
			},
		}, {
			APIGroups:     []string{""},
			ResourceNames: []string{"config-image-policies"},
			Resources:     []string{"configmaps"},
			Verbs: []string{
				"get",
				"list",
				"create",
				"update",
				"patch",
				"watch",
			},
		}, {
			APIGroups:     []string{""},
			ResourceNames: []string{"config-sigstore-keys"},
			Resources:     []string{"configmaps"},
			Verbs: []string{
				"get",
				"list",
				"create",
				"update",
				"patch",
				"watch",
			},
		}, {
			APIGroups: []string{"policy.sigstore.dev"},
			Resources: []string{"trustroots"},
			Verbs:     []string{"get", "list"},
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "rbac.authorization.k8s.io/v1",
		Kind:       "Role",
	},
}