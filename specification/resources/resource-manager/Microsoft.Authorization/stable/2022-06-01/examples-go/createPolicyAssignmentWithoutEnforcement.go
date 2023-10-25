package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignmentWithoutEnforcement.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithoutEnforcingPolicyEffectDuringResourceCreationOrUpdate() {
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
		Properties: &armpolicy.AssignmentProperties{
			Description:     to.Ptr("Force resource names to begin with given DeptA and end with -LC"),
			DisplayName:     to.Ptr("Enforce resource naming rules"),
			EnforcementMode: to.Ptr(armpolicy.EnforcementModeDoNotEnforce),
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
