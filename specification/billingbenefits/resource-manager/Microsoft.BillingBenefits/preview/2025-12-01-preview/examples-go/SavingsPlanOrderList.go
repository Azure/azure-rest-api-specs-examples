package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/SavingsPlanOrderList.json
func ExampleSavingsPlanOrderClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSavingsPlanOrderClient().NewListPager(nil)
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
		// page = armbillingbenefits.SavingsPlanOrderClientListResponse{
		// 	SavingsPlanOrderModelList: armbillingbenefits.SavingsPlanOrderModelList{
		// 		Value: []*armbillingbenefits.SavingsPlanOrderModel{
		// 			{
		// 				Name: to.Ptr("20000000-0000-0000-0000-000000000000"),
		// 				Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders"),
		// 				ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000"),
		// 				Properties: &armbillingbenefits.SavingsPlanOrderModelProperties{
		// 					BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000004:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 					BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
		// 					BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000004:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 					BillingScopeID: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 					DisplayName: to.Ptr("Compute_SavingsPlan_10-28-2022_16-38"),
		// 					ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-28T23:40:30.6028303Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
		// 					SavingsPlans: []*string{
		// 						to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/20000000-0000-0000-0000-000000000001"),
		// 					},
		// 					Term: to.Ptr(armbillingbenefits.TermP3Y),
		// 				},
		// 				SKU: &armbillingbenefits.ResourceSKU{
		// 					Name: to.Ptr("Compute_Savings_Plan"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
