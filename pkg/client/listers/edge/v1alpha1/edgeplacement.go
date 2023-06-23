//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeStellar Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	edgev1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/edge/v1alpha1"
)

// EdgePlacementClusterLister can list EdgePlacements across all workspaces, or scope down to a EdgePlacementLister for one workspace.
// All objects returned here must be treated as read-only.
type EdgePlacementClusterLister interface {
	// List lists all EdgePlacements in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.EdgePlacement, err error)
	// Cluster returns a lister that can list and get EdgePlacements in one workspace.
	Cluster(clusterName logicalcluster.Name) EdgePlacementLister
	EdgePlacementClusterListerExpansion
}

type edgePlacementClusterLister struct {
	indexer cache.Indexer
}

// NewEdgePlacementClusterLister returns a new EdgePlacementClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewEdgePlacementClusterLister(indexer cache.Indexer) *edgePlacementClusterLister {
	return &edgePlacementClusterLister{indexer: indexer}
}

// List lists all EdgePlacements in the indexer across all workspaces.
func (s *edgePlacementClusterLister) List(selector labels.Selector) (ret []*edgev1alpha1.EdgePlacement, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*edgev1alpha1.EdgePlacement))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get EdgePlacements.
func (s *edgePlacementClusterLister) Cluster(clusterName logicalcluster.Name) EdgePlacementLister {
	return &edgePlacementLister{indexer: s.indexer, clusterName: clusterName}
}

// EdgePlacementLister can list all EdgePlacements, or get one in particular.
// All objects returned here must be treated as read-only.
type EdgePlacementLister interface {
	// List lists all EdgePlacements in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*edgev1alpha1.EdgePlacement, err error)
	// Get retrieves the EdgePlacement from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*edgev1alpha1.EdgePlacement, error)
	EdgePlacementListerExpansion
}

// edgePlacementLister can list all EdgePlacements inside a workspace.
type edgePlacementLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all EdgePlacements in the indexer for a workspace.
func (s *edgePlacementLister) List(selector labels.Selector) (ret []*edgev1alpha1.EdgePlacement, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.EdgePlacement))
	})
	return ret, err
}

// Get retrieves the EdgePlacement from the indexer for a given workspace and name.
func (s *edgePlacementLister) Get(name string) (*edgev1alpha1.EdgePlacement, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("EdgePlacement"), name)
	}
	return obj.(*edgev1alpha1.EdgePlacement), nil
}

// NewEdgePlacementLister returns a new EdgePlacementLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewEdgePlacementLister(indexer cache.Indexer) *edgePlacementScopedLister {
	return &edgePlacementScopedLister{indexer: indexer}
}

// edgePlacementScopedLister can list all EdgePlacements inside a workspace.
type edgePlacementScopedLister struct {
	indexer cache.Indexer
}

// List lists all EdgePlacements in the indexer for a workspace.
func (s *edgePlacementScopedLister) List(selector labels.Selector) (ret []*edgev1alpha1.EdgePlacement, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*edgev1alpha1.EdgePlacement))
	})
	return ret, err
}

// Get retrieves the EdgePlacement from the indexer for a given workspace and name.
func (s *edgePlacementScopedLister) Get(name string) (*edgev1alpha1.EdgePlacement, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(edgev1alpha1.Resource("EdgePlacement"), name)
	}
	return obj.(*edgev1alpha1.EdgePlacement), nil
}
