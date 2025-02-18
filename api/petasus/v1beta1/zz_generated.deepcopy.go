//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorConfig) DeepCopyInto(out *AcceleratorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorConfig.
func (in *AcceleratorConfig) DeepCopy() *AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := new(AcceleratorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcceleratorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorConfigList) DeepCopyInto(out *AcceleratorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AcceleratorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorConfigList.
func (in *AcceleratorConfigList) DeepCopy() *AcceleratorConfigList {
	if in == nil {
		return nil
	}
	out := new(AcceleratorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcceleratorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorConfigSpec) DeepCopyInto(out *AcceleratorConfigSpec) {
	*out = *in
	in.Providers.DeepCopyInto(&out.Providers)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorConfigSpec.
func (in *AcceleratorConfigSpec) DeepCopy() *AcceleratorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AcceleratorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorConfigStatus) DeepCopyInto(out *AcceleratorConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorConfigStatus.
func (in *AcceleratorConfigStatus) DeepCopy() *AcceleratorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AcceleratorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonsetsSpec) DeepCopyInto(out *DaemonsetsSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonsetsSpec.
func (in *DaemonsetsSpec) DeepCopy() *DaemonsetsSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonsetsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvidersSpec) DeepCopyInto(out *ProvidersSpec) {
	*out = *in
	if in.Rebellions != nil {
		in, out := &in.Rebellions, &out.Rebellions
		*out = new(RebellionsConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvidersSpec.
func (in *ProvidersSpec) DeepCopy() *ProvidersSpec {
	if in == nil {
		return nil
	}
	out := new(ProvidersSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebellionsConfigSpec) DeepCopyInto(out *RebellionsConfigSpec) {
	*out = *in
	if in.Daemonsets != nil {
		in, out := &in.Daemonsets, &out.Daemonsets
		*out = new(DaemonsetsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Driver != nil {
		in, out := &in.Driver, &out.Driver
		*out = new(RebellionsDriverSpec)
		**out = **in
	}
	if in.DevicePlugin != nil {
		in, out := &in.DevicePlugin, &out.DevicePlugin
		*out = new(RebellionsDevicePluginSpec)
		**out = **in
	}
	if in.MetricExporter != nil {
		in, out := &in.MetricExporter, &out.MetricExporter
		*out = new(RebellionsMetricExporterSpec)
		**out = **in
	}
	if in.NPUFeatureDiscovery != nil {
		in, out := &in.NPUFeatureDiscovery, &out.NPUFeatureDiscovery
		*out = new(RebellionsNPUFeatureDiscoverySpec)
		**out = **in
	}
	if in.VFIOManager != nil {
		in, out := &in.VFIOManager, &out.VFIOManager
		*out = new(VFIOManagerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebellionsConfigSpec.
func (in *RebellionsConfigSpec) DeepCopy() *RebellionsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(RebellionsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebellionsDevicePluginSpec) DeepCopyInto(out *RebellionsDevicePluginSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebellionsDevicePluginSpec.
func (in *RebellionsDevicePluginSpec) DeepCopy() *RebellionsDevicePluginSpec {
	if in == nil {
		return nil
	}
	out := new(RebellionsDevicePluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebellionsDriverSpec) DeepCopyInto(out *RebellionsDriverSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebellionsDriverSpec.
func (in *RebellionsDriverSpec) DeepCopy() *RebellionsDriverSpec {
	if in == nil {
		return nil
	}
	out := new(RebellionsDriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebellionsMetricExporterSpec) DeepCopyInto(out *RebellionsMetricExporterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebellionsMetricExporterSpec.
func (in *RebellionsMetricExporterSpec) DeepCopy() *RebellionsMetricExporterSpec {
	if in == nil {
		return nil
	}
	out := new(RebellionsMetricExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebellionsNPUFeatureDiscoverySpec) DeepCopyInto(out *RebellionsNPUFeatureDiscoverySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebellionsNPUFeatureDiscoverySpec.
func (in *RebellionsNPUFeatureDiscoverySpec) DeepCopy() *RebellionsNPUFeatureDiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(RebellionsNPUFeatureDiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VFIOManagerSpec) DeepCopyInto(out *VFIOManagerSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]EnvVar, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VFIOManagerSpec.
func (in *VFIOManagerSpec) DeepCopy() *VFIOManagerSpec {
	if in == nil {
		return nil
	}
	out := new(VFIOManagerSpec)
	in.DeepCopyInto(out)
	return out
}
