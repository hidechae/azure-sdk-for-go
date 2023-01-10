// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package accountapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/datalake/analytics/mgmt/2015-10-01-preview/account"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	AddDataLakeStoreAccount(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, parameters account.AddDataLakeStoreParameters) (result autorest.Response, err error)
	AddStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters account.AddStorageAccountParameters) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeAnalyticsAccount) (result account.CreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result account.DeleteFuture, err error)
	DeleteDataLakeStoreAccount(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string) (result autorest.Response, err error)
	DeleteStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeAnalyticsAccount, err error)
	GetDataLakeStoreAccount(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string) (result account.DataLakeStoreAccountInfo, err error)
	GetStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.StorageAccountInfo, err error)
	GetStorageContainer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.BlobContainer, err error)
	List(ctx context.Context, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListResultIterator, err error)
	ListDataLakeStoreAccounts(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListDataLakeStoreResultPage, err error)
	ListDataLakeStoreAccountsComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListDataLakeStoreResultIterator, err error)
	ListSasTokens(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.ListSasTokensResultPage, err error)
	ListSasTokensComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.ListSasTokensResultIterator, err error)
	ListStorageAccounts(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListStorageAccountsResultPage, err error)
	ListStorageAccountsComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, expand string, selectParameter string, orderby string, count *bool, search string, formatParameter string) (result account.DataLakeAnalyticsAccountListStorageAccountsResultIterator, err error)
	ListStorageContainers(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.ListBlobContainersResultPage, err error)
	ListStorageContainersComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.ListBlobContainersResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, name string, parameters account.DataLakeAnalyticsAccount) (result account.UpdateFuture, err error)
	UpdateStorageAccount(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters account.AddStorageAccountParameters) (result autorest.Response, err error)
}

var _ ClientAPI = (*account.Client)(nil)
