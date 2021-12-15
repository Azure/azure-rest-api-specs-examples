Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmsubscriptions%2Fv0.1.1/sdk/resourcemanager/resources/armsubscriptions/README.md) on how to add the SDK to your project and authenticate.

```go
package armsubscriptions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/GetSubscription.json
func ExampleSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsubscriptions.NewSubscriptionsClient(cred, nil)
	res, err := client.Get(ctx,
		"<subscription-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Subscription.ID: %s\n", *res.ID)
}
```
