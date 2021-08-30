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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openebs/device-localpv/pkg/apis/openebs.io/scp/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StoragePoolLister helps list StoragePools.
// All objects returned here must be treated as read-only.
type StoragePoolLister interface {
	// List lists all StoragePools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StoragePool, err error)
	// StoragePools returns an object that can list and get StoragePools.
	StoragePools(namespace string) StoragePoolNamespaceLister
	StoragePoolListerExpansion
}

// storagePoolLister implements the StoragePoolLister interface.
type storagePoolLister struct {
	indexer cache.Indexer
}

// NewStoragePoolLister returns a new StoragePoolLister.
func NewStoragePoolLister(indexer cache.Indexer) StoragePoolLister {
	return &storagePoolLister{indexer: indexer}
}

// List lists all StoragePools in the indexer.
func (s *storagePoolLister) List(selector labels.Selector) (ret []*v1alpha1.StoragePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragePool))
	})
	return ret, err
}

// StoragePools returns an object that can list and get StoragePools.
func (s *storagePoolLister) StoragePools(namespace string) StoragePoolNamespaceLister {
	return storagePoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StoragePoolNamespaceLister helps list and get StoragePools.
// All objects returned here must be treated as read-only.
type StoragePoolNamespaceLister interface {
	// List lists all StoragePools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StoragePool, err error)
	// Get retrieves the StoragePool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StoragePool, error)
	StoragePoolNamespaceListerExpansion
}

// storagePoolNamespaceLister implements the StoragePoolNamespaceLister
// interface.
type storagePoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StoragePools in the indexer for a given namespace.
func (s storagePoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StoragePool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragePool))
	})
	return ret, err
}

// Get retrieves the StoragePool from the indexer for a given namespace and name.
func (s storagePoolNamespaceLister) Get(name string) (*v1alpha1.StoragePool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagepool"), name)
	}
	return obj.(*v1alpha1.StoragePool), nil
}