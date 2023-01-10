// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package frontdoorapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2019-10-01/frontdoor"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	CheckFrontDoorNameAvailability(ctx context.Context, checkFrontDoorNameAvailabilityInput frontdoor.CheckNameAvailabilityInput) (result frontdoor.CheckNameAvailabilityOutput, err error)
	CheckFrontDoorNameAvailabilityWithSubscription(ctx context.Context, checkFrontDoorNameAvailabilityInput frontdoor.CheckNameAvailabilityInput) (result frontdoor.CheckNameAvailabilityOutput, err error)
}

var _ BaseClientAPI = (*frontdoor.BaseClient)(nil)

// FrontDoorsClientAPI contains the set of methods on the FrontDoorsClient type.
type FrontDoorsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, frontDoorName string, frontDoorParameters frontdoor.FrontDoor) (result frontdoor.FrontDoorsCreateOrUpdateFutureType, err error)
	Delete(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontDoorsDeleteFutureType, err error)
	Get(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontDoor, err error)
	List(ctx context.Context) (result frontdoor.ListResultPage, err error)
	ListComplete(ctx context.Context) (result frontdoor.ListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result frontdoor.ListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result frontdoor.ListResultIterator, err error)
	ValidateCustomDomain(ctx context.Context, resourceGroupName string, frontDoorName string, customDomainProperties frontdoor.ValidateCustomDomainInput) (result frontdoor.ValidateCustomDomainOutput, err error)
}

var _ FrontDoorsClientAPI = (*frontdoor.FrontDoorsClient)(nil)

// FrontendEndpointsClientAPI contains the set of methods on the FrontendEndpointsClient type.
type FrontendEndpointsClientAPI interface {
	DisableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result frontdoor.FrontendEndpointsDisableHTTPSFuture, err error)
	EnableHTTPS(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string, customHTTPSConfiguration frontdoor.CustomHTTPSConfiguration) (result frontdoor.FrontendEndpointsEnableHTTPSFuture, err error)
	Get(ctx context.Context, resourceGroupName string, frontDoorName string, frontendEndpointName string) (result frontdoor.FrontendEndpoint, err error)
	ListByFrontDoor(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontendEndpointsListResultPage, err error)
	ListByFrontDoorComplete(ctx context.Context, resourceGroupName string, frontDoorName string) (result frontdoor.FrontendEndpointsListResultIterator, err error)
}

var _ FrontendEndpointsClientAPI = (*frontdoor.FrontendEndpointsClient)(nil)

// EndpointsClientAPI contains the set of methods on the EndpointsClient type.
type EndpointsClientAPI interface {
	PurgeContent(ctx context.Context, resourceGroupName string, frontDoorName string, contentFilePaths frontdoor.PurgeParameters) (result frontdoor.EndpointsPurgeContentFuture, err error)
}

var _ EndpointsClientAPI = (*frontdoor.EndpointsClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters frontdoor.WebApplicationFirewallPolicy) (result frontdoor.PoliciesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, policyName string) (result frontdoor.PoliciesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, policyName string) (result frontdoor.WebApplicationFirewallPolicy, err error)
	List(ctx context.Context, resourceGroupName string) (result frontdoor.WebApplicationFirewallPolicyListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result frontdoor.WebApplicationFirewallPolicyListIterator, err error)
}

var _ PoliciesClientAPI = (*frontdoor.PoliciesClient)(nil)

// ManagedRuleSetsClientAPI contains the set of methods on the ManagedRuleSetsClient type.
type ManagedRuleSetsClientAPI interface {
	List(ctx context.Context) (result frontdoor.ManagedRuleSetDefinitionListPage, err error)
	ListComplete(ctx context.Context) (result frontdoor.ManagedRuleSetDefinitionListIterator, err error)
}

var _ ManagedRuleSetsClientAPI = (*frontdoor.ManagedRuleSetsClient)(nil)
