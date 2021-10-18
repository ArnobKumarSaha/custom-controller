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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	arnobcomv1alpha1 "github.com/Arnobkumarsaha/messi/pkg/apis/arnob.com/v1alpha1"
	versioned "github.com/Arnobkumarsaha/messi/pkg/client/clientset/versioned"
	internalinterfaces "github.com/Arnobkumarsaha/messi/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Arnobkumarsaha/messi/pkg/client/listers/arnob.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MessiInformer provides access to a shared informer and lister for
// Messis.
type MessiInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MessiLister
}

type messiInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMessiInformer constructs a new informer for Messi type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMessiInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMessiInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMessiInformer constructs a new informer for Messi type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMessiInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArnobV1alpha1().Messis(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArnobV1alpha1().Messis(namespace).Watch(context.TODO(), options)
			},
		},
		&arnobcomv1alpha1.Messi{},
		resyncPeriod,
		indexers,
	)
}

func (f *messiInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMessiInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *messiInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&arnobcomv1alpha1.Messi{}, f.defaultInformer)
}

func (f *messiInformer) Lister() v1alpha1.MessiLister {
	return v1alpha1.NewMessiLister(f.Informer().GetIndexer())
}
