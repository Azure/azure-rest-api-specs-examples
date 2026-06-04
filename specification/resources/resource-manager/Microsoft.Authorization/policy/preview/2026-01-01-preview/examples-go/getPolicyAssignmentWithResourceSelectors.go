package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/getPolicyAssignmentWithResourceSelectors.json
func ExampleAssignmentsClient_Get_retrieveAPolicyAssignmentWithResourceSelectors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().Get(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "CostManagement", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicy.AssignmentsClientGetResponse{
	// 	Assignment: armpolicy.Assignment{
	// 		Name: to.Ptr("CostManagement"),
	// 		Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
	// 		ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
	// 		Properties: &armpolicy.AssignmentProperties{
	// 			Description: to.Ptr("Limit the resource location and resource SKU"),
	// 			DefinitionVersion: to.Ptr("1.*.*"),
	// 			DisplayName: to.Ptr("Limit the resource location and resource SKU"),
	// 			EnforcementMode: to.Ptr(armpolicy.EnforcementModeDefault),
	// 			InstanceID: to.Ptr("a3c4d5e6-f7a8-9b0c-1d2e-3f4a5b6c7d8e"),
	// 			Metadata: map[string]any{
	// 				"assignedBy": "Special Someone",
	// 			},
	// 			NotScopes: []*string{
	// 			},
	// 			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
	// 			ResourceSelectors: []*armpolicy.ResourceSelector{
	// 				{
	// 					Name: to.Ptr("SDPRegions"),
	// 					Selectors: []*armpolicy.Selector{
	// 						{
	// 							In: []*string{
	// 								to.Ptr("eastus2euap"),
	// 								to.Ptr("centraluseuap"),
	// 							},
	// 							Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
	// 		},
	// 	},
	// }
}
