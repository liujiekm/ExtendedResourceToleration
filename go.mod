module ExtendedResourceToleration

go 1.13

require (
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/kubernetes v1.16.15
)

replace (
	k8s.io/api => ./staging/src/k8s.io/api
	k8s.io/apiextensions-apiserver => ./staging/src/k8s.io/apiextensions-apiserver
	k8s.io/apimachinery => ./staging/src/k8s.io/apimachinery
	k8s.io/apiserver => ./staging/src/k8s.io/apiserver
	k8s.io/cli-runtime => ./staging/src/k8s.io/cli-runtime
	k8s.io/client-go => ./staging/src/k8s.io/client-go
	k8s.io/cloud-provider => ./staging/src/k8s.io/cloud-provider
	k8s.io/cluster-bootstrap => ./staging/src/k8s.io/cluster-bootstrap
	k8s.io/code-generator => ./staging/src/k8s.io/code-generator
	k8s.io/component-base => ./staging/src/k8s.io/component-base
	k8s.io/cri-api => ./staging/src/k8s.io/cri-api
	k8s.io/csi-translation-lib => ./staging/src/k8s.io/csi-translation-lib
	k8s.io/gengo => k8s.io/gengo v0.0.0-20190822140433-26a664648505
	k8s.io/heapster => k8s.io/heapster v1.2.0-beta.1
	k8s.io/klog => k8s.io/klog v1.0.0
	k8s.io/kube-aggregator => ./staging/src/k8s.io/kube-aggregator
	k8s.io/kube-controller-manager => ./staging/src/k8s.io/kube-controller-manager
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410163147-594e756bea31 // release-1.16
	k8s.io/kube-proxy => ./staging/src/k8s.io/kube-proxy
	k8s.io/kube-scheduler => ./staging/src/k8s.io/kube-scheduler
	k8s.io/kubectl => ./staging/src/k8s.io/kubectl
	k8s.io/kubelet => ./staging/src/k8s.io/kubelet
	k8s.io/legacy-cloud-providers => ./staging/src/k8s.io/legacy-cloud-providers
	k8s.io/metrics => ./staging/src/k8s.io/metrics
	k8s.io/node-api => ./staging/src/k8s.io/node-api
	k8s.io/repo-infra => k8s.io/repo-infra v0.0.1-alpha.1
	k8s.io/sample-apiserver => ./staging/src/k8s.io/sample-apiserver
	k8s.io/sample-cli-plugin => ./staging/src/k8s.io/sample-cli-plugin
	k8s.io/sample-controller => ./staging/src/k8s.io/sample-controller
	k8s.io/utils => k8s.io/utils v0.0.0-20190801114015-581e00157fb1
)
