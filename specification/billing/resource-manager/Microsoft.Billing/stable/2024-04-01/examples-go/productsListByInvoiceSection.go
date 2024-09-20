package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/productsListByInvoiceSection.json
func ExampleProductsClient_NewListByInvoiceSectionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductsClient().NewListByInvoiceSectionPager("a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31", "ea36e548-1505-41db-bebc-46fff3d37998", "Q7GV-UUVA-PJA-TGB", &armbilling.ProductsClientListByInvoiceSectionOptions{Filter: nil,
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
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998"),
		// 				DisplayName: to.Ptr("My Product"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/invoiceSections/Q7GV-UUVA-PJA-TGB"),
		// 				SKUDescription: to.Ptr("Microsoft Azure Dev/Test"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CD0BBD7A-461A-4D9A-9C59-EAE51A0D5D12"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProducts/CD0BBD7A-461A-4D9A-9C59-EAE51A0D5D12"),
		// 			Properties: &armbilling.ProductProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998"),
		// 				DisplayName: to.Ptr("Test Product"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/invoiceSections/Q7GV-UUVA-PJA-TGB"),
		// 				SKUDescription: to.Ptr("Microsoft Azure Dev/Test"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("90D31600-D46D-4084-84C6-1235CDAFFB56"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProducts/90D31600-D46D-4084-84C6-1235CDAFFB56"),
		// 			Properties: &armbilling.ProductProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998"),
		// 				DisplayName: to.Ptr("Dev Product"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/invoiceSections/Q7GV-UUVA-PJA-TGB"),
		// 				SKUDescription: to.Ptr("Microsoft Azure Dev/Test"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("9e90f86b-22fc-42f3-bfe2-0ac3e7c01d32"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/products"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProducts/9e90f86b-22fc-42f3-bfe2-0ac3e7c01d32"),
		// 			Properties: &armbilling.ProductProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso operations billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998"),
		// 				DisplayName: to.Ptr("Dev Test Product"),
		// 				InvoiceSectionDisplayName: to.Ptr("Contoso operations invoiceSection"),
		// 				InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31/billingProfiles/ea36e548-1505-41db-bebc-46fff3d37998/invoiceSections/Q7GV-UUVA-PJA-TGB"),
		// 				SKUDescription: to.Ptr("Domain Name Registration ORG"),
		// 				SKUID: to.Ptr("0001"),
		// 				Status: to.Ptr(armbilling.ProductStatusActive),
		// 			},
		// 	}},
		// }
	}
}
