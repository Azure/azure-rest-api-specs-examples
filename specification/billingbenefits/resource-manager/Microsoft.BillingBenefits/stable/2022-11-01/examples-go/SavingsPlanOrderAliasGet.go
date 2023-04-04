package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanOrderAliasGet.json
func ExampleSavingsPlanOrderAliasClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanOrderAliasClient().Get(ctx, "spAlias123", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanOrderAliasModel = armbillingbenefits.SavingsPlanOrderAliasModel{
	// 	Name: to.Ptr("spAlias123"),
	// 	Type: to.Ptr("Microsoft.BillingBenefits/savingsPlanOrderAliases"),
	// 	ID: to.Ptr("/providers/microsoft.billingbenefits/savingsPlanOrderAliases/spAlias123"),
	// 	Properties: &armbillingbenefits.SavingsPlanOrderAliasProperties{
	// 		AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
	// 			ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
	// 		},
	// 		AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
	// 		BillingPlan: to.Ptr(armbillingbenefits.BillingPlanP1M),
	// 		BillingScopeID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
	// 		Commitment: &armbillingbenefits.Commitment{
	// 			Amount: to.Ptr[float64](15.23),
	// 			CurrencyCode: to.Ptr("USD"),
	// 			Grain: to.Ptr(armbillingbenefits.CommitmentGrainHourly),
	// 		},
	// 		DisplayName: to.Ptr("ComputeSavingsPlan_2021-07-01"),
	// 		ProvisioningState: to.Ptr(armbillingbenefits.ProvisioningStateSucceeded),
	// 		SavingsPlanOrderID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/30000000-0000-0000-0000-000000000000"),
	// 		Term: to.Ptr(armbillingbenefits.TermP1Y),
	// 	},
	// 	SKU: &armbillingbenefits.SKU{
	// 		Name: to.Ptr("Compute_Savings_Plan"),
	// 	},
	// }
}
