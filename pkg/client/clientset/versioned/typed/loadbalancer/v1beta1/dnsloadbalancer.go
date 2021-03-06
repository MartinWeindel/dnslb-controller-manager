/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1beta1 "github.com/gardener/dnslb-controller-manager/pkg/apis/loadbalancer/v1beta1"
	scheme "github.com/gardener/dnslb-controller-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSLoadBalancersGetter has a method to return a DNSLoadBalancerInterface.
// A group's client should implement this interface.
type DNSLoadBalancersGetter interface {
	DNSLoadBalancers(namespace string) DNSLoadBalancerInterface
}

// DNSLoadBalancerInterface has methods to work with DNSLoadBalancer resources.
type DNSLoadBalancerInterface interface {
	Create(*v1beta1.DNSLoadBalancer) (*v1beta1.DNSLoadBalancer, error)
	Update(*v1beta1.DNSLoadBalancer) (*v1beta1.DNSLoadBalancer, error)
	UpdateStatus(*v1beta1.DNSLoadBalancer) (*v1beta1.DNSLoadBalancer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.DNSLoadBalancer, error)
	List(opts v1.ListOptions) (*v1beta1.DNSLoadBalancerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.DNSLoadBalancer, err error)
	DNSLoadBalancerExpansion
}

// dNSLoadBalancers implements DNSLoadBalancerInterface
type dNSLoadBalancers struct {
	client rest.Interface
	ns     string
}

// newDNSLoadBalancers returns a DNSLoadBalancers
func newDNSLoadBalancers(c *LoadbalancerV1beta1Client, namespace string) *dNSLoadBalancers {
	return &dNSLoadBalancers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dNSLoadBalancer, and returns the corresponding dNSLoadBalancer object, and an error if there is any.
func (c *dNSLoadBalancers) Get(name string, options v1.GetOptions) (result *v1beta1.DNSLoadBalancer, err error) {
	result = &v1beta1.DNSLoadBalancer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSLoadBalancers that match those selectors.
func (c *dNSLoadBalancers) List(opts v1.ListOptions) (result *v1beta1.DNSLoadBalancerList, err error) {
	result = &v1beta1.DNSLoadBalancerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSLoadBalancers.
func (c *dNSLoadBalancers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a dNSLoadBalancer and creates it.  Returns the server's representation of the dNSLoadBalancer, and an error, if there is any.
func (c *dNSLoadBalancers) Create(dNSLoadBalancer *v1beta1.DNSLoadBalancer) (result *v1beta1.DNSLoadBalancer, err error) {
	result = &v1beta1.DNSLoadBalancer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		Body(dNSLoadBalancer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dNSLoadBalancer and updates it. Returns the server's representation of the dNSLoadBalancer, and an error, if there is any.
func (c *dNSLoadBalancers) Update(dNSLoadBalancer *v1beta1.DNSLoadBalancer) (result *v1beta1.DNSLoadBalancer, err error) {
	result = &v1beta1.DNSLoadBalancer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		Name(dNSLoadBalancer.Name).
		Body(dNSLoadBalancer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dNSLoadBalancers) UpdateStatus(dNSLoadBalancer *v1beta1.DNSLoadBalancer) (result *v1beta1.DNSLoadBalancer, err error) {
	result = &v1beta1.DNSLoadBalancer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		Name(dNSLoadBalancer.Name).
		SubResource("status").
		Body(dNSLoadBalancer).
		Do().
		Into(result)
	return
}

// Delete takes name of the dNSLoadBalancer and deletes it. Returns an error if one occurs.
func (c *dNSLoadBalancers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSLoadBalancers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dNSLoadBalancer.
func (c *dNSLoadBalancers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.DNSLoadBalancer, err error) {
	result = &v1beta1.DNSLoadBalancer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dnsloadbalancers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
