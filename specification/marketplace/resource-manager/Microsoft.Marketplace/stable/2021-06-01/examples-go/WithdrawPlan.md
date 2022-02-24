Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmarketplace%2Farmmarketplace%2Fv0.2.1/sdk/resourcemanager/marketplace/armmarketplace/README.md) on how to add the SDK to your project and authenticate.

```go
package armmarketplace_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace"
)

// x-ms-original-file: specification/marketplace/resource-manager/Microsoft.Marketplace/stable/2021-06-01/examples/WithdrawPlan.json
func ExamplePrivateStoreClient_WithdrawPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmarketplace.NewPrivateStoreClient(cred, nil)
	_, err = client.WithdrawPlan(ctx,
		"<private-store-id>",
		"<request-approval-id>",
		&armmarketplace.PrivateStoreClientWithdrawPlanOptions{Payload: &armmarketplace.WithdrawProperties{
			Properties: &armmarketplace.WithdrawDetails{
				PlanID:      to.StringPtr("<plan-id>"),
				PublisherID: to.StringPtr("<publisher-id>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
```
