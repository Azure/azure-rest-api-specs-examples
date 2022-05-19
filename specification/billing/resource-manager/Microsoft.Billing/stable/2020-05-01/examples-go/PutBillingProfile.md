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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
func ExampleProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewProfilesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"{billingAccountName}",
		"{billingProfileName}",
		armbilling.Profile{
			Properties: &armbilling.ProfileProperties{
				BillTo: &armbilling.AddressDetails{
					AddressLine1: to.Ptr("Test Address 1"),
					City:         to.Ptr("Redmond"),
					Country:      to.Ptr("US"),
					FirstName:    to.Ptr("Test"),
					LastName:     to.Ptr("User"),
					PostalCode:   to.Ptr("12345"),
					Region:       to.Ptr("WA"),
				},
				DisplayName: to.Ptr("Finance"),
				EnabledAzurePlans: []*armbilling.AzurePlan{
					{
						SKUID: to.Ptr("0001"),
					},
					{
						SKUID: to.Ptr("0002"),
					}},
				InvoiceEmailOptIn: to.Ptr(true),
				PoNumber:          to.Ptr("ABC12345"),
			},
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
