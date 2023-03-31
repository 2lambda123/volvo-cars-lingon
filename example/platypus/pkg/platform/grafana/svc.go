// Code generated by go-kart. EDIT AS MUCH AS YOU LIKE.

package grafana

import (
	"github.com/volvo-cars/lingon/pkg/kubeutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func Service(opts KubeOpts) *corev1.Service {
	return &corev1.Service{
		TypeMeta: kubeutil.TypeServiceV1,
		ObjectMeta: metav1.ObjectMeta{
			Name:      opts.Name,
			Namespace: opts.Namespace,
			Labels:    opts.CommonLabels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       "service",
					Port:       int32(80),
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.IntOrString{IntVal: int32(3000)},
				},
			},
			Selector: opts.CommonLabels,
			Type:     corev1.ServiceTypeClusterIP,
		},
	}
}
