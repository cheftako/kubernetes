/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package options

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	componentbaseconfig "k8s.io/component-base/config"
)

// ServiceControllerConfiguration contains elements describing ServiceController.
type ServiceControllerConfiguration struct {
	// concurrentServiceSyncs is the number of services that are
	// allowed to sync concurrently. Larger number = more responsive service
	// management, but more CPU (and network) load.
	ConcurrentServiceSyncs int32
}

// GenericControllerManagerConfiguration holds configuration for a generic controller-manager
type GenericControllerManagerConfiguration struct {
	// port is the port that the controller-manager's http service runs on.
	Port int32
	// address is the IP address to serve on (set to 0.0.0.0 for all interfaces).
	Address string
	// minResyncPeriod is the resync period in reflectors; will be random between
	// minResyncPeriod and 2*minResyncPeriod.
	MinResyncPeriod metav1.Duration
	// ClientConnection specifies the kubeconfig file and client connection
	// settings for the proxy server to use when communicating with the apiserver.
	ClientConnection componentbaseconfig.ClientConnectionConfiguration
	// How long to wait between starting controller managers
	ControllerStartInterval metav1.Duration
	// leaderElection defines the configuration of leader election client.
	LeaderElection componentbaseconfig.LeaderElectionConfiguration
	// Controllers is the list of controllers to enable or disable
	// '*' means "all enabled by default controllers"
	// 'foo' means "enable 'foo'"
	// '-foo' means "disable 'foo'"
	// first item for a particular name wins
	Controllers []string
	// DebuggingConfiguration holds configuration for Debugging related features.
	Debugging componentbaseconfig.DebuggingConfiguration
}

// KubeCloudSharedConfiguration contains elements shared by both kube-controller manager
// and cloud-controller manager, but not genericconfig.
type KubeCloudSharedConfiguration struct {
	// CloudProviderConfiguration holds configuration for CloudProvider related features.
	CloudProvider CloudProviderConfiguration
	// externalCloudVolumePlugin specifies the plugin to use when cloudProvider is "external".
	// It is currently used by the in repo cloud providers to handle node and volume control in the KCM.
	ExternalCloudVolumePlugin string
	// useServiceAccountCredentials indicates whether controllers should be run with
	// individual service account credentials.
	UseServiceAccountCredentials bool
	// run with untagged cloud instances
	AllowUntaggedCloud bool
	// routeReconciliationPeriod is the period for reconciling routes created for Nodes by cloud provider..
	RouteReconciliationPeriod metav1.Duration
	// nodeMonitorPeriod is the period for syncing NodeStatus in NodeController.
	NodeMonitorPeriod metav1.Duration
	// clusterName is the instance prefix for the cluster.
	ClusterName string
	// clusterCIDR is CIDR Range for Pods in cluster.
	ClusterCIDR string
	// AllocateNodeCIDRs enables CIDRs for Pods to be allocated and, if
	// ConfigureCloudRoutes is true, to be set on the cloud provider.
	AllocateNodeCIDRs bool
	// CIDRAllocatorType determines what kind of pod CIDR allocator will be used.
	CIDRAllocatorType string
	// configureCloudRoutes enables CIDRs allocated with allocateNodeCIDRs
	// to be configured on the cloud provider.
	ConfigureCloudRoutes bool
	// nodeSyncPeriod is the period for syncing nodes from cloudprovider. Longer
	// periods will result in fewer calls to cloud provider, but may delay addition
	// of new nodes to cluster.
	NodeSyncPeriod metav1.Duration
}

// CloudProviderConfiguration contains basically elements about cloud provider.
type CloudProviderConfiguration struct {
	// Name is the provider for cloud services.
	Name string
	// cloudConfigFile is the path to the cloud provider configuration file.
	CloudConfigFile string
}
