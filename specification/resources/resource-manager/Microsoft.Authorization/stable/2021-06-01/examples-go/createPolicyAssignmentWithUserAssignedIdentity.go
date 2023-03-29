package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createPolicyAssignmentWithUserAssignedIdentity.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithAUserAssignedIdentity() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Create(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "EnforceNaming", armpolicy.Assignment{
		Identity: &armpolicy.Identity{
			Type: to.Ptr(armpolicy.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armpolicy.UserAssignedIdentitiesValue{
				"/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/testResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-identity": {},
			},
		},
		Location: to.Ptr("eastus"),
		Properties: &armpolicy.AssignmentProperties{
			Description:     to.Ptr("Force resource names to begin with given DeptA and end with -LC"),
			DisplayName:     to.Ptr("Enforce resource naming rules"),
			EnforcementMode: to.Ptr(armpolicy.EnforcementModeDefault),
			Metadata: map[string]any{
				"assignedBy": "Foo Bar",
			},
			Parameters: map[string]*armpolicy.ParameterValuesValue{
				"prefix": {
					Value: "DeptA",
				},
				"suffix": {
					Value: "-LC",
				},
			},
			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
