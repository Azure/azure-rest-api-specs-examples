package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: 2025-03-01/deletePolicyAssignment.json
func ExampleAssignmentsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().Delete(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "EnforceNaming", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicy.AssignmentsClientDeleteResponse{
	// 	Assignment: &armpolicy.Assignment{
	// 		Name: to.Ptr("EnforceNaming"),
	// 		Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
	// 		ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/EnforceNaming"),
	// 		Properties: &armpolicy.AssignmentProperties{
	// 			Description: to.Ptr("Force resource names to begin with given DeptA and end with -LC"),
	// 			DefinitionVersion: to.Ptr("1.*.*"),
	// 			DisplayName: to.Ptr("Enforce resource naming rules"),
	// 			InstanceID: to.Ptr("e4b0f5a6-7c8d-4e9f-8a1b-2c3d4e5f6a7b"),
	// 			Metadata: map[string]any{
	// 				"assignedBy": "Special Someone",
	// 			},
	// 			NotScopes: []*string{
	// 			},
	// 			Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 				"prefix": &armpolicy.ParameterValuesValue{
	// 					Value: "DeptA",
	// 				},
	// 				"suffix": &armpolicy.ParameterValuesValue{
	// 					Value: "-LC",
	// 				},
	// 			},
	// 			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 			Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
	// 		},
	// 	},
	// }
}
