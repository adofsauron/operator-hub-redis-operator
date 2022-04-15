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

package v1alpha1

import (
	"context"
	v1alpha1 "td-redis-operator/pkg/apis/cache/v1alpha1"
	scheme "td-redis-operator/pkg/client/clientset/scheme"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedisStandbiesGetter has a method to return a RedisStandbyInterface.
// A group's client should implement this interface.
type RedisStandbiesGetter interface {
	RedisStandbies(namespace string) RedisStandbyInterface
}

// RedisStandbyInterface has methods to work with RedisStandby resources.
type RedisStandbyInterface interface {
	Create(ctx context.Context, redisStandby *v1alpha1.RedisStandby, opts v1.CreateOptions) (*v1alpha1.RedisStandby, error)
	Update(ctx context.Context, redisStandby *v1alpha1.RedisStandby, opts v1.UpdateOptions) (*v1alpha1.RedisStandby, error)
	UpdateStatus(ctx context.Context, redisStandby *v1alpha1.RedisStandby, opts v1.UpdateOptions) (*v1alpha1.RedisStandby, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RedisStandby, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RedisStandbyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RedisStandby, err error)
	RedisStandbyExpansion
}

// redisStandbies implements RedisStandbyInterface
type redisStandbies struct {
	client rest.Interface
	ns     string
}

// newRedisStandbies returns a RedisStandbies
func newRedisStandbies(c *CacheV1alpha1Client, namespace string) *redisStandbies {
	return &redisStandbies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the redisStandby, and returns the corresponding redisStandby object, and an error if there is any.
func (c *redisStandbies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RedisStandby, err error) {
	result = &v1alpha1.RedisStandby{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisstandbies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisStandbies that match those selectors.
func (c *redisStandbies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RedisStandbyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RedisStandbyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("redisstandbies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisStandbies.
func (c *redisStandbies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("redisstandbies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a redisStandby and creates it.  Returns the server's representation of the redisStandby, and an error, if there is any.
func (c *redisStandbies) Create(ctx context.Context, redisStandby *v1alpha1.RedisStandby, opts v1.CreateOptions) (result *v1alpha1.RedisStandby, err error) {
	result = &v1alpha1.RedisStandby{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("redisstandbies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisStandby).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a redisStandby and updates it. Returns the server's representation of the redisStandby, and an error, if there is any.
func (c *redisStandbies) Update(ctx context.Context, redisStandby *v1alpha1.RedisStandby, opts v1.UpdateOptions) (result *v1alpha1.RedisStandby, err error) {
	result = &v1alpha1.RedisStandby{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisstandbies").
		Name(redisStandby.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisStandby).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *redisStandbies) UpdateStatus(ctx context.Context, redisStandby *v1alpha1.RedisStandby, opts v1.UpdateOptions) (result *v1alpha1.RedisStandby, err error) {
	result = &v1alpha1.RedisStandby{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("redisstandbies").
		Name(redisStandby.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(redisStandby).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the redisStandby and deletes it. Returns an error if one occurs.
func (c *redisStandbies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisstandbies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisStandbies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("redisstandbies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched redisStandby.
func (c *redisStandbies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RedisStandby, err error) {
	result = &v1alpha1.RedisStandby{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("redisstandbies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}