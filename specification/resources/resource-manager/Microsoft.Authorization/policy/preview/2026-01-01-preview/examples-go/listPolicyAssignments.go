package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/listPolicyAssignments.json
func ExampleAssignmentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssignmentsClient().NewListPager(&armpolicy.AssignmentsClientListOptions{
		Filter: to.Ptr("atScope()"),
		Expand: to.Ptr("LatestDefinitionVersion, EffectiveDefinitionVersion")})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armpolicy.AssignmentsClientListResponse{
		// 	AssignmentListResult: armpolicy.AssignmentListResult{
		// 		Value: []*armpolicy.Assignment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
		// 				Name: to.Ptr("CostManagement"),
		// 				Location: to.Ptr("eastus"),
		// 				Identity: &armpolicy.Identity{
		// 					Type: to.Ptr(armpolicy.ResourceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("e6d23f8d-af97-4fbc-bda6-00604e4e3d0a"),
		// 					TenantID: to.Ptr("4bee2b8a-1bee-47c2-90e9-404241551135"),
		// 				},
		// 				Properties: &armpolicy.AssignmentProperties{
		// 					DisplayName: to.Ptr("Storage Cost Management"),
		// 					Description: to.Ptr("Minimize the risk of accidental cost overruns"),
		// 					Metadata: map[string]any{
		// 						"category": "Cost Management",
		// 					},
		// 					PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/storageSkus"),
		// 					DefinitionVersion: to.Ptr("1.*.*"),
		// 					LatestDefinitionVersion: to.Ptr("1.0.0"),
		// 					EffectiveDefinitionVersion: to.Ptr("1.0.0"),
		// 					Parameters: map[string]*armpolicy.ParameterValuesValue{
		// 						"allowedSkus": &armpolicy.ParameterValuesValue{
		// 							Value: "Standard_A1",
		// 						},
		// 					},
		// 					Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
		// 					NotScopes: []*string{
		// 					},
		// 					InstanceID: to.Ptr("a3c4d5e6-f7a8-9b0c-1d2e-3f4a5b6c7d8e"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/TagEnforcement"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
		// 				Name: to.Ptr("TagEnforcement"),
		// 				Properties: &armpolicy.AssignmentProperties{
		// 					DisplayName: to.Ptr("Enforces a tag key and value"),
		// 					Description: to.Ptr("Ensure a given tag key and value are present on all resources"),
		// 					PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/TagKeyValue"),
		// 					DefinitionVersion: to.Ptr("1.*.*"),
		// 					LatestDefinitionVersion: to.Ptr("1.0.0"),
		// 					EffectiveDefinitionVersion: to.Ptr("1.0.0"),
		// 					Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
		// 					NotScopes: []*string{
		// 					},
		// 					InstanceID: to.Ptr("b6d7e8f9-a0b1-2c3d-4e5f-6a7b8c9d0e1f"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
