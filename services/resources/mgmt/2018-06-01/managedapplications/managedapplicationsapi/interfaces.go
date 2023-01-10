// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package managedapplicationsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-06-01/managedapplications"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	ListOperations(ctx context.Context) (result managedapplications.OperationListResultPage, err error)
	ListOperationsComplete(ctx context.Context) (result managedapplications.OperationListResultIterator, err error)
}

var _ BaseClientAPI = (*managedapplications.BaseClient)(nil)

// ApplicationsClientAPI contains the set of methods on the ApplicationsClient type.
type ApplicationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters managedapplications.Application) (result managedapplications.ApplicationsCreateOrUpdateFuture, err error)
	CreateOrUpdateByID(ctx context.Context, applicationID string, parameters managedapplications.Application) (result managedapplications.ApplicationsCreateOrUpdateByIDFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationName string) (result managedapplications.ApplicationsDeleteFuture, err error)
	DeleteByID(ctx context.Context, applicationID string) (result managedapplications.ApplicationsDeleteByIDFuture, err error)
	Get(ctx context.Context, resourceGroupName string, applicationName string) (result managedapplications.Application, err error)
	GetByID(ctx context.Context, applicationID string) (result managedapplications.Application, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result managedapplications.ApplicationListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result managedapplications.ApplicationListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result managedapplications.ApplicationListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result managedapplications.ApplicationListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, applicationName string, parameters *managedapplications.ApplicationPatchable) (result managedapplications.Application, err error)
	UpdateByID(ctx context.Context, applicationID string, parameters *managedapplications.Application) (result managedapplications.Application, err error)
}

var _ ApplicationsClientAPI = (*managedapplications.ApplicationsClient)(nil)

// ApplicationDefinitionsClientAPI contains the set of methods on the ApplicationDefinitionsClient type.
type ApplicationDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters managedapplications.ApplicationDefinition) (result managedapplications.ApplicationDefinitionsCreateOrUpdateFuture, err error)
	CreateOrUpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters managedapplications.ApplicationDefinition) (result managedapplications.ApplicationDefinitionsCreateOrUpdateByIDFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result managedapplications.ApplicationDefinitionsDeleteFuture, err error)
	DeleteByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result managedapplications.ApplicationDefinitionsDeleteByIDFuture, err error)
	Get(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result managedapplications.ApplicationDefinition, err error)
	GetByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string) (result managedapplications.ApplicationDefinition, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result managedapplications.ApplicationDefinitionListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result managedapplications.ApplicationDefinitionListResultIterator, err error)
}

var _ ApplicationDefinitionsClientAPI = (*managedapplications.ApplicationDefinitionsClient)(nil)
