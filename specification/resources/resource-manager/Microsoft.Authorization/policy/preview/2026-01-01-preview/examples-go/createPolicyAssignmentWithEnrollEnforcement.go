package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/createPolicyAssignmentWithEnrollEnforcement.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentToEnforcePolicyEffectOnlyOnEnrolledResourcesDuringResourceCreationOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Create(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "EnforceNamingEnroll", armpolicy.Assignment{
		Properties: &armpolicy.AssignmentProperties{
			Description:     to.Ptr("Force resource names to begin with given DeptA and end with -LC"),
			DisplayName:     to.Ptr("Enforce resource naming rules"),
			EnforcementMode: to.Ptr(armpolicy.EnforcementModeEnroll),
			Metadata: map[string]any{
				"assignedBy": "Special Someone",
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
