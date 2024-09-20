package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/productsListByCustomer.json
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
	pager := clientFactory.NewProductsClient().NewListByCustomerPager("a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31", "Q7GV-UUVA-PJA-TGB", &armbilling.ProductsClientListByCustomerOptions{Filter: nil,
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
		// 			Name: to.Ptr("6b96d3f2-9008-4a9d-912f-f87744185aa3"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProducts/6b96d3f2-9008-4a9d-912f-f87744185aa3"),
		// 			Properties: &armbilling.ProductProperties{
		// 				BillingProfileDisplayName: to.Ptr("BillingProfile1"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998"),
		// 				CustomerDisplayName: to.Ptr("Customer1"),
		// 				CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/customers/Q7GV-UUVA-PJA-TGB"),
		// 				DisplayName: to.Ptr("My Product"),
		// 				SKUDescription: to.Ptr("Microsoft Azure Standard"),
		// 				SKUID: to.Ptr("0002"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CD0BBD7A-461A-4D9A-9C59-EAE51A0D5D12"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProducts/CD0BBD7A-461A-4D9A-9C59-EAE51A0D5D12"),
		// 			Properties: &armbilling.ProductProperties{
		// 				BillingProfileDisplayName: to.Ptr("BillingProfile1"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998"),
		// 				CustomerDisplayName: to.Ptr("Customer1"),
		// 				CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/customers/Q7GV-UUVA-PJA-TGB"),
		// 				DisplayName: to.Ptr("Test Product"),
		// 				SKUDescription: to.Ptr("Microsoft Azure Standard"),
		// 				SKUID: to.Ptr("0002"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("9e90f86b-22fc-42f3-bfe2-0ac3e7c01d32"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProducts/9e90f86b-22fc-42f3-bfe2-0ac3e7c01d32"),
		// 			Properties: &armbilling.ProductProperties{
		// 				BillingProfileDisplayName: to.Ptr("BillingProfile1"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998"),
		// 				CustomerDisplayName: to.Ptr("Customer1"),
		// 				CustomerID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/customers/Q7GV-UUVA-PJA-TGB"),
		// 				DisplayName: to.Ptr("Dev Test Product"),
		// 				SKUDescription: to.Ptr("Domain Name Registration ORG"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 	}},
		// }
	}
}
