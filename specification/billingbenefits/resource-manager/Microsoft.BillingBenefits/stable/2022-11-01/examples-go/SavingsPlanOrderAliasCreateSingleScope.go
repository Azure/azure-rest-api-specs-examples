package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderAliasCreateSingleScope.json
func ExampleSavingsPlanOrderAliasClient_BeginCreate_savingsPlanOrderAliasCreateSingleScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSavingsPlanOrderAliasClient().BeginCreate(ctx, "spAlias123", armbillingbenefits.SavingsPlanOrderAliasModel{
		Properties: &armbillingbenefits.SavingsPlanOrderAliasProperties{
			AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
				SubscriptionID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000"),
			},
			AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
			BillingPlan:      to.Ptr(armbillingbenefits.BillingPlanP1M),
			BillingScopeID:   to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234567/billingSubscriptions/30000000-0000-0000-0000-000000000000"),
			Commitment: &armbillingbenefits.Commitment{
				Amount:       to.Ptr[float64](0.001),
				CurrencyCode: to.Ptr("USD"),
				Grain:        to.Ptr(armbillingbenefits.CommitmentGrainHourly),
			},
			DisplayName: to.Ptr("Compute_SavingsPlan_10-28-2022_16-38"),
			Term:        to.Ptr(armbillingbenefits.TermP3Y),
		},
		SKU: &armbillingbenefits.SKU{
			Name: to.Ptr("Compute_Savings_Plan"),
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
	// res.SavingsPlanOrderAliasModel = armbillingbenefits.SavingsPlanOrderAliasModel{
	// 	Name: to.Ptr("SavingsPlan_1667000324595"),
	// 	Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrderAliases"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrderAliases/SavingsPlan_1667000324595"),
	// 	Properties: &armbillingbenefits.SavingsPlanOrderAliasProperties{
	// 		AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
	// 			SubscriptionID: to.Ptr("/subscriptions/30000000-0000-0000-0000-000000000000"),
	// 		},
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingScopeID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/1234567/billingSubscriptions/30000000-0000-0000-0000-000000000000"),
	// 		Commitment: &armbillingbenefits.Commitment{
	// 			Amount: to.Ptr[float64](0.001),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("Compute_SavingsPlan_10-28-2022_16-38"),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		SavingsPlanOrderID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/30000000-0000-0000-0000-000000000023"),
	// 		Term: to.Ptr(armbillingbenefits.TermP3Y),
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}
