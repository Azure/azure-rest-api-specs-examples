```go
package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PrivateEndpointConnectionUpdate.json
func ExamplePrivateEndpointConnectionClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewPrivateEndpointConnectionClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"default-azurebatch-japaneast",
		"sampleacct",
		"testprivateEndpointConnection5.24d6b4b5-e65c-4330-bbe9-3a290d62f8e0",
		armbatch.PrivateEndpointConnection{
			Properties: &armbatch.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armbatch.PrivateLinkServiceConnectionState{
					Description: to.Ptr("Approved by xyz.abc@company.com"),
					Status:      to.Ptr(armbatch.PrivateLinkServiceConnectionStatusApproved),
				},
			},
		},
		&armbatch.PrivateEndpointConnectionClientBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbatch%2Farmbatch%2Fv1.0.0/sdk/resourcemanager/batch/armbatch/README.md) on how to add the SDK to your project and authenticate.
