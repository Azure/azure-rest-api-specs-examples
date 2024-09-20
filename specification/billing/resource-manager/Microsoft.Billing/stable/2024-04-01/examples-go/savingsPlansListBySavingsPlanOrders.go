package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlansListBySavingsPlanOrders.json
func ExampleSavingsPlansClient_NewListBySavingsPlanOrderPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSavingsPlansClient().NewListBySavingsPlanOrderPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", nil)
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
		// page.SavingsPlanModelList = armbilling.SavingsPlanModelList{
		// 	Value: []*armbilling.SavingsPlanModel{
		// 		{
		// 			Name: to.Ptr("30000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbilling.SavingsPlanModelProperties{
		// 				AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeShared),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-20T02:36:22.339Z"); return t}()),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
		// 				BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
		// 				Commitment: &armbilling.Commitment{
		// 					Amount: to.Ptr[float64](0.001),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbilling.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("SP1"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-20T02:36:25.089Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-01-20T02:36:22.339Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armbilling.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-20T02:34:38.126Z"); return t}()),
		// 				Renew: to.Ptr(false),
		// 				Term: to.Ptr(armbilling.SavingsPlanTermP3Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Shared"),
		// 				Utilization: &armbilling.Utilization{
		// 					Aggregates: []*armbilling.UtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](0),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr("SAME"),
		// 				},
		// 			},
		// 			SKU: &armbilling.SKU{
		// 				Name: to.Ptr("Compute_Savings_Plan"),
		// 			},
		// 	}},
		// }
	}
}
