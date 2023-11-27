package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlansList.json
func ExampleSavingsPlanClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSavingsPlanClient().NewListAllPager(&armbillingbenefits.SavingsPlanClientListAllOptions{Filter: to.Ptr("(properties%2farchived+eq+false)"),
		Orderby:        to.Ptr("properties/displayName asc"),
		RefreshSummary: nil,
		Skiptoken:      to.Ptr[float32](50),
		SelectedState:  nil,
		Take:           to.Ptr[float32](1),
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
		// page.SavingsPlanModelListResult = armbillingbenefits.SavingsPlanModelListResult{
		// 	AdditionalProperties: []*armbillingbenefits.SavingsPlanSummary{
		// 		{
		// 			Name: to.Ptr("summary"),
		// 			Value: &armbillingbenefits.SavingsPlanSummaryCount{
		// 				CancelledCount: to.Ptr[float32](0),
		// 				ExpiredCount: to.Ptr[float32](0),
		// 				ExpiringCount: to.Ptr[float32](0),
		// 				FailedCount: to.Ptr[float32](0),
		// 				NoBenefitCount: to.Ptr[float32](0),
		// 				PendingCount: to.Ptr[float32](0),
		// 				ProcessingCount: to.Ptr[float32](0),
		// 				SucceededCount: to.Ptr[float32](1),
		// 				WarningCount: to.Ptr[float32](0),
		// 			},
		// 	}},
		// 	Value: []*armbillingbenefits.SavingsPlanModel{
		// 		{
		// 			Name: to.Ptr("20000000-0000-0000-0000-000000000000/30000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
		// 			Properties: &armbillingbenefits.SavingsPlanModelProperties{
		// 				AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
		// 					DisplayName: to.Ptr("Azure subscription 1"),
		// 					SubscriptionID: to.Ptr("/subscriptions/20000000-0000-0000-0000-000000000005"),
		// 				},
		// 				AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
		// 				BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000009"),
		// 				Commitment: &armbillingbenefits.Commitment{
		// 					Amount: to.Ptr[float64](0.001),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("Compute_SavingsPlan_10-19-2022_11-03"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-19T18:05:37.103Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T18:05:36.525Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-19T18:03:56.403Z"); return t}()),
		// 				Renew: to.Ptr(false),
		// 				Term: to.Ptr(armbillingbenefits.TermP1Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Single"),
		// 				Utilization: &armbillingbenefits.Utilization{
		// 					Aggregates: []*armbillingbenefits.UtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](100),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](78),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](78.12),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr(""),
		// 				},
		// 			},
		// 			SKU: &armbillingbenefits.SKU{
		// 				Name: to.Ptr("Compute_Savings_Plan"),
		// 			},
		// 	}},
		// }
	}
}
