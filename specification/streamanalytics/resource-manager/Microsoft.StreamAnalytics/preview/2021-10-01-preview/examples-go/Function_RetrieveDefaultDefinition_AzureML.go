package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_RetrieveDefaultDefinition_AzureML.json
func ExampleFunctionsClient_RetrieveDefaultDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFunctionsClient().RetrieveDefaultDefinition(ctx, "sjrg7", "sj9093", "function588", &armstreamanalytics.FunctionsClientRetrieveDefaultDefinitionOptions{FunctionRetrieveDefaultDefinitionParameters: &armstreamanalytics.AzureMachineLearningStudioFunctionRetrieveDefaultDefinitionParameters{
		BindingType: to.Ptr("Microsoft.MachineLearning/WebService"),
		BindingRetrievalProperties: &armstreamanalytics.AzureMachineLearningStudioFunctionBindingRetrievalProperties{
			ExecuteEndpoint: to.Ptr("someAzureMLExecuteEndpointUrl"),
			UdfType:         to.Ptr("Scalar"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Function = armstreamanalytics.Function{
	// 	Name: to.Ptr("function588"),
	// 	Properties: &armstreamanalytics.ScalarFunctionProperties{
	// 		Type: to.Ptr("Scalar"),
	// 		Properties: &armstreamanalytics.FunctionConfiguration{
	// 			Binding: &armstreamanalytics.AzureMachineLearningStudioFunctionBinding{
	// 				Type: to.Ptr("Microsoft.MachineLearning/WebService"),
	// 				Properties: &armstreamanalytics.AzureMachineLearningStudioFunctionBindingProperties{
	// 					BatchSize: to.Ptr[int32](1000),
	// 					Endpoint: to.Ptr("someAzureMLExecuteEndpointUrl"),
	// 					Inputs: &armstreamanalytics.AzureMachineLearningStudioInputs{
	// 						Name: to.Ptr("input1"),
	// 						ColumnNames: []*armstreamanalytics.AzureMachineLearningStudioInputColumn{
	// 							{
	// 								Name: to.Ptr("tweet"),
	// 								DataType: to.Ptr("string"),
	// 								MapTo: to.Ptr[int32](0),
	// 						}},
	// 					},
	// 					Outputs: []*armstreamanalytics.AzureMachineLearningStudioOutputColumn{
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
