```go
package armbatch_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/PrivateEndpointConnectionUpdate.json
func ExamplePrivateEndpointConnectionClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewPrivateEndpointConnectionClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-endpoint-connection-name>",
		armbatch.PrivateEndpointConnection{
			Properties: &armbatch.PrivateEndpointConnectionProperties{
				PrivateLinkServiceConnectionState: &armbatch.PrivateLinkServiceConnectionState{
					Description: to.StringPtr("<description>"),
					Status:      armbatch.PrivateLinkServiceConnectionStatusApproved.ToPtr(),
				},
			},
		},
		&armbatch.PrivateEndpointConnectionClientBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateEndpointConnectionClientUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbatch%2Farmbatch%2Fv0.2.0/sdk/resourcemanager/batch/armbatch/README.md) on how to add the SDK to your project and authenticate.
