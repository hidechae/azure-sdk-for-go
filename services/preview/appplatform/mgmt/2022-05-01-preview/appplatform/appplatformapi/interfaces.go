// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package appplatformapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/appplatform/mgmt/2022-05-01-preview/appplatform"
	"github.com/Azure/go-autorest/autorest"
)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, availabilityParameters appplatform.NameAvailabilityParameters) (result appplatform.NameAvailability, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, resource appplatform.ServiceResource) (result appplatform.ServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServicesDeleteFuture, err error)
	DisableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	EnableTestEndpoint(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.TestKeys, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServiceResource, err error)
	List(ctx context.Context, resourceGroupName string) (result appplatform.ServiceResourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result appplatform.ServiceResourceListIterator, err error)
	ListBySubscription(ctx context.Context) (result appplatform.ServiceResourceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result appplatform.ServiceResourceListIterator, err error)
	ListTestKeys(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.TestKeys, err error)
	RegenerateTestKey(ctx context.Context, resourceGroupName string, serviceName string, regenerateTestKeyRequest appplatform.RegenerateTestKeyRequestPayload) (result appplatform.TestKeys, err error)
	Start(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServicesStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServicesStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, resource appplatform.ServiceResource) (result appplatform.ServicesUpdateFuture, err error)
}

var _ ServicesClientAPI = (*appplatform.ServicesClient)(nil)

// ConfigServersClientAPI contains the set of methods on the ConfigServersClient type.
type ConfigServersClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ConfigServerResource, err error)
	UpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, configServerResource appplatform.ConfigServerResource) (result appplatform.ConfigServersUpdatePatchFuture, err error)
	UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, configServerResource appplatform.ConfigServerResource) (result appplatform.ConfigServersUpdatePutFuture, err error)
	Validate(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings appplatform.ConfigServerSettings) (result appplatform.ConfigServersValidateFuture, err error)
}

var _ ConfigServersClientAPI = (*appplatform.ConfigServersClient)(nil)

// ConfigurationServicesClientAPI contains the set of methods on the ConfigurationServicesClient type.
type ConfigurationServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, configurationServiceResource appplatform.ConfigurationServiceResource) (result appplatform.ConfigurationServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string) (result appplatform.ConfigurationServicesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string) (result appplatform.ConfigurationServiceResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ConfigurationServiceResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ConfigurationServiceResourceCollectionIterator, err error)
	Validate(ctx context.Context, resourceGroupName string, serviceName string, configurationServiceName string, settings appplatform.ConfigurationServiceSettings) (result appplatform.ConfigurationServicesValidateFuture, err error)
}

var _ ConfigurationServicesClientAPI = (*appplatform.ConfigurationServicesClient)(nil)

// ServiceRegistriesClientAPI contains the set of methods on the ServiceRegistriesClient type.
type ServiceRegistriesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string) (result appplatform.ServiceRegistriesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string) (result appplatform.ServiceRegistriesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, serviceRegistryName string) (result appplatform.ServiceRegistryResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServiceRegistryResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.ServiceRegistryResourceCollectionIterator, err error)
}

var _ ServiceRegistriesClientAPI = (*appplatform.ServiceRegistriesClient)(nil)

// BuildServiceClientAPI contains the set of methods on the BuildServiceClient type.
type BuildServiceClientAPI interface {
	CreateOrUpdateBuild(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, buildName string, buildParameter appplatform.Build) (result appplatform.Build, err error)
	GetBuild(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, buildName string) (result appplatform.Build, err error)
	GetBuildResult(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, buildName string, buildResultName string) (result appplatform.BuildResult, err error)
	GetBuildResultLog(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, buildName string, buildResultName string) (result appplatform.BuildResultLog, err error)
	GetBuildService(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.BuildService, err error)
	GetResourceUploadURL(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.ResourceUploadDefinition, err error)
	GetSupportedBuildpack(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, buildpackName string) (result appplatform.SupportedBuildpackResource, err error)
	GetSupportedStack(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, stackName string) (result appplatform.SupportedStackResource, err error)
	ListBuildResults(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, buildName string) (result appplatform.BuildResultCollectionPage, err error)
	ListBuildResultsComplete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, buildName string) (result appplatform.BuildResultCollectionIterator, err error)
	ListBuilds(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.BuildCollectionPage, err error)
	ListBuildsComplete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.BuildCollectionIterator, err error)
	ListBuildServices(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.BuildServiceCollectionPage, err error)
	ListBuildServicesComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.BuildServiceCollectionIterator, err error)
	ListSupportedBuildpacks(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.SupportedBuildpacksCollection, err error)
	ListSupportedStacks(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.SupportedStacksCollection, err error)
}

var _ BuildServiceClientAPI = (*appplatform.BuildServiceClient)(nil)

// BuildpackBindingClientAPI contains the set of methods on the BuildpackBindingClient type.
type BuildpackBindingClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string, buildpackBinding appplatform.BuildpackBindingResource) (result appplatform.BuildpackBindingCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string) (result appplatform.BuildpackBindingDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, buildpackBindingName string) (result appplatform.BuildpackBindingResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (result appplatform.BuildpackBindingResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (result appplatform.BuildpackBindingResourceCollectionIterator, err error)
}

var _ BuildpackBindingClientAPI = (*appplatform.BuildpackBindingClient)(nil)

// BuildServiceBuilderClientAPI contains the set of methods on the BuildServiceBuilderClient type.
type BuildServiceBuilderClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, builderResource appplatform.BuilderResource) (result appplatform.BuildServiceBuilderCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (result appplatform.BuildServiceBuilderDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string) (result appplatform.BuilderResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.BuilderResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.BuilderResourceCollectionIterator, err error)
}

var _ BuildServiceBuilderClientAPI = (*appplatform.BuildServiceBuilderClient)(nil)

// BuildServiceAgentPoolClientAPI contains the set of methods on the BuildServiceAgentPoolClient type.
type BuildServiceAgentPoolClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string) (result appplatform.BuildServiceAgentPoolResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.BuildServiceAgentPoolResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string) (result appplatform.BuildServiceAgentPoolResourceCollectionIterator, err error)
	UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, agentPoolName string, agentPoolResource appplatform.BuildServiceAgentPoolResource) (result appplatform.BuildServiceAgentPoolUpdatePutFuture, err error)
}

var _ BuildServiceAgentPoolClientAPI = (*appplatform.BuildServiceAgentPoolClient)(nil)

// MonitoringSettingsClientAPI contains the set of methods on the MonitoringSettingsClient type.
type MonitoringSettingsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.MonitoringSettingResource, err error)
	UpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource appplatform.MonitoringSettingResource) (result appplatform.MonitoringSettingsUpdatePatchFuture, err error)
	UpdatePut(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource appplatform.MonitoringSettingResource) (result appplatform.MonitoringSettingsUpdatePutFuture, err error)
}

var _ MonitoringSettingsClientAPI = (*appplatform.MonitoringSettingsClient)(nil)

// AppsClientAPI contains the set of methods on the AppsClient type.
type AppsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource appplatform.AppResource) (result appplatform.AppsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.AppsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, syncStatus string) (result appplatform.AppResource, err error)
	GetResourceUploadURL(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.ResourceUploadDefinition, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.AppResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.AppResourceCollectionIterator, err error)
	SetActiveDeployments(ctx context.Context, resourceGroupName string, serviceName string, appName string, activeDeploymentCollection appplatform.ActiveDeploymentCollection) (result appplatform.AppsSetActiveDeploymentsFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, appResource appplatform.AppResource) (result appplatform.AppsUpdateFuture, err error)
	ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, appName string, validatePayload appplatform.CustomDomainValidatePayload) (result appplatform.CustomDomainValidateResult, err error)
}

var _ AppsClientAPI = (*appplatform.AppsClient)(nil)

// BindingsClientAPI contains the set of methods on the BindingsClient type.
type BindingsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource appplatform.BindingResource) (result appplatform.BindingsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string) (result appplatform.BindingsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string) (result appplatform.BindingResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.BindingResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.BindingResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, bindingName string, bindingResource appplatform.BindingResource) (result appplatform.BindingsUpdateFuture, err error)
}

var _ BindingsClientAPI = (*appplatform.BindingsClient)(nil)

// StoragesClientAPI contains the set of methods on the StoragesClient type.
type StoragesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, storageName string, storageResource appplatform.StorageResource) (result appplatform.StoragesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, storageName string) (result appplatform.StoragesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, storageName string) (result appplatform.StorageResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.StorageResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.StorageResourceCollectionIterator, err error)
}

var _ StoragesClientAPI = (*appplatform.StoragesClient)(nil)

// CertificatesClientAPI contains the set of methods on the CertificatesClient type.
type CertificatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, certificateName string, certificateResource appplatform.CertificateResource) (result appplatform.CertificatesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, certificateName string) (result appplatform.CertificatesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, certificateName string) (result appplatform.CertificateResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.CertificateResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.CertificateResourceCollectionIterator, err error)
}

var _ CertificatesClientAPI = (*appplatform.CertificatesClient)(nil)

// CustomDomainsClientAPI contains the set of methods on the CustomDomainsClient type.
type CustomDomainsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource appplatform.CustomDomainResource) (result appplatform.CustomDomainsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string) (result appplatform.CustomDomainsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string) (result appplatform.CustomDomainResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.CustomDomainResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, appName string) (result appplatform.CustomDomainResourceCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, domainName string, domainResource appplatform.CustomDomainResource) (result appplatform.CustomDomainsUpdateFuture, err error)
}

var _ CustomDomainsClientAPI = (*appplatform.CustomDomainsClient)(nil)

// DeploymentsClientAPI contains the set of methods on the DeploymentsClient type.
type DeploymentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource appplatform.DeploymentResource) (result appplatform.DeploymentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsDeleteFuture, err error)
	GenerateHeapDump(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, diagnosticParameters appplatform.DiagnosticParameters) (result appplatform.DeploymentsGenerateHeapDumpFuture, err error)
	GenerateThreadDump(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, diagnosticParameters appplatform.DiagnosticParameters) (result appplatform.DeploymentsGenerateThreadDumpFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentResource, err error)
	GetLogFileURL(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.LogFileURLResponse, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (result appplatform.DeploymentResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, appName string, version []string) (result appplatform.DeploymentResourceCollectionIterator, err error)
	ListForCluster(ctx context.Context, resourceGroupName string, serviceName string, version []string) (result appplatform.DeploymentResourceCollectionPage, err error)
	ListForClusterComplete(ctx context.Context, resourceGroupName string, serviceName string, version []string) (result appplatform.DeploymentResourceCollectionIterator, err error)
	Restart(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsStartFuture, err error)
	StartJFR(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, diagnosticParameters appplatform.DiagnosticParameters) (result appplatform.DeploymentsStartJFRFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string) (result appplatform.DeploymentsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, appName string, deploymentName string, deploymentResource appplatform.DeploymentResource) (result appplatform.DeploymentsUpdateFuture, err error)
}

var _ DeploymentsClientAPI = (*appplatform.DeploymentsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result appplatform.AvailableOperationsPage, err error)
	ListComplete(ctx context.Context) (result appplatform.AvailableOperationsIterator, err error)
}

var _ OperationsClientAPI = (*appplatform.OperationsClient)(nil)

// RuntimeVersionsClientAPI contains the set of methods on the RuntimeVersionsClient type.
type RuntimeVersionsClientAPI interface {
	ListRuntimeVersions(ctx context.Context) (result appplatform.AvailableRuntimeVersions, err error)
}

var _ RuntimeVersionsClientAPI = (*appplatform.RuntimeVersionsClient)(nil)

// SkusClientAPI contains the set of methods on the SkusClient type.
type SkusClientAPI interface {
	List(ctx context.Context) (result appplatform.ResourceSkuCollectionPage, err error)
	ListComplete(ctx context.Context) (result appplatform.ResourceSkuCollectionIterator, err error)
}

var _ SkusClientAPI = (*appplatform.SkusClient)(nil)

// GatewaysClientAPI contains the set of methods on the GatewaysClient type.
type GatewaysClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, gatewayResource appplatform.GatewayResource) (result appplatform.GatewaysCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result appplatform.GatewaysDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result appplatform.GatewayResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.GatewayResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.GatewayResourceCollectionIterator, err error)
	ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, validatePayload appplatform.CustomDomainValidatePayload) (result appplatform.CustomDomainValidateResult, err error)
}

var _ GatewaysClientAPI = (*appplatform.GatewaysClient)(nil)

// GatewayRouteConfigsClientAPI contains the set of methods on the GatewayRouteConfigsClient type.
type GatewayRouteConfigsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string, gatewayRouteConfigResource appplatform.GatewayRouteConfigResource) (result appplatform.GatewayRouteConfigsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string) (result appplatform.GatewayRouteConfigsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, routeConfigName string) (result appplatform.GatewayRouteConfigResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result appplatform.GatewayRouteConfigResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result appplatform.GatewayRouteConfigResourceCollectionIterator, err error)
}

var _ GatewayRouteConfigsClientAPI = (*appplatform.GatewayRouteConfigsClient)(nil)

// GatewayCustomDomainsClientAPI contains the set of methods on the GatewayCustomDomainsClient type.
type GatewayCustomDomainsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string, gatewayCustomDomainResource appplatform.GatewayCustomDomainResource) (result appplatform.GatewayCustomDomainsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string) (result appplatform.GatewayCustomDomainsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string, domainName string) (result appplatform.GatewayCustomDomainResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result appplatform.GatewayCustomDomainResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, gatewayName string) (result appplatform.GatewayCustomDomainResourceCollectionIterator, err error)
}

var _ GatewayCustomDomainsClientAPI = (*appplatform.GatewayCustomDomainsClient)(nil)

// APIPortalsClientAPI contains the set of methods on the APIPortalsClient type.
type APIPortalsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string, APIPortalResource appplatform.APIPortalResource) (result appplatform.APIPortalsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string) (result appplatform.APIPortalsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string) (result appplatform.APIPortalResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.APIPortalResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string) (result appplatform.APIPortalResourceCollectionIterator, err error)
	ValidateDomain(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string, validatePayload appplatform.CustomDomainValidatePayload) (result appplatform.CustomDomainValidateResult, err error)
}

var _ APIPortalsClientAPI = (*appplatform.APIPortalsClient)(nil)

// APIPortalCustomDomainsClientAPI contains the set of methods on the APIPortalCustomDomainsClient type.
type APIPortalCustomDomainsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string, domainName string, APIPortalCustomDomainResource appplatform.APIPortalCustomDomainResource) (result appplatform.APIPortalCustomDomainsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string, domainName string) (result appplatform.APIPortalCustomDomainsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string, domainName string) (result appplatform.APIPortalCustomDomainResource, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string) (result appplatform.APIPortalCustomDomainResourceCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, APIPortalName string) (result appplatform.APIPortalCustomDomainResourceCollectionIterator, err error)
}

var _ APIPortalCustomDomainsClientAPI = (*appplatform.APIPortalCustomDomainsClient)(nil)
