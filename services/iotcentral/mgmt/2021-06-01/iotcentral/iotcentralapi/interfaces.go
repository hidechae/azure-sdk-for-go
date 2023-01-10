// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package iotcentralapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/iotcentral/mgmt/2021-06-01/iotcentral"
)

// AppsClientAPI contains the set of methods on the AppsClient type.
type AppsClientAPI interface {
	CheckNameAvailability(ctx context.Context, operationInputs iotcentral.OperationInputs) (result iotcentral.AppAvailabilityInfo, err error)
	CheckSubdomainAvailability(ctx context.Context, operationInputs iotcentral.OperationInputs) (result iotcentral.AppAvailabilityInfo, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, app iotcentral.App) (result iotcentral.AppsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceName string) (result iotcentral.AppsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, resourceName string) (result iotcentral.App, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result iotcentral.AppListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result iotcentral.AppListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result iotcentral.AppListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result iotcentral.AppListResultIterator, err error)
	ListTemplates(ctx context.Context) (result iotcentral.AppTemplatesResultPage, err error)
	ListTemplatesComplete(ctx context.Context) (result iotcentral.AppTemplatesResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, resourceName string, appPatch iotcentral.AppPatch) (result iotcentral.AppsUpdateFuture, err error)
}

var _ AppsClientAPI = (*iotcentral.AppsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result iotcentral.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result iotcentral.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*iotcentral.OperationsClient)(nil)
