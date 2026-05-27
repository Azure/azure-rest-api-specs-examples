package armbillingbenefits_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
)

// Generated from example definition: 2025-12-01-preview/SavingsPlanOrderElevate.json
func ExampleSavingsPlanOrderClient_Elevate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbillingbenefits.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSavingsPlanOrderClient().Elevate(ctx, "20000000-0000-0000-0000-000000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbillingbenefits.SavingsPlanOrderClientElevateResponse{
	// 	RoleAssignmentEntity: armbillingbenefits.RoleAssignmentEntity{
	// 		Name: to.Ptr("70000000-0000-0000-0000-000000000005"),
	// 		ID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000009/providers/Microsoft.Authorization/roleAssignments/70000000-0000-0000-0000-000000000005"),
	// 		Properties: &armbillingbenefits.RoleAssignmentEntityProperties{
	// 			PrincipalID: to.Ptr("50000000-0000-0000-0000-000000000000"),
	// 			RoleDefinitionID: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000009/providers/Microsoft.Authorization/roleDefinitions/30000000-0000-0000-0000-000000000008"),
	// 			Scope: to.Ptr("/providers/Microsoft.BillingBenefits/savingsPlanOrders/20000000-0000-0000-0000-000000000009"),
	// 		},
	// 	},
	// }
}
