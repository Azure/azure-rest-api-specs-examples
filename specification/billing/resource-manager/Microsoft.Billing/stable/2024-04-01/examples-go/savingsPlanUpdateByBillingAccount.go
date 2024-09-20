package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/savingsPlanUpdateByBillingAccount.json
func ExampleSavingsPlansClient_BeginUpdateByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSavingsPlansClient().BeginUpdateByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbilling.SavingsPlanUpdateRequest{
		Properties: &armbilling.SavingsPlanUpdateRequestProperties{
			AppliedScopeProperties: &armbilling.AppliedScopeProperties{
				ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/mg1"),
				TenantID:          to.Ptr("80000000-0000-0000-0000-000000000000"),
			},
			AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeManagementGroup),
			DisplayName:      to.Ptr("sp_newName"),
			Renew:            to.Ptr(true),
			RenewProperties: &armbilling.RenewProperties{
				PurchaseProperties: &armbilling.PurchaseRequest{
					Properties: &armbilling.PurchaseRequestProperties{
						AppliedScopeProperties: &armbilling.AppliedScopeProperties{
							ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/mg1"),
							TenantID:          to.Ptr("80000000-0000-0000-0000-000000000000"),
						},
						AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeManagementGroup),
						BillingPlan:      to.Ptr(armbilling.BillingPlanP1M),
						BillingScopeID:   to.Ptr("/subscriptions/50000000-0000-0000-0000-000000000000"),
						Commitment: &armbilling.Commitment{
							Amount:       to.Ptr[float64](0.001),
							CurrencyCode: to.Ptr("USD"),
							Grain:        to.Ptr(armbilling.CommitmentGrainHourly),
						},
						DisplayName: to.Ptr("sp_newName_renewed"),
						Term:        to.Ptr(armbilling.SavingsPlanTermP3Y),
					},
					SKU: &armbilling.SKU{
						Name: to.Ptr("Compute_Savings_Plan"),
					},
				},
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
	// res.SavingsPlanModel = armbilling.SavingsPlanModel{
	// 	Name: to.Ptr("30000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("microsoft.billing/billingAccounts/savingsPlanOrders/savingsPlans"),
	// 	ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armbilling.SavingsPlanModelProperties{
	// 		AppliedScopeProperties: &armbilling.AppliedScopeProperties{
	// 			DisplayName: to.Ptr("mg1"),
	// 			ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/mg1"),
	// 			TenantID: to.Ptr("80000000-0000-0000-0000-000000000000"),
	// 		},
	// 		AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeManagementGroup),
	// 		BenefitStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-04T03:18:31.307Z"); return t}()),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbilling.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/AAAA-BBBB-CCC-DDD"),
	// 		BillingScopeID: to.Ptr("/subscriptions/50000000-0000-0000-0000-000000000000"),
	// 		Commitment: &armbilling.Commitment{
	// 			Amount: to.Ptr[float64](0.001),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbilling.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("sp_newName"),
	// 		DisplayProvisioningState: to.Ptr("Succeeded"),
	// 		EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-04T04:30:49.198Z"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-02-04T03:18:31.307Z"); return t}()),
	// 		ProductCode: to.Ptr("20000000-0000-0000-0000-000000000005"),
	// 		ProvisioningState: to.Ptr(armbilling.ProvisioningStateSucceeded),
	// 		PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-02-04T03:16:38.906Z"); return t}()),
	// 		Renew: to.Ptr(true),
	// 		Term: to.Ptr(armbilling.SavingsPlanTermP3Y),
	// 		UserFriendlyAppliedScopeType: to.Ptr("ManagementGroup"),
	// 		Utilization: &armbilling.Utilization{
	// 			Aggregates: []*armbilling.UtilizationAggregates{
	// 			},
	// 		},
	// 	},
	// 	SKU: &armbilling.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}
