package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fa469a1157c33837a46c9bcd524527e94125189a/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2021-10-01-preview/examples/Function_Create_CSharp.json
func ExampleFunctionsClient_CreateOrReplace_createAClrUdfFunction() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstreamanalytics.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFunctionsClient().CreateOrReplace(ctx, "sjrg", "sjName", "function588", armstreamanalytics.Function{
		Properties: &armstreamanalytics.ScalarFunctionProperties{
			Type: to.Ptr("Scalar"),
			Properties: &armstreamanalytics.FunctionConfiguration{
				Binding: &armstreamanalytics.CSharpFunctionBinding{
					Type: to.Ptr("Microsoft.StreamAnalytics/CLRUdf"),
					Properties: &armstreamanalytics.CSharpFunctionBindingProperties{
						Method:     to.Ptr("SquareFunction"),
						Class:      to.Ptr("ASAEdgeUDFDemo.Class1"),
						DllPath:    to.Ptr("ASAEdgeApplication2_CodeBehind"),
						UpdateMode: to.Ptr(armstreamanalytics.UpdateModeStatic),
					},
				},
				Inputs: []*armstreamanalytics.FunctionInput{
					{
						DataType: to.Ptr("nvarchar(max)"),
					}},
				Output: &armstreamanalytics.FunctionOutput{
					DataType: to.Ptr("nvarchar(max)"),
				},
			},
		},
	}, &armstreamanalytics.FunctionsClientCreateOrReplaceOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Function = armstreamanalytics.Function{
	// 	Name: to.Ptr("function588"),
	// 	Type: to.Ptr("Microsoft.StreamAnalytics/streamingjobs/functions"),
	// 	ID: to.Ptr("/subscriptions/56b5e0a9-b645-407d-99b0-c64f86013e3d/resourceGroups/sjrg/providers/Microsoft.StreamAnalytics/streamingjobs/sjName/functions/function588"),
	// 	Properties: &armstreamanalytics.ScalarFunctionProperties{
	// 		Type: to.Ptr("Scalar"),
	// 		Properties: &armstreamanalytics.FunctionConfiguration{
	// 			Binding: &armstreamanalytics.CSharpFunctionBinding{
	// 				Type: to.Ptr("Microsoft.StreamAnalytics/CLRUdf"),
	// 				Properties: &armstreamanalytics.CSharpFunctionBindingProperties{
	// 					Method: to.Ptr("SquareFunction"),
	// 					Class: to.Ptr("ASAEdgeUDFDemo.Class1"),
	// 					DllPath: to.Ptr("ASAEdgeApplication2_CodeBehind"),
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
