package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionGetByBillingProfile.json
func ExampleSubscriptionsClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsClient().GetByBillingProfile(ctx, "pcn.94077792", "6478903", "6b96d3f2-9008-4a9d-912f-f87744185aa3", &armbilling.SubscriptionsClientGetByBillingProfileOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Subscription = armbilling.Subscription{
	// 	Name: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/billingSubscriptions"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/pcn.94077792/billingProfiles/6478903/billingSubscriptions/6b96d3f2-9008-4a9d-912f-f87744185aa3"),
	// 	Properties: &armbilling.SubscriptionProperties{
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/pcn.94077792/billingProfiles/6478903"),
	// 		BillingProfileName: to.Ptr("47268432"),
	// 		DisplayName: to.Ptr("My Subscription"),
	// 		InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6564892/billingProfiles/11000000-0000-0000-0000-000000000000/invoiceSections/22000000-0000-0000-0000-000000000000"),
	// 		ProvisioningTenantID: to.Ptr("7d62dbd1-1348-4496-b2f8-df6948c103ee"),
	// 		Quantity: to.Ptr[int64](50),
	// 		SKUDescription: to.Ptr("Microsoft Azure Dev/Test"),
	// 		SKUID: to.Ptr("0001"),
	// 		Status: to.Ptr(armbilling.BillingSubscriptionStatusActive),
	// 		SubscriptionID: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
	// 		SystemOverrides: &armbilling.SystemOverrides{
	// 			Cancellation: to.Ptr(armbilling.CancellationAllowed),
	// 			CancellationAllowedEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		},
	// 	},
	// }
}
