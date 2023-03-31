package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Update_AzureML.json
func ExampleFunctionsClient_Update_updateAnAzureMlFunction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFunctionsClient().Update(ctx, "sjrg7", "sj9093", "function588", armstreamanalytics.Function{
		Properties: &armstreamanalytics.ScalarFunctionProperties{
			Type: to.Ptr("Scalar"),
			Properties: &armstreamanalytics.FunctionConfiguration{
				Binding: &armstreamanalytics.AzureMachineLearningWebServiceFunctionBinding{
					Type: to.Ptr("Microsoft.MachineLearning/WebService"),
					Properties: &armstreamanalytics.AzureMachineLearningWebServiceFunctionBindingProperties{
						BatchSize: to.Ptr[int32](5000),
					},
				},
			},
		},
	}, &armstreamanalytics.FunctionsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Function = armstreamanalytics.Function{
	// 	Name: to.Ptr("function588"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg7/providers/Microsoft.StreamAnalytics/streamingjobs/sj9093/functions/function588"),
	// 	Properties: &armstreamanalytics.ScalarFunctionProperties{
	// 		Type: to.Ptr("Scalar"),
	// 		Properties: &armstreamanalytics.FunctionConfiguration{
	// 			Binding: &armstreamanalytics.AzureMachineLearningWebServiceFunctionBinding{
	// 				Type: to.Ptr("Microsoft.MachineLearning/WebService"),
	// 				Properties: &armstreamanalytics.AzureMachineLearningWebServiceFunctionBindingProperties{
	// 					BatchSize: to.Ptr[int32](5000),
	// 					Endpoint: to.Ptr("someAzureMLEndpointURL"),
	// 					Inputs: &armstreamanalytics.AzureMachineLearningWebServiceInputs{
	// 						Name: to.Ptr("input1"),
	// 						ColumnNames: []*armstreamanalytics.AzureMachineLearningWebServiceInputColumn{
	// 							{
	// 								Name: to.Ptr("tweet"),
	// 								DataType: to.Ptr("string"),
	// 								MapTo: to.Ptr[int32](0),
	// 						}},
	// 					},
	// 					Outputs: []*armstreamanalytics.AzureMachineLearningWebServiceOutputColumn{
	// 						{
	// 							Name: to.Ptr("Sentiment"),
	// 							DataType: to.Ptr("string"),
	// 					}},
	// 				},
	// 			},
	// 			Inputs: []*armstreamanalytics.FunctionInput{
	// 				{
	// 					DataType: to.Ptr("nvarchar(max)"),
	// 			}},
	// 			Output: &armstreamanalytics.FunctionOutput{
	// 				DataType: to.Ptr("nvarchar(max)"),
	// 			},
	// 		},
	// 	},
	// }
}
