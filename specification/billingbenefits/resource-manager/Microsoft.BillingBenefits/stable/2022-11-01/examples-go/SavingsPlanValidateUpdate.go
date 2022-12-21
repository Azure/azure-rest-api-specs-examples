package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/billingbenefits/resource-manager/Microsoft.BillingBenefits/stable/2022-11-01/examples/SavingsPlanValidateUpdate.json
func ExampleSavingsPlanClient_ValidateUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbillingbenefits.NewSavingsPlanClient(nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ValidateUpdate(ctx, "20000000-0000-0000-0000-000000000000", "30000000-0000-0000-0000-000000000000", armbillingbenefits.SavingsPlanUpdateValidateRequest{
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
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
