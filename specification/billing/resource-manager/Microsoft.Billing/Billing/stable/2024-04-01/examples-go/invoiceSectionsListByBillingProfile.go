package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/invoiceSectionsListByBillingProfile.json
func ExampleInvoiceSectionsClient_NewListByBillingProfilePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInvoiceSectionsClient().NewListByBillingProfilePager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", &armbilling.InvoiceSectionsClientListByBillingProfileOptions{
		IncludeDeleted: to.Ptr(true)})
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
		// page = armbilling.InvoiceSectionsClientListByBillingProfileResponse{
		// 	InvoiceSectionListResult: armbilling.InvoiceSectionListResult{
		// 		Value: []*armbilling.InvoiceSection{
		// 			{
		// 				Name: to.Ptr("invoice-section-1"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/invoice-section-1"),
		// 				Properties: &armbilling.InvoiceSectionProperties{
		// 					DisplayName: to.Ptr("Invoice Section 1"),
		// 					State: to.Ptr(armbilling.InvoiceSectionStateActive),
		// 					SystemID: to.Ptr("yyyy-yyyy-yyy-yyy"),
		// 					Tags: map[string]*string{
		// 						"costCategory": to.Ptr("Support"),
		// 						"pcCode": to.Ptr("A123456"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("invoice-section-2"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/invoice-section-2"),
		// 				Properties: &armbilling.InvoiceSectionProperties{
		// 					DisplayName: to.Ptr("Invoice Section 2"),
		// 					ReasonCode: to.Ptr(armbilling.InvoiceSectionStateReasonCodePastDue),
		// 					State: to.Ptr(armbilling.InvoiceSectionStateWarned),
		// 					SystemID: to.Ptr("zzzz-zzzz-zzz-zzz"),
		// 					Tags: map[string]*string{
		// 						"costCategory": to.Ptr("Marketing"),
		// 						"pcCode": to.Ptr("Z345678"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("invoice-section-3"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/invoice-section-3"),
		// 				Properties: &armbilling.InvoiceSectionProperties{
		// 					DisplayName: to.Ptr("Invoice Section 3"),
		// 					State: to.Ptr(armbilling.InvoiceSectionStateDeleted),
		// 					SystemID: to.Ptr("aaaa-aaaa-aaa-aaa"),
		// 					Tags: map[string]*string{
		// 						"costCategory": to.Ptr("Support"),
		// 						"pcCode": to.Ptr("A123456"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
