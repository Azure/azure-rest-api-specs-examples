package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/productsListByBillingAccount.json
func ExampleProductsClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductsClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", &armbilling.ProductsClientListByBillingAccountOptions{Filter: nil,
		OrderBy: nil,
		Top:     nil,
		Skip:    nil,
		Count:   nil,
		Search:  nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ProductListResult = armbilling.ProductListResult{
		// 	Value: []*armbilling.Product{
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/products/90000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.ProductProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				DisplayName: to.Ptr("My product"),
		// 				InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				ProductType: to.Ptr("Seat-Based Product Type"),
		// 				ProductTypeID: to.Ptr("XYZ56789"),
		// 				PurchaseDate: to.Ptr("2023-01-04T22:39:34.2606750Z"),
		// 				Quantity: to.Ptr[int64](1),
		// 				SKUDescription: to.Ptr("SKU Description"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/products/90000000-0000-0000-0000-000000000001"),
		// 			Properties: &armbilling.ProductProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				DisplayName: to.Ptr("Test product"),
		// 				InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				ProductType: to.Ptr("Software Product Type"),
		// 				ProductTypeID: to.Ptr("EFG456"),
		// 				PurchaseDate: to.Ptr("2023-01-06T22:39:34.2606750Z"),
		// 				Quantity: to.Ptr[int64](1),
		// 				SKUDescription: to.Ptr("SKU Description"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90000000-0000-0000-0000-000000000002"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/products/90000000-0000-0000-0000-000000000002"),
		// 			Properties: &armbilling.ProductProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				DisplayName: to.Ptr("Dev product"),
		// 				InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 				ProductType: to.Ptr("Reservation Product Type"),
		// 				ProductTypeID: to.Ptr("JKL789"),
		// 				PurchaseDate: to.Ptr("2023-01-05T22:39:34.2606750Z"),
		// 				Quantity: to.Ptr[int64](1),
		// 				SKUDescription: to.Ptr("SKU Description"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusSuspended),
		// 			},
		// 	}},
		// }
	}
}
