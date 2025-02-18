/*
Copyright 2025 SK Telecom.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AcceleratorConfigSpec defines the desired state of AcceleratorConfig
type AcceleratorConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Providers ProvidersSpec `json:"providers"`
}

// ProvidersSpec defines each providers' component configurations
type ProvidersSpec struct {
	// TODO: Nvidia gpu configuration
	// Nvidia *nvidiav1.ClusterPolicySpec `json:"nvidia,omitempty"`
	// Rebellions npu configuration
	Rebellions *RebellionsConfigSpec `json:"rebellions,omitempty"`
}

type RebellionsConfigSpec struct {
	// Daemonsets defines common configurations for rebellions daemonsets
	Daemonsets *DaemonsetsSpec `json:"daemonsets,omitempty"`
	// Driver spec
	Driver *RebellionsDriverSpec `json:"driver,omitempty"`
	// DevicePlugin component spec
	DevicePlugin *RebellionsDevicePluginSpec `json:"devicePlugin,omitempty"`
	// MetricExporter spec
	MetricExporter *RebellionsMetricExporterSpec `json:"metricExporter,omitempty"`
	// NPUFeatureDiscovery spec
	NPUFeatureDiscovery *RebellionsNPUFeatureDiscoverySpec `json:"nfd,omitempty"`
	// VFIOManager component spec
	VFIOManager *VFIOManagerSpec `json:"vfioManager,omitempty"`
}

// DaemonsetsSpec indicates common configuration for all Daemonsets managed by Accelerator Operator
type DaemonsetsSpec struct {
	// Set labels
	Labels map[string]string `json:"labels,omitempty"`
	// Set annotations
	Annotations map[string]string `json:"annotations,omitempty"`
	// Set tolerations
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
	// +kubebuilder:validation:Optional
	PriorityClassName string `json:"priorityClassName,omitempty"`
}

type RebellionsDriverSpec struct {
	Enabled bool `json:"enabled"`
}

type RebellionsDevicePluginSpec struct {
	Enabled bool `json:"enabled"`
}

type RebellionsMetricExporterSpec struct {
	Enabled bool `json:"enabled"`
}

type RebellionsNPUFeatureDiscoverySpec struct {
	Enabled bool `json:"enabled"`
}

type VFIOManagerSpec struct {
	// Flag which indicates deploying VFIO Manager with operator is enabled
	Enabled bool `json:"enabled"`
	// VFIO manager container image
	Image *ImageSpec `json:"image,omitempty"`
	// container resource spec
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// additional container environments
	Env []EnvVar `json:"env,omitempty"`
}

type ImageSpec struct {
	// Image registry
	// +kubebuilder:validation:Optional
	Registry string `json:"registry,omitempty"`

	// Image name
	// +kubebuilder:validation:Pattern=[a-zA-Z0-9\-]+
	Name string `json:"name,omitempty"`

	// Image tag
	// +kubebuilder:validation:Optional
	Tag string `json:"tag,omitempty"`

	// Image pull policy
	// +kubebuilder:validation:Optional
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	// Image pull secrets
	// +kubebuilder:validation:Optional
	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`
}

type EnvVar struct {
	// Name of the environment variable.
	Name string `json:"name"`

	// Value of the environment variable.
	Value string `json:"value,omitempty"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed.
	// +kubebuilder:validation:Optional
	Limits corev1.ResourceList `json:"limits,omitempty"`
	// Requests describes the minimum amount of compute resources required.
	// +kubebuilder:validation:Optional
	Requests corev1.ResourceList `json:"requests,omitempty"`
}

// AcceleratorConfigStatus defines the observed state of AcceleratorConfig
type AcceleratorConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:scope=Cluster,shortName={"ac","accelconfig","accelconf"}

// AcceleratorConfig is the Schema for the acceleratorconfigs API
type AcceleratorConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AcceleratorConfigSpec   `json:"spec,omitempty"`
	Status AcceleratorConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AcceleratorConfigList contains a list of AcceleratorConfig
type AcceleratorConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AcceleratorConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AcceleratorConfig{}, &AcceleratorConfigList{})
}
