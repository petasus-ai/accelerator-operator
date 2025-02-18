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
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	petasusv1beta1 "github.com/petasus-ai/accelerator-operator/api/petasus/v1beta1"
	scheme "github.com/petasus-ai/accelerator-operator/api/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AcceleratorConfigsGetter has a method to return a AcceleratorConfigInterface.
// A group's client should implement this interface.
type AcceleratorConfigsGetter interface {
	AcceleratorConfigs() AcceleratorConfigInterface
}

// AcceleratorConfigInterface has methods to work with AcceleratorConfig resources.
type AcceleratorConfigInterface interface {
	Create(ctx context.Context, acceleratorConfig *petasusv1beta1.AcceleratorConfig, opts v1.CreateOptions) (*petasusv1beta1.AcceleratorConfig, error)
	Update(ctx context.Context, acceleratorConfig *petasusv1beta1.AcceleratorConfig, opts v1.UpdateOptions) (*petasusv1beta1.AcceleratorConfig, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, acceleratorConfig *petasusv1beta1.AcceleratorConfig, opts v1.UpdateOptions) (*petasusv1beta1.AcceleratorConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*petasusv1beta1.AcceleratorConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*petasusv1beta1.AcceleratorConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *petasusv1beta1.AcceleratorConfig, err error)
	AcceleratorConfigExpansion
}

// acceleratorConfigs implements AcceleratorConfigInterface
type acceleratorConfigs struct {
	*gentype.ClientWithList[*petasusv1beta1.AcceleratorConfig, *petasusv1beta1.AcceleratorConfigList]
}

// newAcceleratorConfigs returns a AcceleratorConfigs
func newAcceleratorConfigs(c *PetasusV1beta1Client) *acceleratorConfigs {
	return &acceleratorConfigs{
		gentype.NewClientWithList[*petasusv1beta1.AcceleratorConfig, *petasusv1beta1.AcceleratorConfigList](
			"acceleratorconfigs",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *petasusv1beta1.AcceleratorConfig { return &petasusv1beta1.AcceleratorConfig{} },
			func() *petasusv1beta1.AcceleratorConfigList { return &petasusv1beta1.AcceleratorConfigList{} },
		),
	}
}
