package cilium

import (
	"fmt"

	"github.com/volvo-cars/lingon/pkg/kube"
	"github.com/volvo-cars/lingon/pkg/kubeutil"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

var _ kube.Exporter = (*CiliumApp)(nil)

// todo:
// - generate certificate
// - Hubble + Hubble UI + hubble certificates

type CiliumApp struct {
	kube.App

	CACertSecret *corev1.Secret
	Config       *corev1.ConfigMap

	MattDamon  *appsv1.DaemonSet
	NodeInit   *appsv1.DaemonSet
	AppSA      *corev1.ServiceAccount
	AppBinding *rbacv1.ClusterRoleBinding
	AppRole    *rbacv1.ClusterRole

	Operator  *appsv1.Deployment
	OpSA      *corev1.ServiceAccount
	OpBinding *rbacv1.ClusterRoleBinding
	OpRole    *rbacv1.ClusterRole

	HubbleSecret *corev1.Secret
	HubbleCert   *corev1.Secret
	HubbleSvc    *corev1.Service
}

type ClusterConfig struct {
	ClusterName string
	ClusterID   int
}

func New(config ClusterConfig) *CiliumApp {
	ciliumCM := map[string]string{
		"cluster-name": config.ClusterName,
		"cluster-id":   fmt.Sprintf("%d", config.ClusterID),
	}

	ns := "kube-system"

	// HACK: this is freckin' horrible!!!!!!!!
	// Must remove and generate once and store somewhere, e.g. Vault
	caCerts := map[string][]byte{
		"ca.crt": []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGRENDQWZ5Z0F3SUJBZ0lSQUlKUnROOEV0SFY3b3Vxci8ySnZoYUV3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl5TVRJeE5qQTVNVGd6TmxvWERUSTFNVEl4TlRBNQpNVGd6Tmxvd0ZERVNNQkFHQTFVRUF4TUpRMmxzYVhWdElFTkJNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DCkFROEFNSUlCQ2dLQ0FRRUExR2Z1M00zQzVlK3dFYnNWSkxYQ2ExdVBiWjJYWXBDQUZPWHQ2UmZEYzhzM2p2MjUKRFFlU2Z6VFdQNE1oZkkvQk1BbUZreHRTcjhmZUtHS2RIcndLN1B5bEhhRTQ3Y09aZ3VPMzdiR0tZWXhEVml4dQorSW5yTnJXRWhlMlR0MWNhM3h3MFp2d3orNkZWeFNJZEFFU1BwcVNxdlY2OEpkS2JZQkRmV2oveERoSUFwdGdGCmVSM2F5U3V3aGZ6Slk0WE4xQkFxeWhoWmFaZUdjRC9seUdmcmRhZmtDdmVQRk5jRVV6Q1pHdHA0ZXFsOXVTdXEKNjJhdXBDcUJuTzBnZkJ0Q0NVUW5tZDRnMXlYMTErRUFmQXFVVmE2QWUxYnlIVlorRTI0dFgveXhwNU9yNGtRWApJUUVuUlUvYTJ5RDJ0NWNYZTZrWGQxWi8xWXJzcE9XV0l4cXRsUUlEQVFBQm8yRXdYekFPQmdOVkhROEJBZjhFCkJBTUNBcVF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01BOEdBMVVkRXdFQi93UUYKTUFNQkFmOHdIUVlEVlIwT0JCWUVGRHgyZTkrblRKNUVFYWFiakxJRUV3N2w1ZWw4TUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQmRMRXRTZkFCWlRSTUhoeVk0Nmd2QnNBRnVSTzY5V2xOSmNCUnlyNE9FSDVZRGpYOGU4L2hWCkFmY0NFMU5WaGxoN2JOMTg4TVNkaTY5eHMzOFpmMVc4c0d1UHpxaisrQmZPNlRKTWhxTUVZbkk5VWtodHFQNXYKNVRpR3JscGY3MCtVQXZyYkpXVXI5VHV1UHN3NTk5T25Tb3NST2hwbkxQNDIzL1lSaWhKbDFMdGNsdmhBMGNiVgphSEtHZHNjekxuNGpHeW8ydnRRcnJFTFVJb1BWOW5Mb2x0cGNGOFcyeVJhNEp5b2JZY294WHorWmxrWTNRanlkClB3dnp5blc5cEZQWWtHL0I1YS9SNGtPNHpzR0tZbWZnc0Z1cG94WW5iMXpyVzkxTS9TSm5OZTRyNXl5Qkw0NS8KN1hzbWdnUWFQV1A2ckZ3N3FhVHdhZk5rU0twcHVMY3gKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="),
		"ca.key": []byte("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMUdmdTNNM0M1ZSt3RWJzVkpMWENhMXVQYloyWFlwQ0FGT1h0NlJmRGM4czNqdjI1CkRRZVNmelRXUDRNaGZJL0JNQW1Ga3h0U3I4ZmVLR0tkSHJ3SzdQeWxIYUU0N2NPWmd1TzM3YkdLWVl4RFZpeHUKK0luck5yV0VoZTJUdDFjYTN4dzBadnd6KzZGVnhTSWRBRVNQcHFTcXZWNjhKZEtiWUJEZldqL3hEaElBcHRnRgplUjNheVN1d2hmekpZNFhOMUJBcXloaFphWmVHY0QvbHlHZnJkYWZrQ3ZlUEZOY0VVekNaR3RwNGVxbDl1U3VxCjYyYXVwQ3FCbk8wZ2ZCdENDVVFubWQ0ZzF5WDExK0VBZkFxVVZhNkFlMWJ5SFZaK0UyNHRYL3l4cDVPcjRrUVgKSVFFblJVL2EyeUQydDVjWGU2a1hkMVovMVlyc3BPV1dJeHF0bFFJREFRQUJBb0lCQUUzdFE2a21wRmFQdFYwTAo4aG5oeFU1MTdRMGVRQ2dkTTZCM0t1M1ZsaE9wZnR5cklYVXlUZ0QxZFpVZm11MkVJREJyamVJR3FETnRkSWdFCmhmaDhyTlY5YTJhUGU3OWZmN2FSclMwN2NiV1FMRFExWVJFMktHR04vdXpUMk5udXp5RUR6QVhzaVhYTUh4ZVEKQ0d2TXU1YzcycGhYWlZmTENNNFo0cGZOMWJaL0lNZWphTEF1b3RDSW1LWFQ3SW9CQjJwenQ2SlBVT0N2ekI1eQpZazZCS3ZQdzA3c3RWa09hbGFIbzVjOVErMUVYMXhsQmVYK1dYeElacURaVkl3UVI5MWtva3drN2RvUitXYm50ClU2dlNGN01kR3orVUduR0oyYjZtRHNBSldJNW9KWDFDZW5JbGVWajV3ME83eGQyQnBRQmcvcE9sa1l5dDZhWUEKKy9kUnBZRUNnWUVBK3NXZkZHckRWWGxkRWQ4SVpQa1Bud1NHNHQxMG5QaVlMOWxZTzJlSDVpRmx5ZmFKSmQwTQplQ2M1LzBhcFFrNG9pWDJZU2thZDZIaTFBSnhLYkIyUGtnWVRyR0wvWm81aEdwU25vcGtEUlBkcCtiWnRVbS9LCitzdGtNaVJRQXcxc0I1MHZiRTVGRFcyRU1hbDk2RW8vMjk3a2hqNjVOTXdpM0p3cURZY0R2UDBDZ1lFQTJOV04KRjl2TzRpbm8xM3RzUytncW5aYlJScXJ2dysvY2s4N1c5bm1LR3FkNmlRWGxyN0VvM2p3VDBkU2hpRU0yWVI2OApzNk0xdXVEMnZGUVlSdjB0Zk45MVdGQ204SFVZTGNEMHA0S2h2Y21mKzBqR0ZWVXJwMDVZSGRYWUhiZnFzdEgyCjBOa2Z1SlpOeERqQ2JWZnNoWkUwZ0U2QTM3Q0hvM29JWXBNcjRua0NnWUVBaFMydzExSC8zUVB3Tm82QlVjYW4KMGliQVQvbVdkY3JjWUFVSWc3dnZBM3ZYS0JRak1CV2VDcTJpY24wZlpOUkhXUVYzZkhMV1orQzdGOURwQVZRTgpyRnBIMW5SWStTbENUckNGK3FkU2dpejNmaU94R1dlL244T211YTVwUThXOENxc2l4VjBuOVFLbGd3NWxqSmpxCkl5WFRyYXZnQmpjbmlJdnpGVzRQd1MwQ2dZQS81Z2F6UGpwa1diZGRNT2tFSVlIVmVPbHBLdHlIWURzZmI3ZlgKWUkrN05SbWVJWmZEUTdEb2RNbmVid3UvTFJkc1hYTjhlSjlQMkJXK0FBWTdmVWFYWXY2Z3JQdlZKcHllZHh0aAo5ZFFXS3NHemVvbXRKYkU4bDVET0VnT2pGbHphbjVkZGltNlhwZXQ2NU1NZkYrY0NvWHpZUnNvaG9WTUhjT0hoCnNyOGUrUUtCZ1FEd2ZnSFkwZmRZYkYwWExVcndGUTVONUJlNTFmM05tTVorSjhNS3NsajF2ZUF2dHc4V1VsOGcKY3BRck9UY1FXSTQwQ3FkU3lFL1BSTlJhcUo3TXpWM3JVWlpyS2hjbUU5S0V0emdIQUl2VWl2MUZlb0ZBZ1Q2SQpZOEtXS2lhck9oVi9TQXhidFg3QVFTT21nMzhRbDYxeWVPNm8zakMzQVlrOVhkMHl4dUpkamc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="),
	}
	hubbleCACerts := map[string][]byte{
		"ca.crt": []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGRENDQWZ5Z0F3SUJBZ0lSQUlKUnROOEV0SFY3b3Vxci8ySnZoYUV3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl5TVRJeE5qQTVNVGd6TmxvWERUSTFNVEl4TlRBNQpNVGd6Tmxvd0ZERVNNQkFHQTFVRUF4TUpRMmxzYVhWdElFTkJNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DCkFROEFNSUlCQ2dLQ0FRRUExR2Z1M00zQzVlK3dFYnNWSkxYQ2ExdVBiWjJYWXBDQUZPWHQ2UmZEYzhzM2p2MjUKRFFlU2Z6VFdQNE1oZkkvQk1BbUZreHRTcjhmZUtHS2RIcndLN1B5bEhhRTQ3Y09aZ3VPMzdiR0tZWXhEVml4dQorSW5yTnJXRWhlMlR0MWNhM3h3MFp2d3orNkZWeFNJZEFFU1BwcVNxdlY2OEpkS2JZQkRmV2oveERoSUFwdGdGCmVSM2F5U3V3aGZ6Slk0WE4xQkFxeWhoWmFaZUdjRC9seUdmcmRhZmtDdmVQRk5jRVV6Q1pHdHA0ZXFsOXVTdXEKNjJhdXBDcUJuTzBnZkJ0Q0NVUW5tZDRnMXlYMTErRUFmQXFVVmE2QWUxYnlIVlorRTI0dFgveXhwNU9yNGtRWApJUUVuUlUvYTJ5RDJ0NWNYZTZrWGQxWi8xWXJzcE9XV0l4cXRsUUlEQVFBQm8yRXdYekFPQmdOVkhROEJBZjhFCkJBTUNBcVF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01BOEdBMVVkRXdFQi93UUYKTUFNQkFmOHdIUVlEVlIwT0JCWUVGRHgyZTkrblRKNUVFYWFiakxJRUV3N2w1ZWw4TUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQmRMRXRTZkFCWlRSTUhoeVk0Nmd2QnNBRnVSTzY5V2xOSmNCUnlyNE9FSDVZRGpYOGU4L2hWCkFmY0NFMU5WaGxoN2JOMTg4TVNkaTY5eHMzOFpmMVc4c0d1UHpxaisrQmZPNlRKTWhxTUVZbkk5VWtodHFQNXYKNVRpR3JscGY3MCtVQXZyYkpXVXI5VHV1UHN3NTk5T25Tb3NST2hwbkxQNDIzL1lSaWhKbDFMdGNsdmhBMGNiVgphSEtHZHNjekxuNGpHeW8ydnRRcnJFTFVJb1BWOW5Mb2x0cGNGOFcyeVJhNEp5b2JZY294WHorWmxrWTNRanlkClB3dnp5blc5cEZQWWtHL0I1YS9SNGtPNHpzR0tZbWZnc0Z1cG94WW5iMXpyVzkxTS9TSm5OZTRyNXl5Qkw0NS8KN1hzbWdnUWFQV1A2ckZ3N3FhVHdhZk5rU0twcHVMY3gKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="),
		"ca.key": []byte("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMUdmdTNNM0M1ZSt3RWJzVkpMWENhMXVQYloyWFlwQ0FGT1h0NlJmRGM4czNqdjI1CkRRZVNmelRXUDRNaGZJL0JNQW1Ga3h0U3I4ZmVLR0tkSHJ3SzdQeWxIYUU0N2NPWmd1TzM3YkdLWVl4RFZpeHUKK0luck5yV0VoZTJUdDFjYTN4dzBadnd6KzZGVnhTSWRBRVNQcHFTcXZWNjhKZEtiWUJEZldqL3hEaElBcHRnRgplUjNheVN1d2hmekpZNFhOMUJBcXloaFphWmVHY0QvbHlHZnJkYWZrQ3ZlUEZOY0VVekNaR3RwNGVxbDl1U3VxCjYyYXVwQ3FCbk8wZ2ZCdENDVVFubWQ0ZzF5WDExK0VBZkFxVVZhNkFlMWJ5SFZaK0UyNHRYL3l4cDVPcjRrUVgKSVFFblJVL2EyeUQydDVjWGU2a1hkMVovMVlyc3BPV1dJeHF0bFFJREFRQUJBb0lCQUUzdFE2a21wRmFQdFYwTAo4aG5oeFU1MTdRMGVRQ2dkTTZCM0t1M1ZsaE9wZnR5cklYVXlUZ0QxZFpVZm11MkVJREJyamVJR3FETnRkSWdFCmhmaDhyTlY5YTJhUGU3OWZmN2FSclMwN2NiV1FMRFExWVJFMktHR04vdXpUMk5udXp5RUR6QVhzaVhYTUh4ZVEKQ0d2TXU1YzcycGhYWlZmTENNNFo0cGZOMWJaL0lNZWphTEF1b3RDSW1LWFQ3SW9CQjJwenQ2SlBVT0N2ekI1eQpZazZCS3ZQdzA3c3RWa09hbGFIbzVjOVErMUVYMXhsQmVYK1dYeElacURaVkl3UVI5MWtva3drN2RvUitXYm50ClU2dlNGN01kR3orVUduR0oyYjZtRHNBSldJNW9KWDFDZW5JbGVWajV3ME83eGQyQnBRQmcvcE9sa1l5dDZhWUEKKy9kUnBZRUNnWUVBK3NXZkZHckRWWGxkRWQ4SVpQa1Bud1NHNHQxMG5QaVlMOWxZTzJlSDVpRmx5ZmFKSmQwTQplQ2M1LzBhcFFrNG9pWDJZU2thZDZIaTFBSnhLYkIyUGtnWVRyR0wvWm81aEdwU25vcGtEUlBkcCtiWnRVbS9LCitzdGtNaVJRQXcxc0I1MHZiRTVGRFcyRU1hbDk2RW8vMjk3a2hqNjVOTXdpM0p3cURZY0R2UDBDZ1lFQTJOV04KRjl2TzRpbm8xM3RzUytncW5aYlJScXJ2dysvY2s4N1c5bm1LR3FkNmlRWGxyN0VvM2p3VDBkU2hpRU0yWVI2OApzNk0xdXVEMnZGUVlSdjB0Zk45MVdGQ204SFVZTGNEMHA0S2h2Y21mKzBqR0ZWVXJwMDVZSGRYWUhiZnFzdEgyCjBOa2Z1SlpOeERqQ2JWZnNoWkUwZ0U2QTM3Q0hvM29JWXBNcjRua0NnWUVBaFMydzExSC8zUVB3Tm82QlVjYW4KMGliQVQvbVdkY3JjWUFVSWc3dnZBM3ZYS0JRak1CV2VDcTJpY24wZlpOUkhXUVYzZkhMV1orQzdGOURwQVZRTgpyRnBIMW5SWStTbENUckNGK3FkU2dpejNmaU94R1dlL244T211YTVwUThXOENxc2l4VjBuOVFLbGd3NWxqSmpxCkl5WFRyYXZnQmpjbmlJdnpGVzRQd1MwQ2dZQS81Z2F6UGpwa1diZGRNT2tFSVlIVmVPbHBLdHlIWURzZmI3ZlgKWUkrN05SbWVJWmZEUTdEb2RNbmVid3UvTFJkc1hYTjhlSjlQMkJXK0FBWTdmVWFYWXY2Z3JQdlZKcHllZHh0aAo5ZFFXS3NHemVvbXRKYkU4bDVET0VnT2pGbHphbjVkZGltNlhwZXQ2NU1NZkYrY0NvWHpZUnNvaG9WTUhjT0hoCnNyOGUrUUtCZ1FEd2ZnSFkwZmRZYkYwWExVcndGUTVONUJlNTFmM05tTVorSjhNS3NsajF2ZUF2dHc4V1VsOGcKY3BRck9UY1FXSTQwQ3FkU3lFL1BSTlJhcUo3TXpWM3JVWlpyS2hjbUU5S0V0emdIQUl2VWl2MUZlb0ZBZ1Q2SQpZOEtXS2lhck9oVi9TQXhidFg3QVFTT21nMzhRbDYxeWVPNm8zakMzQVlrOVhkMHl4dUpkamc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="),
	}
	hubbleServerCerts := map[string][]byte{
		"ca.crt":  []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGRENDQWZ5Z0F3SUJBZ0lSQUlKUnROOEV0SFY3b3Vxci8ySnZoYUV3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl5TVRJeE5qQTVNVGd6TmxvWERUSTFNVEl4TlRBNQpNVGd6Tmxvd0ZERVNNQkFHQTFVRUF4TUpRMmxzYVhWdElFTkJNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DCkFROEFNSUlCQ2dLQ0FRRUExR2Z1M00zQzVlK3dFYnNWSkxYQ2ExdVBiWjJYWXBDQUZPWHQ2UmZEYzhzM2p2MjUKRFFlU2Z6VFdQNE1oZkkvQk1BbUZreHRTcjhmZUtHS2RIcndLN1B5bEhhRTQ3Y09aZ3VPMzdiR0tZWXhEVml4dQorSW5yTnJXRWhlMlR0MWNhM3h3MFp2d3orNkZWeFNJZEFFU1BwcVNxdlY2OEpkS2JZQkRmV2oveERoSUFwdGdGCmVSM2F5U3V3aGZ6Slk0WE4xQkFxeWhoWmFaZUdjRC9seUdmcmRhZmtDdmVQRk5jRVV6Q1pHdHA0ZXFsOXVTdXEKNjJhdXBDcUJuTzBnZkJ0Q0NVUW5tZDRnMXlYMTErRUFmQXFVVmE2QWUxYnlIVlorRTI0dFgveXhwNU9yNGtRWApJUUVuUlUvYTJ5RDJ0NWNYZTZrWGQxWi8xWXJzcE9XV0l4cXRsUUlEQVFBQm8yRXdYekFPQmdOVkhROEJBZjhFCkJBTUNBcVF3SFFZRFZSMGxCQll3RkFZSUt3WUJCUVVIQXdFR0NDc0dBUVVGQndNQ01BOEdBMVVkRXdFQi93UUYKTUFNQkFmOHdIUVlEVlIwT0JCWUVGRHgyZTkrblRKNUVFYWFiakxJRUV3N2w1ZWw4TUEwR0NTcUdTSWIzRFFFQgpDd1VBQTRJQkFRQmRMRXRTZkFCWlRSTUhoeVk0Nmd2QnNBRnVSTzY5V2xOSmNCUnlyNE9FSDVZRGpYOGU4L2hWCkFmY0NFMU5WaGxoN2JOMTg4TVNkaTY5eHMzOFpmMVc4c0d1UHpxaisrQmZPNlRKTWhxTUVZbkk5VWtodHFQNXYKNVRpR3JscGY3MCtVQXZyYkpXVXI5VHV1UHN3NTk5T25Tb3NST2hwbkxQNDIzL1lSaWhKbDFMdGNsdmhBMGNiVgphSEtHZHNjekxuNGpHeW8ydnRRcnJFTFVJb1BWOW5Mb2x0cGNGOFcyeVJhNEp5b2JZY294WHorWmxrWTNRanlkClB3dnp5blc5cEZQWWtHL0I1YS9SNGtPNHpzR0tZbWZnc0Z1cG94WW5iMXpyVzkxTS9TSm5OZTRyNXl5Qkw0NS8KN1hzbWdnUWFQV1A2ckZ3N3FhVHdhZk5rU0twcHVMY3gKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="),
		"tls.crt": []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURWekNDQWorZ0F3SUJBZ0lSQUkwZ3NqSjgzSHZaQUpsOHNsZUdXVEV3RFFZSktvWklodmNOQVFFTEJRQXcKRkRFU01CQUdBMVVFQXhNSlEybHNhWFZ0SUVOQk1CNFhEVEl5TVRJeE5qQTVNVGd6TjFvWERUSTFNVEl4TlRBNQpNVGd6TjFvd0tqRW9NQ1lHQTFVRUF3d2ZLaTVrWldaaGRXeDBMbWgxWW1Kc1pTMW5jbkJqTG1OcGJHbDFiUzVwCmJ6Q0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQUwrZ3BqcWlSYzlYSE1ScUt3NWsKR1dVZHpsN2hvaThpcmU4M0JpeHhJb3NjQWN4R0tZTFhBUytFRTB3R05ZbmFZMHhmTHNNbU8vM29OZ3FyTzM4aQp4OHB3eVRPeHZ2Sk83dXhaYU41U1U1ZnEyaVJwSHZ3UmtsZkJVblhsYmhrdFo2Vkg4bkZrRVZSdWpzVnd4alh3CktGV0xHWTA4YTl6K1dKbUs5eHN1NUt2bnZKQmV3WWFvUjdUcGN2UTBWYk40L0hXSVFsTWt1VyszcjU2azRYYk0KNHBESWdhUmplbzh2dkZzS1VVbWVlVU5tczlsUTdEWm53K2Y0bW10NUZmWjRGdDhiTlNiTmhER2J0OU5OT3puWgpKME1keWRrUmFJUHM2NVZRUWI5eS9WVks4SjBrODZUY3FKZnpuSzRPMm5Ua3Z1bHBNVXJLdktwUldTeEhXYUxKCnN1c0NBd0VBQWFPQmpUQ0JpakFPQmdOVkhROEJBZjhFQkFNQ0JhQXdIUVlEVlIwbEJCWXdGQVlJS3dZQkJRVUgKQXdFR0NDc0dBUVVGQndNQ01Bd0dBMVVkRXdFQi93UUNNQUF3SHdZRFZSMGpCQmd3Rm9BVVBIWjczNmRNbmtRUgpwcHVNc2dRVER1WGw2WHd3S2dZRFZSMFJCQ013SVlJZktpNWtaV1poZFd4MExtaDFZbUpzWlMxbmNuQmpMbU5wCmJHbDFiUzVwYnpBTkJna3Foa2lHOXcwQkFRc0ZBQU9DQVFFQVRlQmFwL0x5Rk5xV2lUL1NhbDRwMUt2UmM3cEUKaU5tWmtNZ3o0N0d3TERCaXNpamRZNEpWV1BERyt3ME1zb3Q4RDArUkJSUThMRVU1K0dsUXFHNWhUdk50QktVYwpNYWJ0d083R1VWNHF6ZElEMUdUMzRibEtUWHNlZVpCZ1ppY2VUYWRlR2JoZVZ1QS9JbEt2WTFZeEtucWhZbmlmCmExSlBLK1h5QUdDcHBZRjludnBUQWMyY1Q4dlpZeG93Rm9IKzllcXZ2Wm9YVzVsYjdDM3V5RDlWakgyZk55MHMKMUFkdmNpZ1duTmVRaEh6ZlI1d2xXMlNrcE0vQXlOYUNWZFZ2RHRhNnE5MitiSmxnZ0VKKzFnRUI2bUJBRi91YQo5U2FqQXhSQWE3YWVOaWtMeGd3ZVVIU0RBeVg5a3NraXNnOVVpTjYwNmRNLzY3VGsrQWxUeXFxWi93PT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="),
		"tls.key": []byte("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBdjZDbU9xSkZ6MWNjeEdvckRtUVpaUjNPWHVHaUx5S3Q3emNHTEhFaWl4d0J6RVlwCmd0Y0JMNFFUVEFZMWlkcGpURjh1d3lZNy9lZzJDcXM3ZnlMSHluREpNN0crOGs3dTdGbG8zbEpUbCtyYUpHa2UKL0JHU1Y4RlNkZVZ1R1MxbnBVZnljV1FSVkc2T3hYREdOZkFvVllzWmpUeHIzUDVZbVlyM0d5N2txK2U4a0Y3QgpocWhIdE9seTlEUlZzM2o4ZFloQ1V5UzViN2V2bnFUaGRzemlrTWlCcEdONmp5KzhXd3BSU1o1NVEyYXoyVkRzCk5tZkQ1L2lhYTNrVjluZ1czeHMxSnMyRU1adTMwMDA3T2RrblF4M0oyUkZvZyt6cmxWQkJ2M0w5VlVyd25TVHoKcE55b2wvT2NyZzdhZE9TKzZXa3hTc3E4cWxGWkxFZFpvc215NndJREFRQUJBb0lCQVFDVkdybDlVaHFqdEpLcgp1amg2WUNUcWF3RFREeG9WTnhURDE3cTBCZXZzOWdQb0lJZllTTmVoVTFGNGpEUklhV2R1VzNtVlcwQysxbHFHCmZxb3l5S3RRdCtXMmxZMlFHUjhMUko2MnJyUmd0dHE2RGhtUDVWUkxlQjlqb1B2RUYzSllSdDA4b2JKaVVneEIKVVBqSnEyNlc4VDhXaUhjZFk4TW81ZHBVaW01ZjJ2NUdGWitZNTdtdlg5RkxUOWFZODd3c1A0d0dkNUZqWVErUgoyQmJEZnVzRFlUL0xVeGJvVC9DYTk5b2NsWkRQakh2cnlPUzRLQXExZ0ViU0JnbnlrVTRmY3NLMTNJbFdzQmFPCm9TYWdQcTBLUzZza1J1aVZYK3dVNVVXSXd0WFd0SHZUOFVyOWptYmdib3RvbzdYY3QwRFFkWk0rOFV1QXM5NjIKYWVqU3d1dWhBb0dCQU5xZXlpRkhOM0dQZTFqUFlTdEdiaTUxNmJheXpMUnBuU1IxR25VaEVrcGwwdjlCRFVmYwpKaGZqaUtDZkNHOHJZRzk1UlRtNkFULzFZRGtuWXY4Z2M1QXRLTHNZSm91NGJyc1NTOGZNTGcvbGRXWXZZaTFzCkFTZ2JNeWpqMkdDcGN3eDNHd3pvL28rQkJrNlVKRUFGNHd0ODA5Z2Fsb3hnbktUeUFQN1BtdlZwQW9HQkFPQmsKWE9kQ3lhRjlKVDdpMHRyY283bHRBc2Ird1hidis0RUVWUzBJUWo5bUw1NEVMZTlVMEhmRExjUmRDU0t5ZitFSQp2WWU2aVdHaHdVcityc2lnMlBVWWRJZDgrRWdnc1Uydm0yeVM1dEN0TkV1R3B6NlE1U2prRlZCS2F2OVhMbHMrCjRnNXJWaXMvbW9DeTZwbko5UDM1dlVmUTJQV1hmZkpXQ3g3S1AzY3pBb0dCQU1lUUM1TTFIemRhY280dlA5UHgKQnNNQ2Y4VjJrY1plWWtQVlljRnAzdmhxMnFDSEVVaDNmWTV3OVZjcDFOa21EM0d5a2E2UVRIUEYyWUJTbzl1ZwpFOTJZVzRYdUZjR1ZLZjg2UkZLdDM1NURKMWVRQ1Y3TktJRWoweCtRWFFSZnFkWEhJN28xTmFwcGJRaHQwbWxlCjlsS1dNQXNrdWlpS1NIT0pOYjlrWTErQkFvR0FjenMyOWsrMjZhWXhsVXk0Q1VxYkRXTHN0VElvT2FMZG5oQ1MKaVJDZnJKMFdRT2hXaW53Y25oUHVFZFBSR0M5Z09qalowN1M0VGhuYUFHQXZjN29lRUNkaDJCNFdCanc3c1BCSQpPWVpxMzZqQ25USmwrbHhBUWpKMnU0ZXIwTHA5aE1BVEtHSjRtcmNMNmFGM2xrZys5cG5rV05mb1FwNXNRQ0Z0CmpuOC8vajhDZ1lBanA2VFU5M2dUdTBROE5BdXY0Yy8zSHMvS0JFVitDSm0rY2hRdUpnNFpFZjVVaDBucXhQb0YKb1JRT3JaZW5jRy8zOHVsci9nVFpSSG5yMGRmR0NoQjgxVUNOZndyKzlCRkFnM25CeGxCR3VFM2Y1aEhCRHkreQp6MkJZNEt4YzE0em1xZXVtUTJGWElsS01hUVhFNW1vWTM3VXo1Zk9Sam81U0ZtdHFIdGUzdWc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="),
	}

	opSA := kubeutil.SimpleSA("cilium-operator", ns)
	cSA := kubeutil.SimpleSA("cilium", ns)

	cas := kubeutil.Secret("cilium-ca", ns, caCerts)
	hus := kubeutil.Secret("hubble-ca-secret", ns, hubbleCACerts)
	huCert := kubeutil.Secret(
		"hubble-server-certs",
		ns,
		hubbleServerCerts,
	)

	return &CiliumApp{
		CACertSecret: cas,
		Config: Config(
			mergeConfigMapData(
				DefaultConfigData,
				ciliumCM,
			),
		),

		MattDamon:  Daemon,
		NodeInit:   NodeInit,
		AppSA:      cSA,
		AppBinding: kubeutil.SimpleCRB(cSA, CR),
		AppRole:    CR,

		OpSA:      opSA,
		OpBinding: kubeutil.SimpleCRB(opSA, OperatorCR),
		OpRole:    OperatorCR,
		Operator:  Operator,

		HubbleSecret: hus,
		HubbleSvc:    HubblePeerSvc,
		HubbleCert:   huCert,
	}
}

func mergeConfigMapData(orig, extra map[string]string) map[string]string {
	for k, v := range extra {
		orig[k] = v
	}

	return orig
}
