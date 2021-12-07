Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.1.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

```go
package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateCustomerPolicy.json
func ExamplePoliciesClient_UpdateCustomer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewPoliciesClient(cred, nil)
	res, err := client.UpdateCustomer(ctx,
		"<billing-account-name>",
		"<customer-name>",
		armbilling.CustomerPolicy{
			Properties: &armbilling.CustomerPolicyProperties{
				ViewCharges: armbilling.ViewChargesNotAllowed.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CustomerPolicy.ID: %s\n", *res.ID)
}
```
