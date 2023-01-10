// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deploymentmanager/armdeploymentmanager). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package deploymentmanagerapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/deploymentmanager/mgmt/2019-11-01-preview/deploymentmanager"
	"github.com/Azure/go-autorest/autorest"
)

// ServiceTopologiesClientAPI contains the set of methods on the ServiceTopologiesClient type.
type ServiceTopologiesClientAPI interface {
	CreateOrUpdate(ctx context.Context, serviceTopologyInfo deploymentmanager.ServiceTopologyResource, resourceGroupName string, serviceTopologyName string) (result deploymentmanager.ServiceTopologyResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceTopologyName string) (result deploymentmanager.ServiceTopologyResource, err error)
	List(ctx context.Context, resourceGroupName string) (result deploymentmanager.ListServiceTopologyResource, err error)
}

var _ ServiceTopologiesClientAPI = (*deploymentmanager.ServiceTopologiesClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceInfo deploymentmanager.ServiceResource) (result deploymentmanager.ServiceResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string) (result deploymentmanager.ServiceResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceTopologyName string) (result deploymentmanager.ListServiceResource, err error)
}

var _ ServicesClientAPI = (*deploymentmanager.ServicesClient)(nil)

// ServiceUnitsClientAPI contains the set of methods on the ServiceUnitsClient type.
type ServiceUnitsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string, serviceUnitInfo deploymentmanager.ServiceUnitResource) (result deploymentmanager.ServiceUnitsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string, serviceUnitName string) (result deploymentmanager.ServiceUnitResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceTopologyName string, serviceName string) (result deploymentmanager.ListServiceUnitResource, err error)
}

var _ ServiceUnitsClientAPI = (*deploymentmanager.ServiceUnitsClient)(nil)

// StepsClientAPI contains the set of methods on the StepsClient type.
type StepsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, stepName string, stepInfo *deploymentmanager.StepResource) (result deploymentmanager.StepResource, err error)
	Delete(ctx context.Context, resourceGroupName string, stepName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, stepName string) (result deploymentmanager.StepResource, err error)
	List(ctx context.Context, resourceGroupName string) (result deploymentmanager.ListStepResource, err error)
}

var _ StepsClientAPI = (*deploymentmanager.StepsClient)(nil)

// RolloutsClientAPI contains the set of methods on the RolloutsClient type.
type RolloutsClientAPI interface {
	Cancel(ctx context.Context, resourceGroupName string, rolloutName string) (result deploymentmanager.Rollout, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, rolloutName string, rolloutRequest *deploymentmanager.RolloutRequest) (result deploymentmanager.RolloutsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, rolloutName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, rolloutName string, retryAttempt *int32) (result deploymentmanager.Rollout, err error)
	List(ctx context.Context, resourceGroupName string) (result deploymentmanager.ListRollout, err error)
	Restart(ctx context.Context, resourceGroupName string, rolloutName string, skipSucceeded *bool) (result deploymentmanager.Rollout, err error)
}

var _ RolloutsClientAPI = (*deploymentmanager.RolloutsClient)(nil)

// ArtifactSourcesClientAPI contains the set of methods on the ArtifactSourcesClient type.
type ArtifactSourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, artifactSourceName string, artifactSourceInfo *deploymentmanager.ArtifactSource) (result deploymentmanager.ArtifactSource, err error)
	Delete(ctx context.Context, resourceGroupName string, artifactSourceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, artifactSourceName string) (result deploymentmanager.ArtifactSource, err error)
	List(ctx context.Context, resourceGroupName string) (result deploymentmanager.ListArtifactSource, err error)
}

var _ ArtifactSourcesClientAPI = (*deploymentmanager.ArtifactSourcesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result deploymentmanager.OperationsList, err error)
}

var _ OperationsClientAPI = (*deploymentmanager.OperationsClient)(nil)
