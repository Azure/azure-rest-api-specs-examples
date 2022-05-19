Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatashare%2Farmdatashare%2Fv1.0.0/sdk/resourcemanager/datashare/armdatashare/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatashare_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datashare/resource-manager/Microsoft.DataShare/stable/2020-09-01/examples/ShareSubscriptions_CancelSynchronization.json
func ExampleShareSubscriptionsClient_BeginCancelSynchronization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatashare.NewShareSubscriptionsClient("12345678-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCancelSynchronization(ctx,
		"SampleResourceGroup",
		"Account1",
		"ShareSubscription1",
		armdatashare.ShareSubscriptionSynchronization{
			SynchronizationID: to.Ptr("7d0536a6-3fa5-43de-b152-3d07c4f6b2bb"),
		},
		nil)
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
