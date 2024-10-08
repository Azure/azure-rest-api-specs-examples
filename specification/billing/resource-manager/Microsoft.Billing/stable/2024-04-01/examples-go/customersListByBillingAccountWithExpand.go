package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/customersListByBillingAccountWithExpand.json
func ExampleCustomersClient_NewListByBillingAccountPager_customersListByBillingAccountWithExpand() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomersClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", &armbilling.CustomersClientListByBillingAccountOptions{Expand: to.Ptr("enabledAzurePlans,resellers"),
		Filter:  nil,
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
		// page.CustomerListResult = armbilling.CustomerListResult{
		// 	Value: []*armbilling.Customer{
		// 		{
		// 			Name: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/customers"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/customers/11111111-1111-1111-1111-111111111111"),
		// 			Properties: &armbilling.CustomerProperties{
		// 				BillingProfileDisplayName: to.Ptr("Contoso Operations Billing"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 				DisplayName: to.Ptr("customer1"),
		// 				EnabledAzurePlans: []*armbilling.AzurePlan{
		// 					{
		// 						SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 						SKUID: to.Ptr("0002"),
		// 				}},
		// 				Resellers: []*armbilling.Reseller{
		// 					{
		// 						Description: to.Ptr("Reseller1"),
		// 						ResellerID: to.Ptr("89e87bdf-a2a2-4687-925f-4c18b27bccfd"),
		// 					},
		// 					{
		// 						Description: to.Ptr("Reseller2"),
		// 						ResellerID: to.Ptr("3b65b5a8-bd4f-4084-90e9-e1bd667a2b19"),
		// 				}},
		// 				Status: to.Ptr(armbilling.CustomerStatusActive),
		// 				SystemID: to.Ptr("yyyy-yyyy-yyy-yyy"),
		// 				Tags: map[string]*string{
		// 					"costCategory": to.Ptr("Support"),
		// 					"pcCode": to.Ptr("A123456"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
