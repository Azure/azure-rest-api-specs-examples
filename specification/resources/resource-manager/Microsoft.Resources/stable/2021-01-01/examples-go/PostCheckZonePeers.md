Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmsubscriptions%2Fv0.2.1/sdk/resourcemanager/resources/armsubscriptions/README.md) on how to add the SDK to your project and authenticate.

```go
package armsubscriptions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/PostCheckZonePeers.json
func ExampleClient_CheckZonePeers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsubscriptions.NewClient(cred, nil)
	res, err := client.CheckZonePeers(ctx,
		"<subscription-id>",
		armsubscriptions.CheckZonePeersRequest{
			Location: to.StringPtr("<location>"),
			SubscriptionIDs: []*string{
				to.StringPtr("subscriptions/11111111-1111-1111-1111-111111111111")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientCheckZonePeersResult)
}
```
