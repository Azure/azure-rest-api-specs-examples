package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanUpdate.json
func ExampleSavingsPlanClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanClient().Update(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbillingbenefits.SavingsPlanUpdateRequest{
		Properties: &armbillingbenefits.SavingsPlanUpdateRequestProperties{
			AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
				ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
			},
			AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
			DisplayName:      to.Ptr("TestDisplayName"),
			Renew:            to.Ptr(true),
			RenewProperties: &armbillingbenefits.RenewProperties{
				PurchaseProperties: &armbillingbenefits.PurchaseRequest{
					Properties: &armbillingbenefits.PurchaseRequestProperties{
						AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
							ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
						},
						AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
						BillingPlan:      to.Ptr(armbillingbenefits.BillingPlanP1M),
						BillingScopeID:   to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
						Commitment: &armbillingbenefits.Commitment{
							Amount:       to.Ptr[float64](15.23),
							CurrencyCode: to.Ptr("USD"),
							Grain:        to.Ptr(armbillingbenefits.CommitmentGrainHourly),
						},
						DisplayName: to.Ptr("TestDisplayName_renewed"),
						Renew:       to.Ptr(false),
						Term:        to.Ptr(armbillingbenefits.TermP1Y),
					},
					SKU: &armbillingbenefits.SKU{
						Name: to.Ptr("Compute_Savings_Plan"),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanModel = armbillingbenefits.SavingsPlanModel{
	// 	Name: to.Ptr("30000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("microsoft.billingbenefits/savingsPlanOrders/savingsPlans"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
	// 	Properties: &armbillingbenefits.SavingsPlanModelProperties{
	// 		AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
	// 			DisplayName: to.Ptr("Azure subscription 1"),
	// 			SubscriptionID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
	// 		},
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
	// 		BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-27T00:34:33.669Z"); return t}()),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 		BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
	// 		Commitment: &armbillingbenefits.Commitment{
	// 			Amount: to.Ptr[float64](0.001),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("riName"),
	// 		DisplayProvisioningState: to.Ptr("Succeeded"),
	// 		EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-10T00:12:54.549Z"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-27T00:34:33.669Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-27T00:32:45.582Z"); return t}()),
	// 		Renew: to.Ptr(true),
	// 		Term: to.Ptr(armbillingbenefits.TermP3Y),
	// 		UserFriendlyAppliedScopeType: to.Ptr("Single"),
	// 		Utilization: &armbillingbenefits.Utilization{
	// 			Aggregates: []*armbillingbenefits.UtilizationAggregates{
	// 				{
	// 					Grain: to.Ptr[float32](1),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](100),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](7),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](37),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](30),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](53.85),
	// 					ValueUnit: to.Ptr("percentage"),
	// 			}},
	// 			Trend: to.Ptr("DOWN"),
	// 		},
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}
