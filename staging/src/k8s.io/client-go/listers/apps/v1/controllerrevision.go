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

package v1

import (
	appsv1 "k8s.io/api/apps/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// ControllerRevisionLister helps list ControllerRevisions.
// All objects returned here must be treated as read-only.
type ControllerRevisionLister interface {
	// List lists all ControllerRevisions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1.ControllerRevision, err error)
	// ControllerRevisions returns an object that can list and get ControllerRevisions.
	ControllerRevisions(namespace string) ControllerRevisionNamespaceLister
	ControllerRevisionListerExpansion
}

// controllerRevisionLister implements the ControllerRevisionLister interface.
type controllerRevisionLister struct {
	listers.ResourceIndexer[*appsv1.ControllerRevision]
}

// NewControllerRevisionLister returns a new ControllerRevisionLister.
func NewControllerRevisionLister(indexer cache.Indexer) ControllerRevisionLister {
	return &controllerRevisionLister{listers.New[*appsv1.ControllerRevision](indexer, appsv1.Resource("controllerrevision"))}
}

// ControllerRevisions returns an object that can list and get ControllerRevisions.
func (s *controllerRevisionLister) ControllerRevisions(namespace string) ControllerRevisionNamespaceLister {
	return controllerRevisionNamespaceLister{listers.NewNamespaced[*appsv1.ControllerRevision](s.ResourceIndexer, namespace)}
}

// ControllerRevisionNamespaceLister helps list and get ControllerRevisions.
// All objects returned here must be treated as read-only.
type ControllerRevisionNamespaceLister interface {
	// List lists all ControllerRevisions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1.ControllerRevision, err error)
	// Get retrieves the ControllerRevision from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*appsv1.ControllerRevision, error)
	ControllerRevisionNamespaceListerExpansion
}

// controllerRevisionNamespaceLister implements the ControllerRevisionNamespaceLister
// interface.
type controllerRevisionNamespaceLister struct {
	listers.ResourceIndexer[*appsv1.ControllerRevision]
}
