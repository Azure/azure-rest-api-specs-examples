Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fbilling%2Farmbilling%2Fv0.4.0/sdk/resourcemanager/billing/armbilling/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/PutBillingProfile.json
func ExampleProfilesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armbilling.NewProfilesClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<billing-account-name>",
		"<billing-profile-name>",
		armbilling.Profile{
			Properties: &armbilling.ProfileProperties{
				BillTo: &armbilling.AddressDetails{
					AddressLine1: to.Ptr("<address-line1>"),
					City:         to.Ptr("<city>"),
					Country:      to.Ptr("<country>"),
					FirstName:    to.Ptr("<first-name>"),
					LastName:     to.Ptr("<last-name>"),
					PostalCode:   to.Ptr("<postal-code>"),
					Region:       to.Ptr("<region>"),
				},
				DisplayName: to.Ptr("<display-name>"),
				EnabledAzurePlans: []*armbilling.AzurePlan{
					{
						SKUID: to.Ptr("<skuid>"),
					},
					{
						SKUID: to.Ptr("<skuid>"),
					}},
				InvoiceEmailOptIn: to.Ptr(true),
				PoNumber:          to.Ptr("<po-number>"),
			},
		},
		&armbilling.ProfilesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
