package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlansListByBillingAccount.json
func ExampleSavingsPlansClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSavingsPlansClient().NewListByBillingAccountPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", &armbilling.SavingsPlansClientListByBillingAccountOptions{Filter: nil,
		OrderBy:        nil,
		Skiptoken:      nil,
		Take:           to.Ptr[float32](3),
		SelectedState:  to.Ptr("Succeeded"),
		RefreshSummary: to.Ptr("true"),
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
		// page.SavingsPlanModelListResult = armbilling.SavingsPlanModelListResult{
		// 	Value: []*armbilling.SavingsPlanModel{
		// 		{
		// 			Name: to.Ptr("30000000-0000-0000-0000-000000000000"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armbilling.SavingsPlanModelProperties{
		// 				AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeShared),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-16T02:17:03.739Z"); return t}()),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
		// 				BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
		// 				Commitment: &armbilling.Commitment{
		// 					Amount: to.Ptr[float64](10),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbilling.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("SP1"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-16T02:17:04.989Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-16T02:17:03.739Z"); return t}()),
		// 				ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 				ProvisioningState: to.Ptr(armbilling.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-16T02:14:58.230Z"); return t}()),
		// 				Renew: to.Ptr(false),
		// 				Term: to.Ptr(armbilling.SavingsPlanTermP1Y),
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
		// 		},
		// 		{
		// 			Name: to.Ptr("30000000-0000-0000-0000-000000000001"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000001/savingsPlans/30000000-0000-0000-0000-000000000001"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armbilling.SavingsPlanModelProperties{
		// 				AppliedScopeProperties: &armbilling.AppliedScopeProperties{
		// 					DisplayName: to.Ptr("TestRg"),
		// 					ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/TestRg"),
		// 					TenantID: to.Ptr("50000000-0000-0000-0000-000000000000"),
		// 				},
		// 				AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeManagementGroup),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-16T02:22:22.725Z"); return t}()),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
		// 				BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
		// 				Commitment: &armbilling.Commitment{
		// 					Amount: to.Ptr[float64](10),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbilling.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("SP2"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-16T02:22:24.553Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-12-16T02:22:22.725Z"); return t}()),
		// 				ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 				ProvisioningState: to.Ptr(armbilling.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-16T02:20:23.473Z"); return t}()),
		// 				Renew: to.Ptr(false),
		// 				Term: to.Ptr(armbilling.SavingsPlanTermP1Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("ManagementGroup"),
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
		// 		},
		// 		{
		// 			Name: to.Ptr("30000000-0000-0000-0000-000000000002"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000002/savingsPlans/30000000-0000-0000-0000-000000000002"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armbilling.SavingsPlanModelProperties{
		// 				AppliedScopeProperties: &armbilling.AppliedScopeProperties{
		// 					DisplayName: to.Ptr("Azure subscription 1"),
		// 					SubscriptionID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
		// 				},
		// 				AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeSingle),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T21:18:55.296Z"); return t}()),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
		// 				BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingSubscriptions/10000000-0000-0000-0000-000000000000"),
		// 				Commitment: &armbilling.Commitment{
		// 					Amount: to.Ptr[float64](0.025),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbilling.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("SP3"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-07T23:54:43.823Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-14T21:18:55.296Z"); return t}()),
		// 				ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 				ProvisioningState: to.Ptr(armbilling.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-14T21:16:57.234Z"); return t}()),
		// 				Renew: to.Ptr(true),
		// 				Term: to.Ptr(armbilling.SavingsPlanTermP3Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("Single"),
		// 				Utilization: &armbilling.Utilization{
		// 					Aggregates: []*armbilling.UtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](66),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](66),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](65.52),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr("SAME"),
		// 				},
		// 			},
		// 			SKU: &armbilling.SKU{
		// 				Name: to.Ptr("Compute_Savings_Plan"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("30000000-0000-0000-0000-000000000003"),
		// 			Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders/savingsPlans"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000003/savingsPlans/30000000-0000-0000-0000-000000000003"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armbilling.SavingsPlanModelProperties{
		// 				AppliedScopeProperties: &armbilling.AppliedScopeProperties{
		// 					DisplayName: to.Ptr("testRG"),
		// 					ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourcegroups/testRG"),
		// 				},
		// 				AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeSingle),
		// 				BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-08T00:00:06.363Z"); return t}()),
		// 				BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
		// 				BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
		// 				BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
		// 				BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
		// 				Commitment: &armbilling.Commitment{
		// 					Amount: to.Ptr[float64](0.001),
		// 					CurrencyCode: to.Ptr("USD"),
		// 					Grain: to.Ptr(armbilling.CommitmentGrainHourly),
		// 				},
		// 				DisplayName: to.Ptr("SP4"),
		// 				DisplayProvisioningState: to.Ptr("Succeeded"),
		// 				EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-30T21:18:12.196Z"); return t}()),
		// 				ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-11-08T00:00:06.363Z"); return t}()),
		// 				ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
		// 				ProvisioningState: to.Ptr(armbilling.ProvisioningStateSucceeded),
		// 				PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-07T23:57:58.242Z"); return t}()),
		// 				Renew: to.Ptr(false),
		// 				Term: to.Ptr(armbilling.SavingsPlanTermP3Y),
		// 				UserFriendlyAppliedScopeType: to.Ptr("ResourceGroup"),
		// 				Utilization: &armbilling.Utilization{
		// 					Aggregates: []*armbilling.UtilizationAggregates{
		// 						{
		// 							Grain: to.Ptr[float32](1),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](100),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](7),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](100),
		// 							ValueUnit: to.Ptr("percentage"),
		// 						},
		// 						{
		// 							Grain: to.Ptr[float32](30),
		// 							GrainUnit: to.Ptr("days"),
		// 							Value: to.Ptr[float32](100),
		// 							ValueUnit: to.Ptr("percentage"),
		// 					}},
		// 					Trend: to.Ptr("SAME"),
		// 				},
		// 			},
		// 			SKU: &armbilling.SKU{
		// 				Name: to.Ptr("Compute_Savings_Plan"),
		// 			},
		// 	}},
		// 	Summary: &armbilling.SavingsPlanSummaryCount{
		// 		CancelledCount: to.Ptr[float32](0),
		// 		ExpiredCount: to.Ptr[float32](0),
		// 		ExpiringCount: to.Ptr[float32](0),
		// 		FailedCount: to.Ptr[float32](0),
		// 		NoBenefitCount: to.Ptr[float32](0),
		// 		PendingCount: to.Ptr[float32](0),
		// 		ProcessingCount: to.Ptr[float32](0),
		// 		SucceededCount: to.Ptr[float32](3),
		// 		WarningCount: to.Ptr[float32](0),
		// 	},
		// }
	}
}
