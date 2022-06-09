```go
package armstreamanalytics_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/streamanalytics/armstreamanalytics"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Function_RetrieveDefaultDefinition_AzureML.json
func ExampleFunctionsClient_RetrieveDefaultDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstreamanalytics.NewFunctionsClient("56b5e0a9-b645-407d-99b0-c64f86013e3d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.RetrieveDefaultDefinition(ctx,
		"sjrg7",
		"sj9093",
		"function588",
		&armstreamanalytics.FunctionsClientRetrieveDefaultDefinitionOptions{FunctionRetrieveDefaultDefinitionParameters: &armstreamanalytics.AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters{
			BindingType: to.Ptr("Microsoft.MachineLearning/WebService"),
			BindingRetrievalProperties: &armstreamanalytics.AzureMachineLearningWebServiceFunctionBindingRetrievalProperties{
				ExecuteEndpoint: to.Ptr("someAzureMLExecuteEndpointUrl"),
				UdfType:         to.Ptr("Scalar"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstreamanalytics%2Farmstreamanalytics%2Fv1.0.0/sdk/resourcemanager/streamanalytics/armstreamanalytics/README.md) on how to add the SDK to your project and authenticate.
