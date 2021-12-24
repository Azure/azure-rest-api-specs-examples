Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsubscription%2Farmsubscription%2Fv0.1.0/sdk/resourcemanager/subscription/armsubscription/README.md) on how to add the SDK to your project and authenticate.

```go
package armsubscription_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
)

// x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/acceptSubscriptionOwnership.json
func ExampleSubscriptionClient_BeginAcceptOwnership() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsubscription.NewSubscriptionClient(cred, nil)
	poller, err := client.BeginAcceptOwnership(ctx,
		"<subscription-id>",
		armsubscription.AcceptOwnershipRequest{
			Properties: &armsubscription.AcceptOwnershipRequestProperties{
				DisplayName:       to.StringPtr("<display-name>"),
				ManagementGroupID: to.StringPtr("<management-group-id>"),
				Tags: map[string]*string{
					"tag1": to.StringPtr("Messi"),
					"tag2": to.StringPtr("Ronaldo"),
					"tag3": to.StringPtr("Lebron"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
