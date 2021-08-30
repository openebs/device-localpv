/*
Copyright 2021 The OpenEBS Authors

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/openebs/device-localpv/pkg/apis/openebs.io/device/v1alpha1"
	scheme "github.com/openebs/device-localpv/pkg/generated/clientset/device/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeviceNodesGetter has a method to return a DeviceNodeInterface.
// A group's client should implement this interface.
type DeviceNodesGetter interface {
	DeviceNodes(namespace string) DeviceNodeInterface
}

// DeviceNodeInterface has methods to work with DeviceNode resources.
type DeviceNodeInterface interface {
	Create(ctx context.Context, deviceNode *v1alpha1.DeviceNode, opts v1.CreateOptions) (*v1alpha1.DeviceNode, error)
	Update(ctx context.Context, deviceNode *v1alpha1.DeviceNode, opts v1.UpdateOptions) (*v1alpha1.DeviceNode, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DeviceNode, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DeviceNodeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeviceNode, err error)
	DeviceNodeExpansion
}

// deviceNodes implements DeviceNodeInterface
type deviceNodes struct {
	client rest.Interface
	ns     string
}

// newDeviceNodes returns a DeviceNodes
func newDeviceNodes(c *LocalV1alpha1Client, namespace string) *deviceNodes {
	return &deviceNodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deviceNode, and returns the corresponding deviceNode object, and an error if there is any.
func (c *deviceNodes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeviceNode, err error) {
	result = &v1alpha1.DeviceNode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicenodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeviceNodes that match those selectors.
func (c *deviceNodes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeviceNodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeviceNodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deviceNodes.
func (c *deviceNodes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devicenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deviceNode and creates it.  Returns the server's representation of the deviceNode, and an error, if there is any.
func (c *deviceNodes) Create(ctx context.Context, deviceNode *v1alpha1.DeviceNode, opts v1.CreateOptions) (result *v1alpha1.DeviceNode, err error) {
	result = &v1alpha1.DeviceNode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devicenodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deviceNode).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deviceNode and updates it. Returns the server's representation of the deviceNode, and an error, if there is any.
func (c *deviceNodes) Update(ctx context.Context, deviceNode *v1alpha1.DeviceNode, opts v1.UpdateOptions) (result *v1alpha1.DeviceNode, err error) {
	result = &v1alpha1.DeviceNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicenodes").
		Name(deviceNode.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deviceNode).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deviceNode and deletes it. Returns an error if one occurs.
func (c *deviceNodes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicenodes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deviceNodes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicenodes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deviceNode.
func (c *deviceNodes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeviceNode, err error) {
	result = &v1alpha1.DeviceNode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devicenodes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}