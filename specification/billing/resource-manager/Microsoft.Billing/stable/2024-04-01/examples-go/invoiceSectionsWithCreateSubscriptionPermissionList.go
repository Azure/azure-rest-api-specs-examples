package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoiceSectionsWithCreateSubscriptionPermissionList.json
func ExampleAccountsClient_NewListInvoiceSectionsByCreateSubscriptionPermissionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListInvoiceSectionsByCreateSubscriptionPermissionPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", &armbilling.AccountsClientListInvoiceSectionsByCreateSubscriptionPermissionOptions{Filter: nil})
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
		// page.InvoiceSectionWithCreateSubPermissionListResult = armbilling.InvoiceSectionWithCreateSubPermissionListResult{
		// 	Value: []*armbilling.InvoiceSectionWithCreateSubPermission{
		// 		{
		// 			BillingProfileDisplayName: to.Ptr("Billing Profile Display Name"),
		// 			BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx"),
		// 			BillingProfileSpendingLimit: to.Ptr(armbilling.SpendingLimitOn),
		// 			BillingProfileStatus: to.Ptr(armbilling.BillingProfileStatusWarned),
		// 			BillingProfileStatusReasonCode: to.Ptr(armbilling.BillingProfileStatusReasonCodePastDue),
		// 			BillingProfileSystemID: to.Ptr("33000000-0000-0000-0000-000000000000"),
		// 			EnabledAzurePlans: []*armbilling.AzurePlan{
		// 				{
		// 					SKUDescription: to.Ptr("Microsoft Azure Plan"),
		// 					SKUID: to.Ptr("0001"),
		// 				},
		// 				{
		// 					SKUDescription: to.Ptr("Microsoft Azure Plan for DevTest"),
		// 					SKUID: to.Ptr("0002"),
		// 			}},
		// 			InvoiceSectionDisplayName: to.Ptr("Invoice Section Display Name"),
		// 			InvoiceSectionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/invoiceSections/yyyy-yyyy-yyy-yyy"),
		// 			InvoiceSectionSystemID: to.Ptr("22000000-0000-0000-0000-000000000000"),
		// 	}},
		// }
	}
}
