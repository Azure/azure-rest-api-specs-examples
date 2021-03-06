package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/UpdateBillingAccount.json
func ExampleAccountsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbilling.NewAccountsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"{billingAccountName}",
		armbilling.AccountUpdateRequest{
			Properties: &armbilling.AccountProperties{
				DisplayName: to.Ptr("Test Account"),
				SoldTo: &armbilling.AddressDetails{
					AddressLine1: to.Ptr("Test Address 1"),
					City:         to.Ptr("Redmond"),
					CompanyName:  to.Ptr("Contoso"),
					Country:      to.Ptr("US"),
					FirstName:    to.Ptr("Test"),
					LastName:     to.Ptr("User"),
					PostalCode:   to.Ptr("12345"),
					Region:       to.Ptr("WA"),
				},
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
