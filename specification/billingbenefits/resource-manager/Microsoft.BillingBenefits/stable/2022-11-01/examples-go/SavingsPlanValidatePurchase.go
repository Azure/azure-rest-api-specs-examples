package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f790e624d0d080b89d962a3bd19c65bc6a6b2f5e/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanValidatePurchase.json
func ExampleRPClient_ValidatePurchase() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRPClient().ValidatePurchase(ctx, armbillingbenefits.SavingsPlanPurchaseValidateRequest{
		Benefits: []*armbillingbenefits.SavingsPlanOrderAliasModel{
			{
				Properties: &armbillingbenefits.SavingsPlanOrderAliasProperties{
					AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
						ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/testrg"),
					},
					AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
					BillingScopeID:   to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
					Commitment: &armbillingbenefits.Commitment{
						Amount:       to.Ptr[float64](15.23),
						CurrencyCode: to.Ptr("USD"),
						Grain:        to.Ptr(armbillingbenefits.CommitmentGrainHourly),
					},
					DisplayName: to.Ptr("ComputeSavingsPlan_2021-07-01"),
					Term:        to.Ptr(armbillingbenefits.TermP1Y),
				},
				SKU: &armbillingbenefits.SKU{
					Name: to.Ptr("Compute_Savings_Plan"),
				},
			},
			{
				Properties: &armbillingbenefits.SavingsPlanOrderAliasProperties{
					AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
						ResourceGroupID: to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000/resourceGroups/RG"),
					},
					AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeSingle),
					BillingScopeID:   to.Ptr("/subscriptions/10000000-0000-0000-0000-000000000000"),
					Commitment: &armbillingbenefits.Commitment{
						Amount:       to.Ptr[float64](20),
						CurrencyCode: to.Ptr("USD"),
						Grain:        to.Ptr(armbillingbenefits.CommitmentGrainHourly),
					},
					DisplayName: to.Ptr("ComputeSavingsPlan_2021-07-01"),
					Term:        to.Ptr(armbillingbenefits.TermP1Y),
				},
				SKU: &armbillingbenefits.SKU{
					Name: to.Ptr("Compute_Savings_Plan"),
				},
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SavingsPlanValidateResponse = armbillingbenefits.SavingsPlanValidateResponse{
	// 	Benefits: []*armbillingbenefits.SavingsPlanValidResponseProperty{
	// 		{
	// 			Reason: to.Ptr("Your provider has not enabled Savings Plan purchase on this subscription."),
	// 			ReasonCode: to.Ptr("CustomerCannotPurchaseSavingsPlan"),
	// 			Valid: to.Ptr(false),
	// 		},
	// 		{
	// 			Valid: to.Ptr(true),
	// 	}},
	// }
}
