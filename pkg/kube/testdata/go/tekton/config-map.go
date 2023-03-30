// Code generated by lingon. EDIT AS MUCH AS YOU LIKE.

package tekton

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var BundleresolverConfigCM = &corev1.ConfigMap{
	Data: map[string]string{
		// The default layer kind in the bundle image.
		"default-kind": "task",
		// the default service account name to use for bundle requests.
		"default-service-account": "default",
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "bundleresolver-config",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ClusterResolverConfigCM = &corev1.ConfigMap{
	Data: map[string]string{
		// An optional comma-separated list of namespaces which the resolver is allowed to access. Defaults to empty, meaning all namespaces are allowed.
		"allowed-namespaces": "",
		// An optional comma-separated list of namespaces which the resolver is blocked from accessing. Defaults to empty, meaning all namespaces are allowed.
		"blocked-namespaces": "",
		// The default kind to fetch.
		"default-kind": "task",
		// The default namespace to look for resources in.
		"default-namespace": "",
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "cluster-resolver-config",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigArtifactBucketCM = &corev1.ConfigMap{
	Data: nil,
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-artifact-bucket",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigArtifactPvcCM = &corev1.ConfigMap{
	Data: nil,
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-artifact-pvc",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigDefaultsCM = &corev1.ConfigMap{
	Data: map[string]string{
		"_example": `
################################
#                              #
#    EXAMPLE CONFIGURATION     #
#                              #
################################

# This block is not actually functional configuration,
# but serves to illustrate the available configuration
# options and document them in a way that is accessible
# to users that "kubectl edit" this config map.
#
# These sample configuration options may be copied out of
# this example block and unindented to be in the data block
# to actually change the configuration.

# default-timeout-minutes contains the default number of
# minutes to use for TaskRun and PipelineRun, if none is specified.
default-timeout-minutes: "60"  # 60 minutes

# default-service-account contains the default service account name
# to use for TaskRun and PipelineRun, if none is specified.
default-service-account: "default"

# default-managed-by-label-value contains the default value given to the
# "app.kubernetes.io/managed-by" label applied to all Pods created for
# TaskRuns. If a user's requested TaskRun specifies another value for this
# label, the user's request supercedes.
default-managed-by-label-value: "tekton-pipelines"

# default-pod-template contains the default pod template to use for
# TaskRun and PipelineRun. If a pod template is specified on the
# PipelineRun, the default-pod-template is merged with that one.
# default-pod-template:

# default-affinity-assistant-pod-template contains the default pod template
# to use for affinity assistant pods. If a pod template is specified on the
# PipelineRun, the default-affinity-assistant-pod-template is merged with
# that one.
# default-affinity-assistant-pod-template:

# default-cloud-events-sink contains the default CloudEvents sink to be
# used for TaskRun and PipelineRun, when no sink is specified.
# Note that right now it is still not possible to set a PipelineRun or
# TaskRun specific sink, so the default is the only option available.
# If no sink is specified, no CloudEvent is generated
# default-cloud-events-sink:

# default-task-run-workspace-binding contains the default workspace
# configuration provided for any Workspaces that a Task declares
# but that a TaskRun does not explicitly provide.
# default-task-run-workspace-binding: |
#   emptyDir: {}

# default-max-matrix-combinations-count contains the default maximum number
# of combinations from a Matrix, if none is specified.
default-max-matrix-combinations-count: "256"

# default-forbidden-env contains comma seperated environment variables that cannot be
# overridden by podTemplate.
default-forbidden-env:

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-defaults",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigLeaderElectionCM = &corev1.ConfigMap{
	Data: map[string]string{
		"_example": `
################################
#                              #
#    EXAMPLE CONFIGURATION     #
#                              #
################################
# This block is not actually functional configuration,
# but serves to illustrate the available configuration
# options and document them in a way that is accessible
# to users that "kubectl edit" this config map.
#
# These sample configuration options may be copied out of
# this example block and unindented to be in the data block
# to actually change the configuration.
# lease-duration is how long non-leaders will wait to try to acquire the
# lock; 15 seconds is the value used by core kubernetes controllers.
lease-duration: "60s"
# renew-deadline is how long a leader will try to renew the lease before
# giving up; 10 seconds is the value used by core kubernetes controllers.
renew-deadline: "40s"
# retry-period is how long the leader election client waits between tries of
# actions; 2 seconds is the value used by core kubernetes controllers.
retry-period: "10s"
# buckets is the number of buckets used to partition key space of each
# Reconciler. If this number is M and the replica number of the controller
# is N, the N replicas will compete for the M buckets. The owner of a
# bucket will take care of the reconciling for the keys partitioned into
# that bucket.
buckets: "1"

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-leader-election",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigLeaderElectionCM1 = &corev1.ConfigMap{
	Data: map[string]string{
		"_example": `
################################
#                              #
#    EXAMPLE CONFIGURATION     #
#                              #
################################
# This block is not actually functional configuration,
# but serves to illustrate the available configuration
# options and document them in a way that is accessible
# to users that "kubectl edit" this config map.
#
# These sample configuration options may be copied out of
# this example block and unindented to be in the data block
# to actually change the configuration.
# lease-duration is how long non-leaders will wait to try to acquire the
# lock; 15 seconds is the value used by core kubernetes controllers.
lease-duration: "60s"
# renew-deadline is how long a leader will try to renew the lease before
# giving up; 10 seconds is the value used by core kubernetes controllers.
renew-deadline: "40s"
# retry-period is how long the leader election client waits between tries of
# actions; 2 seconds is the value used by core kubernetes controllers.
retry-period: "10s"
# buckets is the number of buckets used to partition key space of each
# Reconciler. If this number is M and the replica number of the controller
# is N, the N replicas will compete for the M buckets. The owner of a
# bucket will take care of the reconciling for the keys partitioned into
# that bucket.
buckets: "1"

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "config-leader-election",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigLoggingCM = &corev1.ConfigMap{
	Data: map[string]string{
		// Log level overrides
		"loglevel.controller": "info",
		"loglevel.webhook":    "info",
		"zap-logger-config": `
{
  "level": "info",
  "development": false,
  "sampling": {
    "initial": 100,
    "thereafter": 100
  },
  "outputPaths": ["stdout"],
  "errorOutputPaths": ["stderr"],
  "encoding": "json",
  "encoderConfig": {
    "timeKey": "timestamp",
    "levelKey": "severity",
    "nameKey": "logger",
    "callerKey": "caller",
    "messageKey": "message",
    "stacktraceKey": "stacktrace",
    "lineEnding": "",
    "levelEncoder": "",
    "timeEncoder": "iso8601",
    "durationEncoder": "",
    "callerEncoder": ""
  }
}

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-logging",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigLoggingCM2 = &corev1.ConfigMap{
	Data: map[string]string{
		// Log level overrides
		"loglevel.controller": "info",
		"loglevel.webhook":    "info",
		"zap-logger-config": `
{
  "level": "info",
  "development": false,
  "sampling": {
    "initial": 100,
    "thereafter": 100
  },
  "outputPaths": ["stdout"],
  "errorOutputPaths": ["stderr"],
  "encoding": "json",
  "encoderConfig": {
    "timeKey": "timestamp",
    "levelKey": "severity",
    "nameKey": "logger",
    "callerKey": "caller",
    "messageKey": "message",
    "stacktraceKey": "stacktrace",
    "lineEnding": "",
    "levelEncoder": "",
    "timeEncoder": "iso8601",
    "durationEncoder": "",
    "callerEncoder": ""
  }
}

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "config-logging",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigObservabilityCM = &corev1.ConfigMap{
	Data: map[string]string{
		"_example": `
################################
#                              #
#    EXAMPLE CONFIGURATION     #
#                              #
################################

# This block is not actually functional configuration,
# but serves to illustrate the available configuration
# options and document them in a way that is accessible
# to users that "kubectl edit" this config map.
#
# These sample configuration options may be copied out of
# this example block and unindented to be in the data block
# to actually change the configuration.

# metrics.backend-destination field specifies the system metrics destination.
# It supports either prometheus (the default) or stackdriver.
# Note: Using Stackdriver will incur additional charges.
metrics.backend-destination: prometheus

# metrics.stackdriver-project-id field specifies the Stackdriver project ID. This
# field is optional. When running on GCE, application default credentials will be
# used and metrics will be sent to the cluster's project if this field is
# not provided.
metrics.stackdriver-project-id: "<your stackdriver project id>"

# metrics.allow-stackdriver-custom-metrics indicates whether it is allowed
# to send metrics to Stackdriver using "global" resource type and custom
# metric type. Setting this flag to "true" could cause extra Stackdriver
# charge.  If metrics.backend-destination is not Stackdriver, this is
# ignored.
metrics.allow-stackdriver-custom-metrics: "false"
metrics.taskrun.level: "task"
metrics.taskrun.duration-type: "histogram"
metrics.pipelinerun.level: "pipeline"
metrics.pipelinerun.duration-type: "histogram"

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-observability",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigObservabilityCM3 = &corev1.ConfigMap{
	Data: map[string]string{
		"_example": `
################################
#                              #
#    EXAMPLE CONFIGURATION     #
#                              #
################################

# This block is not actually functional configuration,
# but serves to illustrate the available configuration
# options and document them in a way that is accessible
# to users that "kubectl edit" this config map.
#
# These sample configuration options may be copied out of
# this example block and unindented to be in the data block
# to actually change the configuration.

# metrics.backend-destination field specifies the system metrics destination.
# It supports either prometheus (the default) or stackdriver.
# Note: Using stackdriver will incur additional charges
metrics.backend-destination: prometheus

# metrics.request-metrics-backend-destination specifies the request metrics
# destination. If non-empty, it enables queue proxy to send request metrics.
# Currently supported values: prometheus, stackdriver.
metrics.request-metrics-backend-destination: prometheus

# metrics.stackdriver-project-id field specifies the stackdriver project ID. This
# field is optional. When running on GCE, application default credentials will be
# used if this field is not provided.
metrics.stackdriver-project-id: "<your stackdriver project id>"

# metrics.allow-stackdriver-custom-metrics indicates whether it is allowed to send metrics to
# Stackdriver using "global" resource type and custom metric type if the
# metrics are not supported by "knative_revision" resource type. Setting this
# flag to "true" could cause extra Stackdriver charge.
# If metrics.backend-destination is not Stackdriver, this is ignored.
metrics.allow-stackdriver-custom-metrics: "false"

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "config-observability",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigRegistryCertCM = &corev1.ConfigMap{
	Data: nil,
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-registry-cert",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigSpireCM = &corev1.ConfigMap{
	Data: map[string]string{
		"_example": `
################################
#                              #
#    EXAMPLE CONFIGURATION     #
#                              #
################################
# This block is not actually functional configuration,
# but serves to illustrate the available configuration
# options and document them in a way that is accessible
# to users that "kubectl edit" this config map.
#
# These sample configuration options may be copied out of
# this example block and unindented to be in the data block
# to actually change the configuration.
#
# spire-trust-domain specifies the SPIRE trust domain to use.
# spire-trust-domain: "example.org"
#
# spire-socket-path specifies the SPIRE agent socket for SPIFFE workload API.
# spire-socket-path: "unix:///spiffe-workload-api/spire-agent.sock"
#
# spire-server-addr specifies the SPIRE server address for workload/node registration.
# spire-server-addr: "spire-server.spire.svc.cluster.local:8081"
#
# spire-node-alias-prefix specifies the SPIRE node alias prefix to use.
# spire-node-alias-prefix: "/tekton-node/"

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-spire",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ConfigTrustedResourcesCM = &corev1.ConfigMap{
	Data: map[string]string{
		"_example": `
################################
#                              #
#    EXAMPLE CONFIGURATION     #
#                              #
################################
# This block is not actually functional configuration,
# but serves to illustrate the available configuration
# options and document them in a way that is accessible
# to users that "kubectl edit" this config map.
#
# These sample configuration options may be copied out of
# this example block and unindented to be in the data block
# to actually change the configuration.

# publickeys specifies the list of public keys, the paths are separated by comma
# publickeys: "/etc/verification-secrets/cosign.pub,
# gcpkms://projects/tekton/locations/us/keyRings/trusted-resources/cryptoKeys/trusted-resources"

`,
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "config-trusted-resources",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var FeatureFlagsCM = &corev1.ConfigMap{
	Data: map[string]string{
		/*
		   Setting this flag to "false" will stop Tekton from waiting for a
		   TaskRun's sidecar containers to be running before starting the first
		   step. This will allow Tasks to be run in environments that don't
		   support the DownwardAPI volume type, but may lead to unintended
		   behaviour if sidecars are used.
		   #
		   See https://github.com/tektoncd/pipeline/issues/4937 for more info.
		*/
		"await-sidecar-readiness": "true",
		/*
		   Setting this flag will determine the version for custom tasks created by PipelineRuns.
		   Acceptable values are "v1beta1" and "v1alpha1".
		   The default is "v1beta1".
		*/
		"custom-task-version": "v1beta1",
		/*
		   Setting this flag to "true" will prevent Tekton to create an
		   Affinity Assistant for every TaskRun sharing a PVC workspace
		   #
		   The default behaviour is for Tekton to create Affinity Assistants
		   #
		   See more in the workspace documentation about Affinity Assistant
		   https://github.com/tektoncd/pipeline/blob/main/docs/workspaces.md#affinity-assistant-and-specifying-workspace-order-in-a-pipeline
		   or https://github.com/tektoncd/pipeline/pull/2630 for more info.
		*/
		"disable-affinity-assistant": "false",
		/*
		   Setting this flag to "true" will prevent Tekton scanning attached
		   service accounts and injecting any credentials it finds into your
		   Steps.
		   #
		   The default behaviour currently is for Tekton to search service
		   accounts for secrets matching a specified format and automatically
		   mount those into your Steps.
		   #
		   Note: setting this to "true" will prevent PipelineResources from
		   working.
		   #
		   See https://github.com/tektoncd/pipeline/issues/2791 for more
		   info.
		*/
		"disable-creds-init": "false",
		/*
		   Setting this flag will determine which gated features are enabled.
		   Acceptable values are "stable", "beta", or "alpha".
		*/
		"enable-api-fields": "stable",
		/*
		   Setting this flag to "true" enables populating the "provenance" field in TaskRun
		   and PipelineRun status. This field contains metadata about resources used
		   in the TaskRun/PipelineRun such as the source from where a remote Task/Pipeline
		   definition was fetched.
		*/
		"enable-provenance-in-status": "false",
		/*
		   Setting this flag to "true" enables the use of Tekton OCI bundle.
		   This is an experimental feature and thus should still be considered
		   an alpha feature.
		*/
		"enable-tekton-oci-bundles": "false",
		/*
		   Setting this flag will determine how Tekton pipelines will handle non-falsifiable provenance.
		   If set to "spire", then SPIRE will be used to ensure non-falsifiable provenance.
		   If set to "none", then Tekton will not have non-falsifiable provenance.
		   This is an experimental feature and thus should still be considered an alpha feature.
		*/
		"enforce-nonfalsifiablity": "none",
		/*
		   Setting this flag to "true" will require that any Git SSH Secret
		   offered to Tekton must have known_hosts included.
		   #
		   See https://github.com/tektoncd/pipeline/issues/2981 for more
		   info.
		*/
		"require-git-ssh-secret-known-hosts": "false",
		/*
		   Setting this flag to "enforce" will enforce verification of tasks/pipeline. Failing to verify
		   will fail the taskrun/pipelinerun. "warn" will only log the err message and "skip"
		   will skip the whole verification
		*/
		"resource-verification-mode": "skip",
		/*
		   This option should be set to false when Pipelines is running in a
		   cluster that does not use injected sidecars such as Istio. Setting
		   it to false should decrease the time it takes for a TaskRun to start
		   running. For clusters that use injected sidecars, setting this
		   option to false can lead to unexpected behavior.
		   #
		   See https://github.com/tektoncd/pipeline/issues/2080 for more info.
		*/
		"running-in-environment-with-injected-sidecars": "true",
		/*
		   Setting this flag to "true" enables CloudEvents for CustomRuns and Runs, as long as a
		   CloudEvents sink is configured in the config-defaults config map
		*/
		"send-cloudevents-for-runs": "false",
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "feature-flags",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var GitResolverConfigCM = &corev1.ConfigMap{
	Data: map[string]string{
		// The key in the API token secret containing the actual token. Required when using the authenticated API.
		"api-token-secret-key": "",
		// The Kubernetes secret containing the API token for the SCM provider. Required when using the authenticated API.
		"api-token-secret-name": "",
		// The namespace containing the API token secret. Defaults to "default".
		"api-token-secret-namespace": "default",
		/*
		   The default organization to look for repositories under when using the authenticated API,
		   if not specified in the resolver parameters. Optional.
		*/
		"default-org": "",
		// The git revision to fetch the remote resource from with either anonymous cloning or the authenticated API.
		"default-revision": "main",
		// The git url to fetch the remote resource from when using anonymous cloning.
		"default-url": "https://github.com/tektoncd/catalog.git",
		// The maximum amount of time a single anonymous cloning resolution may take.
		"fetch-timeout": "1m",
		// The SCM type to use with the authenticated API. Can be github, gitlab, gitea, bitbucketserver, bitbucketcloud
		"scm-type": "github",
		// The SCM server URL to use with the authenticated API. Not needed when using github.com, gitlab.com, or BitBucket Cloud
		"server-url": "",
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "git-resolver-config",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var HubresolverConfigCM = &corev1.ConfigMap{
	Data: map[string]string{
		// the default Artifact Hub Pipeline catalog from where to pull the resource.
		"default-artifact-hub-pipeline-catalog": "tekton-catalog-pipelines",
		// the default Artifact Hub Task catalog from where to pull the resource.
		"default-artifact-hub-task-catalog": "tekton-catalog-tasks",
		// the default layer kind in the hub image.
		"default-kind": "task",
		// the default Tekton Hub catalog from where to pull the resource.
		"default-tekton-hub-catalog": "Tekton",
		// the default hub source to pull the resource from.
		"default-type": "artifact",
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "hubresolver-config",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var PipelinesInfoCM = &corev1.ConfigMap{
	Data: map[string]string{
		/*
		   Contains pipelines version which can be queried by external
		   tools such as CLI. Elevated permissions are already given to
		   this ConfigMap such that even if we don't have access to
		   other resources in the namespace we still can have access to
		   this ConfigMap.
		*/
		"version": "v0.45.0",
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/instance": "default",
			"app.kubernetes.io/part-of":  "tekton-pipelines",
		},
		Name:      "pipelines-info",
		Namespace: "tekton-pipelines",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}

var ResolversFeatureFlagsCM = &corev1.ConfigMap{
	Data: map[string]string{
		// Setting this flag to "true" enables remote resolution of Tekton OCI bundles.
		"enable-bundles-resolver": "true",
		// Setting this flag to "true" enables remote resolution of tasks and pipelines from other namespaces within the cluster.
		"enable-cluster-resolver": "true",
		// Setting this flag to "true" enables remote resolution of tasks and pipelines from Git repositories.
		"enable-git-resolver": "true",
		// Setting this flag to "true" enables remote resolution of tasks and pipelines via the Tekton Hub.
		"enable-hub-resolver": "true",
	},
	ObjectMeta: metav1.ObjectMeta{
		Labels: map[string]string{
			"app.kubernetes.io/component": "resolvers",
			"app.kubernetes.io/instance":  "default",
			"app.kubernetes.io/part-of":   "tekton-pipelines",
		},
		Name:      "resolvers-feature-flags",
		Namespace: "tekton-pipelines-resolvers",
	},
	TypeMeta: metav1.TypeMeta{
		APIVersion: "v1",
		Kind:       "ConfigMap",
	},
}