package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-06-01/createPolicyAssignmentWithResourceRolloutPercentageSelector.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithAResourceRolloutPercentageSelector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAssignmentsClient().Create(ctx, "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2", "CostManagement", armpolicy.Assignment{
		Properties: &armpolicy.AssignmentProperties{
			Description: to.Ptr("Limit resources by rollout percentage"),
			DisplayName: to.Ptr("Limit resources by rollout percentage"),
			Metadata: map[string]any{
				"assignedBy": "Special Someone",
			},
			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
			ResourceSelectors: []*armpolicy.ResourceSelector{
				{
					Name: to.Ptr("SDPRollout"),
					Selectors: []*armpolicy.Selector{
						{
							Kind:     to.Ptr(armpolicy.SelectorKindResourceRolloutPercentage),
							Progress: to.Ptr[int32](80),
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
