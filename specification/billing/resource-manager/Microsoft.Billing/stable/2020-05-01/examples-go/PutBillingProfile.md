Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.1.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

```go
package armbilling_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
func ExampleBillingProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbilling.NewBillingProfilesClient(cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		armbilling.BillingProfile{
			Properties: &armbilling.BillingProfileProperties{
				BillTo: &armbilling.AddressDetails{
					AddressLine1: to.StringPtr("<address-line1>"),
					City:         to.StringPtr("<city>"),
					Country:      to.StringPtr("<country>"),
					FirstName:    to.StringPtr("<first-name>"),
					LastName:     to.StringPtr("<last-name>"),
					PostalCode:   to.StringPtr("<postal-code>"),
					Region:       to.StringPtr("<region>"),
				},
				DisplayName: to.StringPtr("<display-name>"),
				EnabledAzurePlans: []*armbilling.AzurePlan{
					{
						SKUID: to.StringPtr("<skuid>"),
					},
					{
						SKUID: to.StringPtr("<skuid>"),
					}},
				InvoiceEmailOptIn: to.BoolPtr(true),
				PoNumber:          to.StringPtr("<po-number>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BillingProfile.ID: %s\n", *res.ID)
}
```
