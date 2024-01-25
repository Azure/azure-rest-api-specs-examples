package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_ListByStreamingJob.json
func ExampleFunctionsClient_NewListByStreamingJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFunctionsClient().NewListByStreamingJobPager("sjrg1637", "sj8653", &armstreamanalytics.FunctionsClientListByStreamingJobOptions{Select: nil})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.FunctionListResult = armstreamanalytics.FunctionListResult{
		// 	Value: []*armstreamanalytics.Function{
		// 		{
		// 			Name: to.Ptr("function588"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg1637/providers/Microsoft.StreamAnalytics/streamingjobs/sj8653/functions/function588"),
		// 			Properties: &armstreamanalytics.ScalarFunctionProperties{
		// 				Type: to.Ptr("Scalar"),
		// 				Etag: to.Ptr("281cbf11-cd50-4a44-b833-cf170ce74748"),
		// 				Properties: &armstreamanalytics.FunctionConfiguration{
		// 					Binding: &armstreamanalytics.AzureMachineLearningStudioFunctionBinding{
		// 						Type: to.Ptr("Microsoft.MachineLearning/WebService"),
		// 						Properties: &armstreamanalytics.AzureMachineLearningStudioFunctionBindingProperties{
		// 							BatchSize: to.Ptr[int32](5000),
		// 							Endpoint: to.Ptr("someAzureMLEndpointURL"),
		// 							Inputs: &armstreamanalytics.AzureMachineLearningStudioInputs{
		// 								Name: to.Ptr("input1"),
		// 								ColumnNames: []*armstreamanalytics.AzureMachineLearningStudioInputColumn{
		// 									{
		// 										Name: to.Ptr("tweet"),
		// 										DataType: to.Ptr("string"),
		// 										MapTo: to.Ptr[int32](0),
		// 								}},
		// 							},
		// 							Outputs: []*armstreamanalytics.AzureMachineLearningStudioOutputColumn{
		// 								{
		// 									Name: to.Ptr("Sentiment"),
		// 									DataType: to.Ptr("string"),
		// 							}},
		// 						},
		// 					},
		// 					Inputs: []*armstreamanalytics.FunctionInput{
		// 						{
		// 							DataType: to.Ptr("nvarchar(max)"),
		// 					}},
		// 					Output: &armstreamanalytics.FunctionOutput{
		// 						DataType: to.Ptr("nvarchar(max)"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("function8197"),
		// 			Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions"),
		// 			ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg1637/providers/Microsoft.StreamAnalytics/streamingjobs/sj8653/functions/function8197"),
		// 			Properties: &armstreamanalytics.ScalarFunctionProperties{
		// 				Type: to.Ptr("Scalar"),
		// 				Etag: to.Ptr("94a512d5-2f59-4e39-b9c8-bca4abd74b7e"),
		// 				Properties: &armstreamanalytics.FunctionConfiguration{
		// 					Binding: &armstreamanalytics.JavaScriptFunctionBinding{
		// 						Type: to.Ptr("Microsoft.StreamAnalytics/JavascriptUdf"),
		// 						Properties: &armstreamanalytics.JavaScriptFunctionBindingProperties{
		// 							Script: to.Ptr("function (a, b) { return a * b; }"),
		// 						},
		// 					},
		// 					Inputs: []*armstreamanalytics.FunctionInput{
		// 						{
		// 							DataType: to.Ptr("Any"),
		// 					}},
		// 					Output: &armstreamanalytics.FunctionOutput{
		// 						DataType: to.Ptr("Any"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
