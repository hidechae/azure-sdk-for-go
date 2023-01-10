// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package trafficmanagerapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/trafficmanager/mgmt/2017-05-01/trafficmanager"
)

// EndpointsClientAPI contains the set of methods on the EndpointsClient type.
type EndpointsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters trafficmanager.Endpoint) (result trafficmanager.Endpoint, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (result trafficmanager.DeleteOperationResult, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (result trafficmanager.Endpoint, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters trafficmanager.Endpoint) (result trafficmanager.Endpoint, err error)
}

var _ EndpointsClientAPI = (*trafficmanager.EndpointsClient)(nil)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	CheckTrafficManagerRelativeDNSNameAvailability(ctx context.Context, parameters trafficmanager.CheckTrafficManagerRelativeDNSNameAvailabilityParameters) (result trafficmanager.NameAvailability, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, parameters trafficmanager.Profile) (result trafficmanager.Profile, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string) (result trafficmanager.DeleteOperationResult, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string) (result trafficmanager.Profile, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result trafficmanager.ProfileListResult, err error)
	ListBySubscription(ctx context.Context) (result trafficmanager.ProfileListResult, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, parameters trafficmanager.Profile) (result trafficmanager.Profile, err error)
}

var _ ProfilesClientAPI = (*trafficmanager.ProfilesClient)(nil)

// GeographicHierarchiesClientAPI contains the set of methods on the GeographicHierarchiesClient type.
type GeographicHierarchiesClientAPI interface {
	GetDefault(ctx context.Context) (result trafficmanager.GeographicHierarchy, err error)
}

var _ GeographicHierarchiesClientAPI = (*trafficmanager.GeographicHierarchiesClient)(nil)
