/*
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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/servicemeshinterface/smi-sdk-go/pkg/apis/specs/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TCPRouteLister helps list TCPRoutes.
type TCPRouteLister interface {
	// List lists all TCPRoutes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.TCPRoute, err error)
	// TCPRoutes returns an object that can list and get TCPRoutes.
	TCPRoutes(namespace string) TCPRouteNamespaceLister
	TCPRouteListerExpansion
}

// tCPRouteLister implements the TCPRouteLister interface.
type tCPRouteLister struct {
	indexer cache.Indexer
}

// NewTCPRouteLister returns a new TCPRouteLister.
func NewTCPRouteLister(indexer cache.Indexer) TCPRouteLister {
	return &tCPRouteLister{indexer: indexer}
}

// List lists all TCPRoutes in the indexer.
func (s *tCPRouteLister) List(selector labels.Selector) (ret []*v1alpha3.TCPRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.TCPRoute))
	})
	return ret, err
}

// TCPRoutes returns an object that can list and get TCPRoutes.
func (s *tCPRouteLister) TCPRoutes(namespace string) TCPRouteNamespaceLister {
	return tCPRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TCPRouteNamespaceLister helps list and get TCPRoutes.
type TCPRouteNamespaceLister interface {
	// List lists all TCPRoutes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.TCPRoute, err error)
	// Get retrieves the TCPRoute from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.TCPRoute, error)
	TCPRouteNamespaceListerExpansion
}

// tCPRouteNamespaceLister implements the TCPRouteNamespaceLister
// interface.
type tCPRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TCPRoutes in the indexer for a given namespace.
func (s tCPRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.TCPRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.TCPRoute))
	})
	return ret, err
}

// Get retrieves the TCPRoute from the indexer for a given namespace and name.
func (s tCPRouteNamespaceLister) Get(name string) (*v1alpha3.TCPRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("tcproute"), name)
	}
	return obj.(*v1alpha3.TCPRoute), nil
}
