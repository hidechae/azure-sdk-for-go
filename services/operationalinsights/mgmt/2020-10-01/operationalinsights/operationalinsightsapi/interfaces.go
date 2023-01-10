// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package operationalinsightsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2020-10-01/operationalinsights"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters operationalinsights.Cluster) (result operationalinsights.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result operationalinsights.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result operationalinsights.Cluster, err error)
	List(ctx context.Context) (result operationalinsights.ClusterListResultPage, err error)
	ListComplete(ctx context.Context) (result operationalinsights.ClusterListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.ClusterListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result operationalinsights.ClusterListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters operationalinsights.ClusterPatch) (result operationalinsights.Cluster, err error)
}

var _ ClustersClientAPI = (*operationalinsights.ClustersClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result operationalinsights.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result operationalinsights.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*operationalinsights.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.Workspace) (result operationalinsights.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, force *bool) (result operationalinsights.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.Workspace, err error)
	List(ctx context.Context) (result operationalinsights.WorkspaceListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.WorkspaceListResult, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.WorkspacePatch) (result operationalinsights.Workspace, err error)
}

var _ WorkspacesClientAPI = (*operationalinsights.WorkspacesClient)(nil)

// DeletedWorkspacesClientAPI contains the set of methods on the DeletedWorkspacesClient type.
type DeletedWorkspacesClientAPI interface {
	List(ctx context.Context) (result operationalinsights.WorkspaceListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result operationalinsights.WorkspaceListResult, err error)
}

var _ DeletedWorkspacesClientAPI = (*operationalinsights.DeletedWorkspacesClient)(nil)
