// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package resourcegraphapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resourcegraph/mgmt/2021-03-01/resourcegraph"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	Resources(ctx context.Context, query resourcegraph.QueryRequest) (result resourcegraph.QueryResponse, err error)
}

var _ BaseClientAPI = (*resourcegraph.BaseClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result resourcegraph.OperationListResult, err error)
}

var _ OperationsClientAPI = (*resourcegraph.OperationsClient)(nil)
