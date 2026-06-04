package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/createOrUpdatePolicyDefinitionExternalEvaluationEnforcementSettings.json
func ExampleDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicyDefinitionWithExternalEvaluationEnforcementSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().CreateOrUpdate(ctx, "RandomizeVMAllocation", armpolicy.Definition{
		Properties: &armpolicy.DefinitionProperties{
			Description: to.Ptr("Randomly disable VM allocation in eastus by having policy rule reference the outcome of invoking an external endpoint using the CoinFlip endpoint that returns random values."),
			DisplayName: to.Ptr("Randomize VM Allocation"),
			ExternalEvaluationEnforcementSettings: &armpolicy.ExternalEvaluationEnforcementSettings{
				EndpointSettings: &armpolicy.ExternalEvaluationEndpointSettings{
					Kind: to.Ptr("CoinFlip"),
					Details: map[string]any{
						"successProbability": 0.5,
					},
				},
				MissingTokenAction: to.Ptr("audit"),
				RoleDefinitionIDs: []*string{
					to.Ptr("subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/roleDefinitions/f0cc2aea-b517-48f6-8f9e-0c01c687907b"),
				},
			},
			Metadata: map[string]any{
				"category": "VM",
			},
			Mode: to.Ptr("Indexed"),
			PolicyRule: map[string]any{
				"if": map[string]any{
					"allOf": []any{
						map[string]any{
							"equals": "Microsoft.Compute/virtualMachines",
							"field":  "type",
						},
						map[string]any{
							"equals": "eastus",
							"field":  "location",
						},
						map[string]any{
							"equals": "false",
							"value":  "[claims().isValid]",
						},
					},
				},
				"then": map[string]any{
					"effect": "deny",
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
