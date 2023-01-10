// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package operationalinsightsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2015-03-20/operationalinsights"
	"github.com/Azure/go-autorest/autorest"
)

// StorageInsightsClientAPI contains the set of methods on the StorageInsightsClient type.
type StorageInsightsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, parameters operationalinsights.StorageInsight) (result operationalinsights.StorageInsight, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string) (result operationalinsights.StorageInsight, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.StorageInsightListResultPage, err error)
	ListByWorkspaceComplete(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.StorageInsightListResultIterator, err error)
}

var _ StorageInsightsClientAPI = (*operationalinsights.StorageInsightsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	AvailableServiceTiers(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.ListAvailableServiceTier, err error)
	DeleteGateways(ctx context.Context, resourceGroupName string, workspaceName string, gatewayID string) (result autorest.Response, err error)
	GetPurgeStatus(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.WorkspacePurgeStatusResponse, err error)
	GetSchema(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SearchGetSchemaResponse, err error)
	ListKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SharedKeys, err error)
	ListLinkTargets(ctx context.Context) (result operationalinsights.ListLinkTarget, err error)
	Purge(ctx context.Context, resourceGroupName string, workspaceName string, body operationalinsights.WorkspacePurgeBody) (result operationalinsights.WorkspacePurgeResponse, err error)
	RegenerateSharedKeys(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SharedKeys, err error)
}

var _ WorkspacesClientAPI = (*operationalinsights.WorkspacesClient)(nil)

// SavedSearchesClientAPI contains the set of methods on the SavedSearchesClient type.
type SavedSearchesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters operationalinsights.SavedSearch) (result operationalinsights.SavedSearch, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string) (result operationalinsights.SavedSearch, err error)
	ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string) (result operationalinsights.SavedSearchesListResult, err error)
}

var _ SavedSearchesClientAPI = (*operationalinsights.SavedSearchesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result operationalinsights.OperationListResult, err error)
}

var _ OperationsClientAPI = (*operationalinsights.OperationsClient)(nil)
