package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/SavingsPlanValidateUpdate.json
func ExampleSavingsPlanClient_ValidateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanClient().ValidateUpdate(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbillingbenefits.SavingsPlanUpdateValidateRequest{
		Benefits: []*armbillingbenefits.SavingsPlanUpdateRequestProperties{
			{
				AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
					ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/30000000-0000-0000-0000-000000000100"),
					TenantID:          to.Ptr("30000000-0000-0000-0000-000000000100"),
				},
				AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeManagementGroup),
			},
			{
				AppliedScopeProperties: &armbillingbenefits.AppliedScopeProperties{
					ManagementGroupID: to.Ptr("/providers/Microsoft.Management/managementGroups/MockMG"),
					TenantID:          to.Ptr("30000000-0000-0000-0000-000000000100"),
				},
				AppliedScopeType: to.Ptr(armbillingbenefits.AppliedScopeTypeManagementGroup),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbillingbenefits.SavingsPlanClientValidateUpdateResponse{
	// 	SavingsPlanValidateResponse: armbillingbenefits.SavingsPlanValidateResponse{
	// 		Benefits: []*armbillingbenefits.SavingsPlanValidResponseProperty{
	// 			{
	// 				Valid: to.Ptr(true),
	// 			},
	// 			{
	// 				Valid: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
