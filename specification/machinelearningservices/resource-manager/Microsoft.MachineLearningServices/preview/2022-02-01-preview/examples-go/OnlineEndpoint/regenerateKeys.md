```go
package armmachinelearning_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineEndpoint/regenerateKeys.json
func ExampleOnlineEndpointsClient_BeginRegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmachinelearning.NewOnlineEndpointsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginRegenerateKeys(ctx,
		"test-rg",
		"my-aml-workspace",
		"testEndpointName",
		armmachinelearning.RegenerateEndpointKeysRequest{
			KeyType:  to.Ptr(armmachinelearning.KeyTypePrimary),
			KeyValue: to.Ptr("string"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearning%2Farmmachinelearning%2Fv2.0.0-beta.1/sdk/resourcemanager/machinelearning/armmachinelearning/README.md) on how to add the SDK to your project and authenticate.
