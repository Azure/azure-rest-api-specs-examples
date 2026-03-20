package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: 2025-03-01/createPolicyAssignmentWithOverrides.json
func ExampleAssignmentsClient_Create_createOrUpdateAPolicyAssignmentWithOverrides() {
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
			Description:       to.Ptr("Limit the resource location and resource SKU"),
			DefinitionVersion: to.Ptr("1.*.*"),
			DisplayName:       to.Ptr("Limit the resource location and resource SKU"),
			Metadata: map[string]any{
				"assignedBy": "Special Someone",
			},
			Overrides: []*armpolicy.Override{
				{
					Kind: to.Ptr(armpolicy.OverrideKindPolicyEffect),
					Selectors: []*armpolicy.Selector{
						{
							In: []*string{
								to.Ptr("Limit_Skus"),
								to.Ptr("Limit_Locations"),
							},
							Kind: to.Ptr(armpolicy.SelectorKindPolicyDefinitionReferenceID),
						},
					},
					Value: to.Ptr("Audit"),
				},
				{
					Kind: to.Ptr(armpolicy.OverrideKindDefinitionVersion),
					Selectors: []*armpolicy.Selector{
						{
							In: []*string{
								to.Ptr("eastUSEuap"),
								to.Ptr("centralUSEuap"),
							},
							Kind: to.Ptr(armpolicy.SelectorKindResourceLocation),
						},
					},
					Value: to.Ptr("2.*.*"),
				},
			},
			PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
