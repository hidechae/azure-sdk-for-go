// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package insightsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2020-10-01/insights"
	"github.com/Azure/go-autorest/autorest"
)

// ActivityLogAlertsClientAPI contains the set of methods on the ActivityLogAlertsClient type.
type ActivityLogAlertsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRule insights.ActivityLogAlertResource) (result insights.ActivityLogAlertResource, err error)
	Delete(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result insights.ActivityLogAlertResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result insights.AlertRuleListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result insights.AlertRuleListIterator, err error)
	ListBySubscriptionID(ctx context.Context) (result insights.AlertRuleListPage, err error)
	ListBySubscriptionIDComplete(ctx context.Context) (result insights.AlertRuleListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRulePatch insights.AlertRulePatchObject) (result insights.ActivityLogAlertResource, err error)
}

var _ ActivityLogAlertsClientAPI = (*insights.ActivityLogAlertsClient)(nil)
