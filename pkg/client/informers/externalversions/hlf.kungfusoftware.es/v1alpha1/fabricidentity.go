/*
 * Copyright Kungfusoftware.es. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package v1alpha1

import (
	"context"
	time "time"

	hlfkungfusoftwareesv1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	versioned "github.com/kfsoftware/hlf-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kfsoftware/hlf-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kfsoftware/hlf-operator/pkg/client/listers/hlf.kungfusoftware.es/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FabricIdentityInformer provides access to a shared informer and lister for
// FabricIdentities.
type FabricIdentityInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FabricIdentityLister
}

type fabricIdentityInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFabricIdentityInformer constructs a new informer for FabricIdentity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFabricIdentityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFabricIdentityInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFabricIdentityInformer constructs a new informer for FabricIdentity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFabricIdentityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HlfV1alpha1().FabricIdentities(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HlfV1alpha1().FabricIdentities(namespace).Watch(context.TODO(), options)
			},
		},
		&hlfkungfusoftwareesv1alpha1.FabricIdentity{},
		resyncPeriod,
		indexers,
	)
}

func (f *fabricIdentityInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFabricIdentityInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fabricIdentityInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hlfkungfusoftwareesv1alpha1.FabricIdentity{}, f.defaultInformer)
}

func (f *fabricIdentityInformer) Lister() v1alpha1.FabricIdentityLister {
	return v1alpha1.NewFabricIdentityLister(f.Informer().GetIndexer())
}