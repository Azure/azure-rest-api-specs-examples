package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/savingsPlanOrderListByBillingAccount.json
func ExampleSavingsPlanOrdersClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSavingsPlanOrdersClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", nil)
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
		// page = armbilling.SavingsPlanOrdersClientListByBillingAccountResponse{
		// 	SavingsPlanOrderModelList: armbilling.SavingsPlanOrderModelList{
		// 		Value: []*armbilling.SavingsPlanOrderModel{
		// 			{
		// 				Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000000"),
		// 				Properties: &armbilling.SavingsPlanOrderModelProperties{
		// 					BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-04T03:18:31.3072635Z"); return t}()),
		// 					BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 					BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
		// 					BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
		// 					BillingScopeID: to.Ptr("10000000-0000-0000-0000-000000000000"),
		// 					DisplayName: to.Ptr("SP1"),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-02-04T03:18:31.3072635Z"); return t}()),
		// 					ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					SavingsPlans: []*string{
		// 						to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
		// 					},
		// 					Term: to.Ptr(armbilling.SavingsPlanTermP3Y),
		// 				},
		// 				SKU: &armbilling.SKU{
		// 					Name: to.Ptr("Compute_Savings_Plan"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("20000000-0000-0000-0000-000000000001"),
		// 				Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000001"),
		// 				Properties: &armbilling.SavingsPlanOrderModelProperties{
		// 					BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-04T03:22:19.7304330Z"); return t}()),
		// 					BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 					BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
		// 					BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
		// 					BillingScopeID: to.Ptr("10000000-0000-0000-0000-000000000000"),
		// 					DisplayName: to.Ptr("SP2"),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-02-04T03:22:19.7304330Z"); return t}()),
		// 					ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					SavingsPlans: []*string{
		// 						to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000001/savingsPlans/30000000-0000-0000-0000-000000000001"),
		// 					},
		// 					Term: to.Ptr(armbilling.SavingsPlanTermP3Y),
		// 				},
		// 				SKU: &armbilling.SKU{
		// 					Name: to.Ptr("Compute_Savings_Plan"),
		// 				},
		// 				Tags: map[string]*string{
		// 					"key1": to.Ptr("value1"),
		// 					"key2": to.Ptr("value2"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
