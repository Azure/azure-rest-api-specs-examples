package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/getPolicyAssignmentWithOverrides.json
func ExampleAssignmentsClient_Get_retrieveAPolicyAssignmentWithOverrides() {
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
	// 			InstanceID: to.Ptr("d2f3a4b5-c6d7-8e9f-0a1b-2c3d4e5f6a7b"),
	// 			Metadata: map[string]any{
	// 				"assignedBy": "Special Someone",
	// 			},
	// 			NotScopes: []*string{
	// 			},
	// 			Overrides: []*armpolicy.Override{
	// 				{
	// 					Kind: to.Ptr(armpolicy.OverrideKindPolicyEffect),
	// 					Selectors: []*armpolicy.Selector{
	// 						{
	// 							In: []*string{
	// 								to.Ptr("Limit_Skus"),
	// 								to.Ptr("Limit_Locations"),
	// 							},
	// 							Kind: to.Ptr(armpolicy.SelectorKindPolicyDefinitionReferenceID),
	// 						},
	// 					},
	// 					Value: to.Ptr("Audit"),
	// 				},
	// 			},
	// 			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
	// 			Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
	// 		},
	// 	},
	// }
}
