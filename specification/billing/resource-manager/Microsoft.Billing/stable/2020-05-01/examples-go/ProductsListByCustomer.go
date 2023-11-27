package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ProductsListByCustomer.json
func ExampleProductsClient_NewListByCustomerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductsClient().NewListByCustomerPager("{billingAccountName}", "{customerName}", nil)
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
		// page.ProductsListResult = armbilling.ProductsListResult{
		// 	TotalCount: to.Ptr[int32](2),
		// 	Value: []*armbilling.Product{
		// 		{
		// 			Name: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/products/00000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.ProductProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				AvailabilityID: to.Ptr("AvailabilityId1"),
		// 				BillingFrequency: to.Ptr(armbilling.BillingFrequencyOneTime),
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				CustomerDisplayName: to.Ptr("Customer 1"),
		// 				CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}"),
		// 				DisplayName: to.Ptr("Eng Reservation (1a13s21awe)"),
		// 				LastCharge: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](5000),
		// 				},
		// 				LastChargeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T17:32:28.000Z"); return t}()),
		// 				ProductType: to.Ptr("Reservation"),
		// 				ProductTypeID: to.Ptr("A12345"),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-01T17:32:28.000Z"); return t}()),
		// 				Quantity: to.Ptr[float32](4),
		// 				Reseller: &armbilling.Reseller{
		// 					Description: to.Ptr("Reseller1"),
		// 					ResellerID: to.Ptr("2c917292-b7bc-42f2-99a2-e498b9087c06"),
		// 				},
		// 				SKUDescription: to.Ptr("Enterprise Agreement Development"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusTypeActive),
		// 				TenantID: to.Ptr("515a6d36-aaf8-4ca2-a5e8-c45deb0c5cce"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("10000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/BillingAccounts/{billingAccountName}/products/10000000-0000-0000-0000-000000000001"),
		// 			Properties: &armbilling.ProductProperties{
		// 				AutoRenew: to.Ptr(armbilling.AutoRenewOn),
		// 				AvailabilityID: to.Ptr("AvailabilityId1"),
		// 				BillingFrequency: to.Ptr(armbilling.BillingFrequencyMonthly),
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/11000000-0000-0000-0000-000000000000"),
		// 				CustomerDisplayName: to.Ptr("Customer 1"),
		// 				CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/customers/{customerName}"),
		// 				DisplayName: to.Ptr("Engineering Email"),
		// 				LastCharge: &armbilling.Amount{
		// 					Currency: to.Ptr("USD"),
		// 					Value: to.Ptr[float32](500),
		// 				},
		// 				LastChargeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-15T17:32:28.000Z"); return t}()),
		// 				ProductType: to.Ptr("Azure subscription"),
		// 				ProductTypeID: to.Ptr("A12345"),
		// 				PurchaseDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-01T17:32:28.000Z"); return t}()),
		// 				Quantity: to.Ptr[float32](4),
		// 				SKUDescription: to.Ptr("Enterprise Agreement Development"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusTypeActive),
		// 				TenantID: to.Ptr("515a6d36-aaf8-4ca2-a5e8-c45deb0c5cce"),
		// 			},
		// 	}},
		// }
	}
}
