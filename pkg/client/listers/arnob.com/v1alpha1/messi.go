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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/Arnobkumarsaha/custom-controller/pkg/apis/arnob.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MessiLister helps list Messis.
// All objects returned here must be treated as read-only.
type MessiLister interface {
	// List lists all Messis in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Messi, err error)
	// Messis returns an object that can list and get Messis.
	Messis(namespace string) MessiNamespaceLister
	MessiListerExpansion
}

// messiLister implements the MessiLister interface.
type messiLister struct {
	indexer cache.Indexer
}

// NewMessiLister returns a new MessiLister.
func NewMessiLister(indexer cache.Indexer) MessiLister {
	return &messiLister{indexer: indexer}
}

// List lists all Messis in the indexer.
func (s *messiLister) List(selector labels.Selector) (ret []*v1alpha1.Messi, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Messi))
	})
	return ret, err
}

// Messis returns an object that can list and get Messis.
func (s *messiLister) Messis(namespace string) MessiNamespaceLister {
	return messiNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MessiNamespaceLister helps list and get Messis.
// All objects returned here must be treated as read-only.
type MessiNamespaceLister interface {
	// List lists all Messis in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Messi, err error)
	// Get retrieves the Messi from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Messi, error)
	MessiNamespaceListerExpansion
}

// messiNamespaceLister implements the MessiNamespaceLister
// interface.
type messiNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Messis in the indexer for a given namespace.
func (s messiNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Messi, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Messi))
	})
	return ret, err
}

// Get retrieves the Messi from the indexer for a given namespace and name.
func (s messiNamespaceLister) Get(name string) (*v1alpha1.Messi, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("messi"), name)
	}
	return obj.(*v1alpha1.Messi), nil
}
