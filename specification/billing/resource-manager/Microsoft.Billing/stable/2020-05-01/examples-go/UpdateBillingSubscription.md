Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.1.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

```go
package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingSubscription.json
func ExampleBillingSubscriptionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingSubscriptionsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<billing-account-name>",
		armbilling.BillingSubscription{
			Properties: &armbilling.BillingSubscriptionProperties{
				CostCenter: to.StringPtr("<cost-center>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingSubscription.ID: %s\n", *res.ID)
}
```
