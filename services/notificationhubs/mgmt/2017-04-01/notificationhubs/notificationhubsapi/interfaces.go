// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package notificationhubsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/notificationhubs/mgmt/2017-04-01/notificationhubs"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result notificationhubs.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result notificationhubs.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*notificationhubs.OperationsClient)(nil)

// NamespacesClientAPI contains the set of methods on the NamespacesClient type.
type NamespacesClientAPI interface {
	CheckAvailability(ctx context.Context, parameters notificationhubs.CheckAvailabilityParameters) (result notificationhubs.CheckAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, parameters notificationhubs.NamespaceCreateOrUpdateParameters) (result notificationhubs.NamespaceResource, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters notificationhubs.SharedAccessAuthorizationRuleCreateOrUpdateParameters) (result notificationhubs.SharedAccessAuthorizationRuleResource, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result notificationhubs.NamespacesDeleteFuture, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result notificationhubs.NamespaceResource, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result notificationhubs.SharedAccessAuthorizationRuleResource, err error)
	List(ctx context.Context, resourceGroupName string) (result notificationhubs.NamespaceListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string) (result notificationhubs.NamespaceListResultIterator, err error)
	ListAll(ctx context.Context) (result notificationhubs.NamespaceListResultPage, err error)
	ListAllComplete(ctx context.Context) (result notificationhubs.NamespaceListResultIterator, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string) (result notificationhubs.SharedAccessAuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result notificationhubs.SharedAccessAuthorizationRuleListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result notificationhubs.ResourceListKeys, err error)
	Patch(ctx context.Context, resourceGroupName string, namespaceName string, parameters notificationhubs.NamespacePatchParameters) (result notificationhubs.NamespaceResource, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters notificationhubs.PolicykeyResource) (result notificationhubs.ResourceListKeys, err error)
}

var _ NamespacesClientAPI = (*notificationhubs.NamespacesClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckNotificationHubAvailability(ctx context.Context, resourceGroupName string, namespaceName string, parameters notificationhubs.CheckAvailabilityParameters) (result notificationhubs.CheckAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, parameters notificationhubs.CreateOrUpdateParameters) (result notificationhubs.ResourceType, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string, parameters notificationhubs.SharedAccessAuthorizationRuleCreateOrUpdateParameters) (result notificationhubs.SharedAccessAuthorizationRuleResource, err error)
	DebugSend(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, parameters *interface{}) (result notificationhubs.DebugSendResponse, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string) (result autorest.Response, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string) (result notificationhubs.ResourceType, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (result notificationhubs.SharedAccessAuthorizationRuleResource, err error)
	GetPnsCredentials(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string) (result notificationhubs.PnsCredentialsResource, err error)
	List(ctx context.Context, resourceGroupName string, namespaceName string) (result notificationhubs.ListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result notificationhubs.ListResultIterator, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string) (result notificationhubs.SharedAccessAuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string) (result notificationhubs.SharedAccessAuthorizationRuleListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string) (result notificationhubs.ResourceListKeys, err error)
	Patch(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, parameters *notificationhubs.PatchParameters) (result notificationhubs.ResourceType, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, notificationHubName string, authorizationRuleName string, parameters notificationhubs.PolicykeyResource) (result notificationhubs.ResourceListKeys, err error)
}

var _ ClientAPI = (*notificationhubs.Client)(nil)
