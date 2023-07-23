// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package dragonflydb

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Deploy = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance":   "dragonflydb",
			"app.kubernetes.io/managed-by": "Helm",
			"app.kubernetes.io/name":       "dragonfly",
			"app.kubernetes.io/version":    "v1.6.1",
			"helm.sh/chart":                "dragonfly-v1.6.1",
		},
		Name:      "dragonflydb",
		Namespace: "dragonflydb",
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: P(int32(1)),
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app.kubernetes.io/instance": "dragonflydb",
				"app.kubernetes.io/name":     "dragonfly",
			},
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: map[string]string{"checksum/tls-secret": "01ba4719c80b6fe911b091a7c05124b64eeece964e09c058ef8f9805daca546b"},
				Labels: map[string]string{
					"app.kubernetes.io/instance": "dragonflydb",
					"app.kubernetes.io/name":     "dragonfly",
				},
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Args: []string{
							"--alsologtostderr",
							"--tls",
							"--tls_cert_file=/etc/dragonfly/tls/tls.crt",
							"--tls_key_file=/etc/dragonfly/tls/tls.key",
						},
						Image:           "docker.dragonflydb.io/dragonflydb/dragonfly:v1.6.1",
						ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
						LivenessProbe: &corev1.Probe{
							FailureThreshold:    int32(3),
							InitialDelaySeconds: int32(10),
							PeriodSeconds:       int32(10),
							ProbeHandler: corev1.ProbeHandler{
								Exec: &corev1.ExecAction{
									Command: []string{
										"/bin/sh",
										"/usr/local/bin/healthcheck.sh",
									},
								},
							},
							SuccessThreshold: int32(1),
							TimeoutSeconds:   int32(5),
						},
						Name: "dragonfly",
						Ports: []corev1.ContainerPort{
							{
								ContainerPort: int32(6379),
								Name:          "dragonfly",
								Protocol:      corev1.Protocol("TCP"),
							},
						},
						ReadinessProbe: &corev1.Probe{
							FailureThreshold:    int32(3),
							InitialDelaySeconds: int32(10),
							PeriodSeconds:       int32(10),
							ProbeHandler: corev1.ProbeHandler{
								Exec: &corev1.ExecAction{
									Command: []string{
										"/bin/sh",
										"/usr/local/bin/healthcheck.sh",
									},
								},
							},
							SuccessThreshold: int32(1),
							TimeoutSeconds:   int32(5),
						},
						Resources: corev1.ResourceRequirements{
							Limits:   map[corev1.ResourceName]resource.Quantity{},
							Requests: map[corev1.ResourceName]resource.Quantity{},
						},
						SecurityContext: &corev1.SecurityContext{
							Capabilities:           &corev1.Capabilities{Drop: []corev1.Capability{corev1.Capability("ALL")}},
							ReadOnlyRootFilesystem: P(true),
							RunAsNonRoot:           P(true),
							RunAsUser:              P(int64(1000)),
						},
						VolumeMounts: []corev1.VolumeMount{
							{
								MountPath: "/etc/dragonfly/tls",
								Name:      "tls",
							},
						},
					},
				},
				SecurityContext:    &corev1.PodSecurityContext{FSGroup: P(int64(2000))},
				ServiceAccountName: "dragonflydb",
				Volumes: []corev1.Volume{
					{
						Name:         "tls",
						VolumeSource: corev1.VolumeSource{Secret: &corev1.SecretVolumeSource{SecretName: "dragonflydb-server-tls"}},
					},
				},
			},
		},
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
	},
}
