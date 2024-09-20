package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoiceSectionsCreateOrUpdate.json
func ExampleInvoiceSectionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInvoiceSectionsClient().BeginCreateOrUpdate(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "invoice-section-1", armbilling.InvoiceSection{
		Properties: &armbilling.InvoiceSectionProperties{
			DisplayName: to.Ptr("Invoice Section 1"),
			Tags: map[string]*string{
				"costCategory": to.Ptr("Support"),
				"pcCode":       to.Ptr("A123456"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InvoiceSection = armbilling.InvoiceSection{
	// 	Name: to.Ptr("invoice-section-1"),
	// 	Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/invoiceSections"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/invoice-section-1"),
	// 	Properties: &armbilling.InvoiceSectionProperties{
	// 		DisplayName: to.Ptr("Invoice Section 1"),
	// 		SystemID: to.Ptr("yyyy-yyyy-yyy-yyy"),
	// 		Tags: map[string]*string{
	// 			"costCategory": to.Ptr("Support"),
	// 			"pcCode": to.Ptr("A123456"),
	// 		},
	// 	},
	// }
}
