// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcemover/armresourcemover). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package resourcemoverapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resourcemover/mgmt/2021-01-01/resourcemover"
)

// MoveCollectionsClientAPI contains the set of methods on the MoveCollectionsClient type.
type MoveCollectionsClientAPI interface {
	BulkRemove(ctx context.Context, resourceGroupName string, moveCollectionName string, body *resourcemover.BulkRemoveRequest) (result resourcemover.MoveCollectionsBulkRemoveFuture, err error)
	Commit(ctx context.Context, resourceGroupName string, moveCollectionName string, body *resourcemover.CommitRequest) (result resourcemover.MoveCollectionsCommitFuture, err error)
	Create(ctx context.Context, resourceGroupName string, moveCollectionName string, body *resourcemover.MoveCollection) (result resourcemover.MoveCollection, err error)
	Delete(ctx context.Context, resourceGroupName string, moveCollectionName string) (result resourcemover.MoveCollectionsDeleteFuture, err error)
	Discard(ctx context.Context, resourceGroupName string, moveCollectionName string, body *resourcemover.DiscardRequest) (result resourcemover.MoveCollectionsDiscardFuture, err error)
	Get(ctx context.Context, resourceGroupName string, moveCollectionName string) (result resourcemover.MoveCollection, err error)
	InitiateMove(ctx context.Context, resourceGroupName string, moveCollectionName string, body *resourcemover.ResourceMoveRequestType) (result resourcemover.MoveCollectionsInitiateMoveFuture, err error)
	ListMoveCollectionsByResourceGroup(ctx context.Context, resourceGroupName string) (result resourcemover.MoveCollectionResultListPage, err error)
	ListMoveCollectionsByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result resourcemover.MoveCollectionResultListIterator, err error)
	ListMoveCollectionsBySubscription(ctx context.Context) (result resourcemover.MoveCollectionResultListPage, err error)
	ListMoveCollectionsBySubscriptionComplete(ctx context.Context) (result resourcemover.MoveCollectionResultListIterator, err error)
	ListRequiredFor(ctx context.Context, resourceGroupName string, moveCollectionName string, sourceID string) (result resourcemover.RequiredForResourcesCollection, err error)
	Prepare(ctx context.Context, resourceGroupName string, moveCollectionName string, body *resourcemover.PrepareRequest) (result resourcemover.MoveCollectionsPrepareFuture, err error)
	ResolveDependencies(ctx context.Context, resourceGroupName string, moveCollectionName string) (result resourcemover.MoveCollectionsResolveDependenciesFuture, err error)
	Update(ctx context.Context, resourceGroupName string, moveCollectionName string, body *resourcemover.UpdateMoveCollectionRequest) (result resourcemover.MoveCollection, err error)
}

var _ MoveCollectionsClientAPI = (*resourcemover.MoveCollectionsClient)(nil)

// MoveResourcesClientAPI contains the set of methods on the MoveResourcesClient type.
type MoveResourcesClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, body *resourcemover.MoveResource) (result resourcemover.MoveResourcesCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (result resourcemover.MoveResourcesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string) (result resourcemover.MoveResource, err error)
	List(ctx context.Context, resourceGroupName string, moveCollectionName string, filter string) (result resourcemover.MoveResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, moveCollectionName string, filter string) (result resourcemover.MoveResourceCollectionIterator, err error)
}

var _ MoveResourcesClientAPI = (*resourcemover.MoveResourcesClient)(nil)

// UnresolvedDependenciesClientAPI contains the set of methods on the UnresolvedDependenciesClient type.
type UnresolvedDependenciesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, moveCollectionName string, dependencyLevel resourcemover.DependencyLevel, orderby string, filter string) (result resourcemover.UnresolvedDependencyCollectionPage, err error)
	GetComplete(ctx context.Context, resourceGroupName string, moveCollectionName string, dependencyLevel resourcemover.DependencyLevel, orderby string, filter string) (result resourcemover.UnresolvedDependencyCollectionIterator, err error)
}

var _ UnresolvedDependenciesClientAPI = (*resourcemover.UnresolvedDependenciesClient)(nil)

// OperationsDiscoveryClientAPI contains the set of methods on the OperationsDiscoveryClient type.
type OperationsDiscoveryClientAPI interface {
	Get(ctx context.Context) (result resourcemover.OperationsDiscoveryCollection, err error)
}

var _ OperationsDiscoveryClientAPI = (*resourcemover.OperationsDiscoveryClient)(nil)
