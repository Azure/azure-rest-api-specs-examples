```go
package armsubscriptions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PostCheckZonePeers.json
func ExampleClient_CheckZonePeers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckZonePeers(ctx,
		"00000000-0000-0000-0000-00000000000000",
		armsubscriptions.CheckZonePeersRequest{
			Location: to.Ptr("eastus"),
			SubscriptionIDs: []*string{
				to.Ptr("subscriptions/11111111-1111-1111-1111-111111111111")},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmsubscriptions%2Fv1.0.0/sdk/resourcemanager/resources/armsubscriptions/README.md) on how to add the SDK to your project and authenticate.
