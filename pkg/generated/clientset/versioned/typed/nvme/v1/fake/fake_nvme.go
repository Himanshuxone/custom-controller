/*
Copyright The Kubernetes Authors.

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

package fake

import (
	"context"
	nvmev1 "custom-controller/pkg/apis/nvme/v1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNvmes implements NvmeInterface
type FakeNvmes struct {
	Fake *FakeNvmeV1
	ns   string
}

var nvmesResource = schema.GroupVersionResource{Group: "nvme.com", Version: "v1", Resource: "nvmes"}

var nvmesKind = schema.GroupVersionKind{Group: "nvme.com", Version: "v1", Kind: "Nvme"}

// Get takes name of the nvme, and returns the corresponding nvme object, and an error if there is any.
func (c *FakeNvmes) Get(ctx context.Context, name string, options v1.GetOptions) (result *nvmev1.Nvme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nvmesResource, c.ns, name), &nvmev1.Nvme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nvmev1.Nvme), err
}

// List takes label and field selectors, and returns the list of Nvmes that match those selectors.
func (c *FakeNvmes) List(ctx context.Context, opts v1.ListOptions) (result *nvmev1.NvmeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nvmesResource, nvmesKind, c.ns, opts), &nvmev1.NvmeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &nvmev1.NvmeList{ListMeta: obj.(*nvmev1.NvmeList).ListMeta}
	for _, item := range obj.(*nvmev1.NvmeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nvmes.
func (c *FakeNvmes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nvmesResource, c.ns, opts))

}

// Create takes the representation of a nvme and creates it.  Returns the server's representation of the nvme, and an error, if there is any.
func (c *FakeNvmes) Create(ctx context.Context, nvme *nvmev1.Nvme, opts v1.CreateOptions) (result *nvmev1.Nvme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nvmesResource, c.ns, nvme), &nvmev1.Nvme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nvmev1.Nvme), err
}

// Update takes the representation of a nvme and updates it. Returns the server's representation of the nvme, and an error, if there is any.
func (c *FakeNvmes) Update(ctx context.Context, nvme *nvmev1.Nvme, opts v1.UpdateOptions) (result *nvmev1.Nvme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nvmesResource, c.ns, nvme), &nvmev1.Nvme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nvmev1.Nvme), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNvmes) UpdateStatus(ctx context.Context, nvme *nvmev1.Nvme, opts v1.UpdateOptions) (*nvmev1.Nvme, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nvmesResource, "status", c.ns, nvme), &nvmev1.Nvme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nvmev1.Nvme), err
}

// Delete takes name of the nvme and deletes it. Returns an error if one occurs.
func (c *FakeNvmes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(nvmesResource, c.ns, name, opts), &nvmev1.Nvme{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNvmes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nvmesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &nvmev1.NvmeList{})
	return err
}

// Patch applies the patch and returns the patched nvme.
func (c *FakeNvmes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *nvmev1.Nvme, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nvmesResource, c.ns, name, pt, data, subresources...), &nvmev1.Nvme{})

	if obj == nil {
		return nil, err
	}
	return obj.(*nvmev1.Nvme), err
}