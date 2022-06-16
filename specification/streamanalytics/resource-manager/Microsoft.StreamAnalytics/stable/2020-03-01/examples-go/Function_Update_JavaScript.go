package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_Update_JavaScript.json
func ExampleFunctionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"sjrg1637",
		"sj8653",
		"function8197",
		armstreamanalytics.Function{
			Properties: &armstreamanalytics.ScalarFunctionProperties{
				Type: to.Ptr("Scalar"),
				Properties: &armstreamanalytics.FunctionConfiguration{
					Binding: &armstreamanalytics.JavaScriptFunctionBinding{
						Type: to.Ptr("Microsoft.StreamAnalytics/JavascriptUdf"),
						Properties: &armstreamanalytics.JavaScriptFunctionBindingProperties{
							Script: to.Ptr("function (a, b) { return a * b; }"),
						},
					},
				},
			},
		},
		&armstreamanalytics.FunctionsClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
