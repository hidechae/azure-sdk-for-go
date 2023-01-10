// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package desktopvirtualizationapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/desktopvirtualization/mgmt/2019-12-10-preview/desktopvirtualization"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result desktopvirtualization.ResourceProviderOperationList, err error)
}

var _ OperationsClientAPI = (*desktopvirtualization.OperationsClient)(nil)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspace desktopvirtualization.Workspace) (result desktopvirtualization.Workspace, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result desktopvirtualization.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result desktopvirtualization.WorkspaceListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result desktopvirtualization.WorkspaceListIterator, err error)
	ListBySubscription(ctx context.Context) (result desktopvirtualization.WorkspaceListPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result desktopvirtualization.WorkspaceListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, workspaceName string, workspace *desktopvirtualization.WorkspacePatch) (result desktopvirtualization.Workspace, err error)
}

var _ WorkspacesClientAPI = (*desktopvirtualization.WorkspacesClient)(nil)

// ApplicationGroupsClientAPI contains the set of methods on the ApplicationGroupsClient type.
type ApplicationGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup desktopvirtualization.ApplicationGroup) (result desktopvirtualization.ApplicationGroup, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, applicationGroupName string) (result desktopvirtualization.ApplicationGroup, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result desktopvirtualization.ApplicationGroupListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result desktopvirtualization.ApplicationGroupListIterator, err error)
	ListBySubscription(ctx context.Context, filter string) (result desktopvirtualization.ApplicationGroupListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, filter string) (result desktopvirtualization.ApplicationGroupListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationGroup *desktopvirtualization.ApplicationGroupPatch) (result desktopvirtualization.ApplicationGroup, err error)
}

var _ ApplicationGroupsClientAPI = (*desktopvirtualization.ApplicationGroupsClient)(nil)

// StartMenuItemsClientAPI contains the set of methods on the StartMenuItemsClient type.
type StartMenuItemsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, applicationGroupName string) (result desktopvirtualization.StartMenuItemListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, applicationGroupName string) (result desktopvirtualization.StartMenuItemListIterator, err error)
}

var _ StartMenuItemsClientAPI = (*desktopvirtualization.StartMenuItemsClient)(nil)

// ApplicationsClientAPI contains the set of methods on the ApplicationsClient type.
type ApplicationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationName string, application desktopvirtualization.Application) (result desktopvirtualization.Application, err error)
	Delete(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationName string) (result desktopvirtualization.Application, err error)
	List(ctx context.Context, resourceGroupName string, applicationGroupName string) (result desktopvirtualization.ApplicationListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, applicationGroupName string) (result desktopvirtualization.ApplicationListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, applicationGroupName string, applicationName string, application *desktopvirtualization.ApplicationPatch) (result desktopvirtualization.Application, err error)
}

var _ ApplicationsClientAPI = (*desktopvirtualization.ApplicationsClient)(nil)

// DesktopsClientAPI contains the set of methods on the DesktopsClient type.
type DesktopsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string) (result desktopvirtualization.Desktop, err error)
	List(ctx context.Context, resourceGroupName string, applicationGroupName string) (result desktopvirtualization.DesktopList, err error)
	Update(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string, desktop *desktopvirtualization.DesktopPatch) (result desktopvirtualization.Desktop, err error)
}

var _ DesktopsClientAPI = (*desktopvirtualization.DesktopsClient)(nil)

// HostPoolsClientAPI contains the set of methods on the HostPoolsClient type.
type HostPoolsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, hostPoolName string, hostPool desktopvirtualization.HostPool) (result desktopvirtualization.HostPool, err error)
	Delete(ctx context.Context, resourceGroupName string, hostPoolName string, force *bool) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, hostPoolName string) (result desktopvirtualization.HostPool, err error)
	List(ctx context.Context) (result desktopvirtualization.HostPoolListPage, err error)
	ListComplete(ctx context.Context) (result desktopvirtualization.HostPoolListIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result desktopvirtualization.HostPoolListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result desktopvirtualization.HostPoolListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, hostPoolName string, hostPool *desktopvirtualization.HostPoolPatch) (result desktopvirtualization.HostPool, err error)
}

var _ HostPoolsClientAPI = (*desktopvirtualization.HostPoolsClient)(nil)

// UserSessionsClientAPI contains the set of methods on the UserSessionsClient type.
type UserSessionsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, force *bool) (result autorest.Response, err error)
	Disconnect(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string) (result desktopvirtualization.UserSession, err error)
	List(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (result desktopvirtualization.UserSessionListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (result desktopvirtualization.UserSessionListIterator, err error)
	ListByHostPool(ctx context.Context, resourceGroupName string, hostPoolName string, filter string) (result desktopvirtualization.UserSessionListPage, err error)
	ListByHostPoolComplete(ctx context.Context, resourceGroupName string, hostPoolName string, filter string) (result desktopvirtualization.UserSessionListIterator, err error)
	SendMessageMethod(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, userSessionID string, sendMessage *desktopvirtualization.SendMessage) (result autorest.Response, err error)
}

var _ UserSessionsClientAPI = (*desktopvirtualization.UserSessionsClient)(nil)

// SessionHostsClientAPI contains the set of methods on the SessionHostsClient type.
type SessionHostsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, force *bool) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string) (result desktopvirtualization.SessionHost, err error)
	List(ctx context.Context, resourceGroupName string, hostPoolName string) (result desktopvirtualization.SessionHostListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, hostPoolName string) (result desktopvirtualization.SessionHostListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, hostPoolName string, sessionHostName string, sessionHost *desktopvirtualization.SessionHostPatch) (result desktopvirtualization.SessionHost, err error)
}

var _ SessionHostsClientAPI = (*desktopvirtualization.SessionHostsClient)(nil)
