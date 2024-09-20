package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionAliasGet.json
func ExampleSubscriptionsAliasesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsAliasesClient().Get(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "c356b7c7-7545-4686-b843-c1a49cf853fc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SubscriptionAlias = armbilling.SubscriptionAlias{
	// 	Name: to.Ptr("c356b7c7-7545-4686-b843-c1a49cf853fc"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingSubscriptionAliases"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptionAliases/c356b7c7-7545-4686-b843-c1a49cf853fc"),
	// 	SystemData: &armbilling.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-25T22:39:34.260Z"); return t}()),
	// 	},
	// 	Properties: &armbilling.SubscriptionAliasProperties{
	// 		AutoRenew: to.Ptr(armbilling.AutoRenewOn),
	// 		BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
	// 		BillingProfileName: to.Ptr("xxxx-xxxx-xxx-xxx"),
	// 		DisplayName: to.Ptr("Billing Subscription Display Name"),
	// 		InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
	// 		InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
	// 		InvoiceSectionName: to.Ptr("yyyy-yyyy-yyy-yyy"),
	// 		ProductCategory: to.Ptr("SeatBased"),
	// 		ProductType: to.Ptr("Seat-Based Product Type"),
	// 		ProductTypeID: to.Ptr("XYZ56789"),
	// 		PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		Quantity: to.Ptr[int64](1),
	// 		SKUDescription: to.Ptr("SKU Description"),
	// 		SKUID: to.Ptr("0001"),
	// 		Status: to.Ptr(armbilling.BillingSubscriptionStatusActive),
	// 		SystemOverrides: &armbilling.SystemOverrides{
	// 			CancellationAllowedEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-01T22:39:34.260Z"); return t}()),
	// 		},
	// 		TermDuration: to.Ptr("P1M"),
	// 		TermEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-05T22:39:34.260Z"); return t}()),
	// 		TermStartDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-05T22:39:34.260Z"); return t}()),
	// 		BillingSubscriptionID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 	},
	// }
}
