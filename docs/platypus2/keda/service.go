// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package keda

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var OperatorMetricsApiserverSVC = &corev1.Service{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app":                          "keda-operator-metrics-apiserver",
			"app.kubernetes.io/component":  "operator",
			"app.kubernetes.io/instance":   "keda",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "keda-operator-metrics-apiserver",
			"app.kubernetes.io/part-of":    "keda-operator",
			"app.kubernetes.io/version":    "2.11.1",
			"helm.sh/chart":                "keda-2.11.1",
		},
		Name:      "keda-operator-metrics-apiserver",
		Namespace: "keda",
	},
	Spec: corev1.ServiceSpec{
		Ports: []corev1.ServicePort{
			{
				Name:       "https",
				Port:       int32(443),
				Protocol:   corev1.Protocol("TCP"),
				TargetPort: intstr.IntOrString{IntVal: int32(6443)},
			}, {
				Name:       "metrics",
				Port:       int32(8080),
				Protocol:   corev1.Protocol("TCP"),
				TargetPort: intstr.IntOrString{IntVal: int32(8080)},
			},
		},
		Selector: map[string]string{"app": "keda-operator-metrics-apiserver"},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "Service",
	},
}

var OperatorSVC = &corev1.Service{
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
	Spec: corev1.ServiceSpec{
		Ports: []corev1.ServicePort{
			{
				Name:       "metricsservice",
				Port:       int32(9666),
				TargetPort: intstr.IntOrString{IntVal: int32(9666)},
			},
		},
		Selector: map[string]string{"app": "keda-operator"},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "Service",
	},
}
