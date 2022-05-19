Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.5.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

```go
package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdatePolicy.json
func ExamplePoliciesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewPoliciesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		armbilling.Policy{
			Properties: &armbilling.PolicyProperties{
				MarketplacePurchases: to.Ptr(armbilling.MarketplacePurchasesPolicyOnlyFreeAllowed),
				ReservationPurchases: to.Ptr(armbilling.ReservationPurchasesPolicyNotAllowed),
				ViewCharges:          to.Ptr(armbilling.ViewChargesPolicyAllowed),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
