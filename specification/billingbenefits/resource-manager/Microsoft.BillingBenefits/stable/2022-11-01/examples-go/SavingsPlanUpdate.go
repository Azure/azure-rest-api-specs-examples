package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanUpdate.json
func ExampleSavingsPlanClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbillingbenefits.NewSavingsPlanClient(nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbillingbenefits.SavingsPlanUpdateRequest{
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
	// TODO: use response item
	_ = res
}
