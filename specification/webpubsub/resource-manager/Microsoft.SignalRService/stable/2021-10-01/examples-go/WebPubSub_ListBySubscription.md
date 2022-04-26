Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwebpubsub%2Farmwebpubsub%2Fv0.4.0/sdk/resourcemanager/webpubsub/armwebpubsub/README.md) on how to add the SDK to your project and authenticate.

```go
package armwebpubsub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSub_ListBySubscription.json
func ExampleClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armwebpubsub.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListBySubscriptionPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
