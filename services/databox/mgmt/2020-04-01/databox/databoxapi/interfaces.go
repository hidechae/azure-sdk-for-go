// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package databoxapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/databox/mgmt/2020-04-01/databox"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result databox.OperationListPage, err error)
	ListComplete(ctx context.Context) (result databox.OperationListIterator, err error)
}

var _ OperationsClientAPI = (*databox.OperationsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	BookShipmentPickUp(ctx context.Context, resourceGroupName string, jobName string, shipmentPickUpRequest databox.ShipmentPickUpRequest) (result databox.ShipmentPickUpResponse, err error)
	Cancel(ctx context.Context, resourceGroupName string, jobName string, cancellationReason databox.CancellationReason) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, jobName string, jobResource databox.JobResource) (result databox.JobsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, jobName string) (result databox.JobsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string, expand string) (result databox.JobResource, err error)
	List(ctx context.Context, skipToken string) (result databox.JobResourceListPage, err error)
	ListComplete(ctx context.Context, skipToken string) (result databox.JobResourceListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string) (result databox.JobResourceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skipToken string) (result databox.JobResourceListIterator, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, jobName string) (result databox.UnencryptedCredentialsList, err error)
	Update(ctx context.Context, resourceGroupName string, jobName string, jobResourceUpdateParameter databox.JobResourceUpdateParameter, ifMatch string) (result databox.JobsUpdateFuture, err error)
}

var _ JobsClientAPI = (*databox.JobsClient)(nil)

// ServiceClientAPI contains the set of methods on the ServiceClient type.
type ServiceClientAPI interface {
	ListAvailableSkusByResourceGroup(ctx context.Context, resourceGroupName string, location string, availableSkuRequest databox.AvailableSkuRequest) (result databox.AvailableSkusResultPage, err error)
	ListAvailableSkusByResourceGroupComplete(ctx context.Context, resourceGroupName string, location string, availableSkuRequest databox.AvailableSkuRequest) (result databox.AvailableSkusResultIterator, err error)
	RegionConfiguration(ctx context.Context, location string, regionConfigurationRequest databox.RegionConfigurationRequest) (result databox.RegionConfigurationResponse, err error)
	RegionConfigurationByResourceGroup(ctx context.Context, resourceGroupName string, location string, regionConfigurationRequest databox.RegionConfigurationRequest) (result databox.RegionConfigurationResponse, err error)
	ValidateAddressMethod(ctx context.Context, location string, validateAddress databox.ValidateAddress) (result databox.AddressValidationOutput, err error)
	ValidateInputs(ctx context.Context, location string, validationRequest databox.BasicValidationRequest) (result databox.ValidationResponse, err error)
	ValidateInputsByResourceGroup(ctx context.Context, resourceGroupName string, location string, validationRequest databox.BasicValidationRequest) (result databox.ValidationResponse, err error)
}

var _ ServiceClientAPI = (*databox.ServiceClient)(nil)
