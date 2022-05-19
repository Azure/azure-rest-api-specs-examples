Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsubscription%2Farmsubscription%2Fv1.0.0/sdk/resourcemanager/subscription/armsubscription/README.md) on how to add the SDK to your project and authenticate.

```go
package armsubscription_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/acceptSubscriptionOwnership.json
func ExampleClient_BeginAcceptOwnership() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsubscription.NewClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAcceptOwnership(ctx,
		"291bba3f-e0a5-47bc-a099-3bdcb2a50a05",
		armsubscription.AcceptOwnershipRequest{
			Properties: &armsubscription.AcceptOwnershipRequestProperties{
				DisplayName: to.Ptr("Test Subscription"),
				Tags: map[string]*string{
					"tag1": to.Ptr("Messi"),
					"tag2": to.Ptr("Ronaldo"),
					"tag3": to.Ptr("Lebron"),
				},
			},
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
