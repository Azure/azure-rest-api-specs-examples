package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyAssignmentWithUserAssignedIdentity.json
func ExampleAssignmentsClient_Get_retrieveAPolicyAssignmentWithAUserAssignedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssignmentsClient().Get(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "EnforceNaming", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assignment = armpolicy.Assignment{
	// 	Name: to.Ptr("EnforceNaming"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyAssignments"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/EnforceNaming"),
	// 	Identity: &armpolicy.Identity{
	// 		Type: to.Ptr(armpolicy.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armpolicy.UserAssignedIdentitiesValue{
	// 			"/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/testResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity": &armpolicy.UserAssignedIdentitiesValue{
	// 				ClientID: to.Ptr("4bee2b8a-1bee-47c2-90e9-404241551135"),
	// 				PrincipalID: to.Ptr("e6d23f8d-af97-4fbc-bda6-00604e4e3d0a"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armpolicy.AssignmentProperties{
	// 		Description: to.Ptr("Force resource names to begin with given DeptA and end with -LC"),
	// 		DisplayName: to.Ptr("Enforce resource naming rules"),
	// 		EnforcementMode: to.Ptr(armpolicy.EnforcementModeDefault),
	// 		Metadata: map[string]any{
	// 			"assignedBy": "Special Someone",
	// 		},
	// 		NotScopes: []*string{
	// 		},
	// 		Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 			"prefix": &armpolicy.ParameterValuesValue{
	// 				Value: "DeptA",
	// 			},
	// 			"suffix": &armpolicy.ParameterValuesValue{
	// 				Value: "-LC",
	// 			},
	// 		},
	// 		PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 		Scope: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
	// 	},
	// }
}
