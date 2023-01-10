// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiembedded/armpowerbiembedded](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiembedded/armpowerbiembedded). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package powerbiembeddedapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/powerbiembedded/mgmt/2016-01-29/powerbiembedded"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	GetAvailableOperations(ctx context.Context) (result powerbiembedded.OperationList, err error)
}

var _ BaseClientAPI = (*powerbiembedded.BaseClient)(nil)

// WorkspaceCollectionsClientAPI contains the set of methods on the WorkspaceCollectionsClient type.
type WorkspaceCollectionsClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, body powerbiembedded.CheckNameRequest) (result powerbiembedded.CheckNameResponse, err error)
	Create(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body powerbiembedded.CreateWorkspaceCollectionRequest) (result powerbiembedded.WorkspaceCollection, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceCollectionsDeleteFuture, err error)
	GetAccessKeys(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceCollectionAccessKeys, err error)
	GetByName(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceCollection, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result powerbiembedded.WorkspaceCollectionList, err error)
	ListBySubscription(ctx context.Context) (result powerbiembedded.WorkspaceCollectionList, err error)
	Migrate(ctx context.Context, resourceGroupName string, body powerbiembedded.MigrateWorkspaceCollectionRequest) (result autorest.Response, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body powerbiembedded.WorkspaceCollectionAccessKey) (result powerbiembedded.WorkspaceCollectionAccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceCollectionName string, body powerbiembedded.UpdateWorkspaceCollectionRequest) (result powerbiembedded.WorkspaceCollection, err error)
}

var _ WorkspaceCollectionsClientAPI = (*powerbiembedded.WorkspaceCollectionsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, workspaceCollectionName string) (result powerbiembedded.WorkspaceList, err error)
}

var _ WorkspacesClientAPI = (*powerbiembedded.WorkspacesClient)(nil)
