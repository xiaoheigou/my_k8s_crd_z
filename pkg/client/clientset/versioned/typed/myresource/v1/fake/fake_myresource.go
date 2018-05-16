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
	myresource_v1 "github.com/xiaoheigou/mycrd/pkg/apis/myresource/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyResources implements MyResourceInterface
type FakeMyResources struct {
	Fake *FakeTrstringerV1
	ns   string
}

var myresourcesResource = schema.GroupVersionResource{Group: "trstringer.com", Version: "v1", Resource: "myresources"}

var myresourcesKind = schema.GroupVersionKind{Group: "trstringer.com", Version: "v1", Kind: "MyResource"}

// Get takes name of the myResource, and returns the corresponding myResource object, and an error if there is any.
func (c *FakeMyResources) Get(name string, options v1.GetOptions) (result *myresource_v1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(myresourcesResource, c.ns, name), &myresource_v1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*myresource_v1.MyResource), err
}

// List takes label and field selectors, and returns the list of MyResources that match those selectors.
func (c *FakeMyResources) List(opts v1.ListOptions) (result *myresource_v1.MyResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(myresourcesResource, myresourcesKind, c.ns, opts), &myresource_v1.MyResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &myresource_v1.MyResourceList{}
	for _, item := range obj.(*myresource_v1.MyResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested myResources.
func (c *FakeMyResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(myresourcesResource, c.ns, opts))

}

// Create takes the representation of a myResource and creates it.  Returns the server's representation of the myResource, and an error, if there is any.
func (c *FakeMyResources) Create(myResource *myresource_v1.MyResource) (result *myresource_v1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(myresourcesResource, c.ns, myResource), &myresource_v1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*myresource_v1.MyResource), err
}

// Update takes the representation of a myResource and updates it. Returns the server's representation of the myResource, and an error, if there is any.
func (c *FakeMyResources) Update(myResource *myresource_v1.MyResource) (result *myresource_v1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(myresourcesResource, c.ns, myResource), &myresource_v1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*myresource_v1.MyResource), err
}

// Delete takes name of the myResource and deletes it. Returns an error if one occurs.
func (c *FakeMyResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(myresourcesResource, c.ns, name), &myresource_v1.MyResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(myresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &myresource_v1.MyResourceList{})
	return err
}

// Patch applies the patch and returns the patched myResource.
func (c *FakeMyResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *myresource_v1.MyResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(myresourcesResource, c.ns, name, data, subresources...), &myresource_v1.MyResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*myresource_v1.MyResource), err
}
