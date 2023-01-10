// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package streamanalyticsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2016-03-01/streamanalytics"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result streamanalytics.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result streamanalytics.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*streamanalytics.OperationsClient)(nil)

// StreamingJobsClientAPI contains the set of methods on the StreamingJobsClient type.
type StreamingJobsClientAPI interface {
	CreateOrReplace(ctx context.Context, streamingJob streamanalytics.StreamingJob, resourceGroupName string, jobName string, ifMatch string, ifNoneMatch string) (result streamanalytics.StreamingJobsCreateOrReplaceFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, jobName string) (result streamanalytics.StreamingJobsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string, expand string) (result streamanalytics.StreamingJob, err error)
	List(ctx context.Context, expand string) (result streamanalytics.StreamingJobListResultPage, err error)
	ListComplete(ctx context.Context, expand string) (result streamanalytics.StreamingJobListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string) (result streamanalytics.StreamingJobListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, expand string) (result streamanalytics.StreamingJobListResultIterator, err error)
	Start(ctx context.Context, resourceGroupName string, jobName string, startJobParameters *streamanalytics.StartStreamingJobParameters) (result streamanalytics.StreamingJobsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, jobName string) (result streamanalytics.StreamingJobsStopFuture, err error)
	Update(ctx context.Context, streamingJob streamanalytics.StreamingJob, resourceGroupName string, jobName string, ifMatch string) (result streamanalytics.StreamingJob, err error)
}

var _ StreamingJobsClientAPI = (*streamanalytics.StreamingJobsClient)(nil)

// InputsClientAPI contains the set of methods on the InputsClient type.
type InputsClientAPI interface {
	CreateOrReplace(ctx context.Context, input streamanalytics.Input, resourceGroupName string, jobName string, inputName string, ifMatch string, ifNoneMatch string) (result streamanalytics.Input, err error)
	Delete(ctx context.Context, resourceGroupName string, jobName string, inputName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string, inputName string) (result streamanalytics.Input, err error)
	ListByStreamingJob(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result streamanalytics.InputListResultPage, err error)
	ListByStreamingJobComplete(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result streamanalytics.InputListResultIterator, err error)
	Test(ctx context.Context, resourceGroupName string, jobName string, inputName string, input *streamanalytics.Input) (result streamanalytics.InputsTestFuture, err error)
	Update(ctx context.Context, input streamanalytics.Input, resourceGroupName string, jobName string, inputName string, ifMatch string) (result streamanalytics.Input, err error)
}

var _ InputsClientAPI = (*streamanalytics.InputsClient)(nil)

// OutputsClientAPI contains the set of methods on the OutputsClient type.
type OutputsClientAPI interface {
	CreateOrReplace(ctx context.Context, output streamanalytics.Output, resourceGroupName string, jobName string, outputName string, ifMatch string, ifNoneMatch string) (result streamanalytics.Output, err error)
	Delete(ctx context.Context, resourceGroupName string, jobName string, outputName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string, outputName string) (result streamanalytics.Output, err error)
	ListByStreamingJob(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result streamanalytics.OutputListResultPage, err error)
	ListByStreamingJobComplete(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result streamanalytics.OutputListResultIterator, err error)
	Test(ctx context.Context, resourceGroupName string, jobName string, outputName string, output *streamanalytics.Output) (result streamanalytics.OutputsTestFuture, err error)
	Update(ctx context.Context, output streamanalytics.Output, resourceGroupName string, jobName string, outputName string, ifMatch string) (result streamanalytics.Output, err error)
}

var _ OutputsClientAPI = (*streamanalytics.OutputsClient)(nil)

// TransformationsClientAPI contains the set of methods on the TransformationsClient type.
type TransformationsClientAPI interface {
	CreateOrReplace(ctx context.Context, transformation streamanalytics.Transformation, resourceGroupName string, jobName string, transformationName string, ifMatch string, ifNoneMatch string) (result streamanalytics.Transformation, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string, transformationName string) (result streamanalytics.Transformation, err error)
	Update(ctx context.Context, transformation streamanalytics.Transformation, resourceGroupName string, jobName string, transformationName string, ifMatch string) (result streamanalytics.Transformation, err error)
}

var _ TransformationsClientAPI = (*streamanalytics.TransformationsClient)(nil)

// FunctionsClientAPI contains the set of methods on the FunctionsClient type.
type FunctionsClientAPI interface {
	CreateOrReplace(ctx context.Context, function streamanalytics.Function, resourceGroupName string, jobName string, functionName string, ifMatch string, ifNoneMatch string) (result streamanalytics.Function, err error)
	Delete(ctx context.Context, resourceGroupName string, jobName string, functionName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string, functionName string) (result streamanalytics.Function, err error)
	ListByStreamingJob(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result streamanalytics.FunctionListResultPage, err error)
	ListByStreamingJobComplete(ctx context.Context, resourceGroupName string, jobName string, selectParameter string) (result streamanalytics.FunctionListResultIterator, err error)
	RetrieveDefaultDefinition(ctx context.Context, resourceGroupName string, jobName string, functionName string, functionRetrieveDefaultDefinitionParameters *streamanalytics.BasicFunctionRetrieveDefaultDefinitionParameters) (result streamanalytics.Function, err error)
	Test(ctx context.Context, resourceGroupName string, jobName string, functionName string, function *streamanalytics.Function) (result streamanalytics.FunctionsTestFuture, err error)
	Update(ctx context.Context, function streamanalytics.Function, resourceGroupName string, jobName string, functionName string, ifMatch string) (result streamanalytics.Function, err error)
}

var _ FunctionsClientAPI = (*streamanalytics.FunctionsClient)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	ListQuotas(ctx context.Context, location string) (result streamanalytics.SubscriptionQuotasListResult, err error)
}

var _ SubscriptionsClientAPI = (*streamanalytics.SubscriptionsClient)(nil)
