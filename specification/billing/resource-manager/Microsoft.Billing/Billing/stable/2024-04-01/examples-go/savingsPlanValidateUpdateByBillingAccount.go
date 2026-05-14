package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/savingsPlanValidateUpdateByBillingAccount.json
func ExampleSavingsPlansClient_ValidateUpdateByBillingAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlansClient().ValidateUpdateByBillingAccount(ctx, "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbilling.SavingsPlanUpdateValidateRequest{
		Benefits: []*armbilling.SavingsPlanUpdateRequestProperties{
			{
				AppliedScopeProperties: &armbilling.AppliedScopeProperties{
					SubscriptionID: to.Ptr("/subscriptions/50000000-0000-0000-0000-000000000000"),
				},
				AppliedScopeType: to.Ptr(armbilling.AppliedScopeTypeSingle),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.SavingsPlansClientValidateUpdateByBillingAccountResponse{
	// 	SavingsPlanValidateResponse: armbilling.SavingsPlanValidateResponse{
	// 		Benefits: []*armbilling.SavingsPlanValidResponseProperty{
	// 			{
	// 				Valid: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
