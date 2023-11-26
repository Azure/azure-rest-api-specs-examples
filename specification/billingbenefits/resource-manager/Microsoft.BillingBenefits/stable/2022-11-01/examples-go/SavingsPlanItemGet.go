package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanItemGet.json
func ExampleSavingsPlanClient_Get_savingsPlanItemGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanClient().Get(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", &armbillingbenefits.SavingsPlanClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanModel = armbillingbenefits.SavingsPlanModel{
	// 	Name: to.Ptr("20000000-0000-0000-0000-000000000000/30000000-0000-0000-0000-000000000000"),
	// 	Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrders/savingsPlans"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000000/savingsPlans/30000000-0000-0000-0000-000000000000"),
	// 	Properties: &armbillingbenefits.SavingsPlanModelProperties{
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeShared),
	// 		BillingAccountID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31"),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingProfileID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000000:20000000-0000-3000-0000-000000000000_2019-05-31/billingProfiles/KPSV-DWNE-BG7-TGB"),
	// 		BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/20000000-1000-0000-0000-000000000009"),
	// 		Commitment: &armbillingbenefits.Commitment{
	// 			Amount: to.Ptr[float64](0.001),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("Compute_SavingsPlan_patch_rename2"),
	// 		DisplayProvisioningState: to.Ptr("Succeeded"),
	// 		EffectiveDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-21T18:15:42.409Z"); return t}()),
	// 		ExpiryDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-10-18T21:16:13.185Z"); return t}()),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		PurchaseDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-18T21:14:26.827Z"); return t}()),
	// 		Renew: to.Ptr(true),
	// 		Term: to.Ptr(armbillingbenefits.TermP3Y),
	// 		UserFriendlyAppliedScopeType: to.Ptr("Shared"),
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
	// 					Value: to.Ptr[float32](84),
	// 					ValueUnit: to.Ptr("percentage"),
	// 				},
	// 				{
	// 					Grain: to.Ptr[float32](30),
	// 					GrainUnit: to.Ptr("days"),
	// 					Value: to.Ptr[float32](83.87),
	// 					ValueUnit: to.Ptr("percentage"),
	// 			}},
	// 			Trend: to.Ptr(""),
	// 		},
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}
