package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: 2025-03-01/createOrUpdatePolicyDefinition.json
func ExampleDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicyDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDefinitionsClient().CreateOrUpdate(ctx, "ResourceNaming", armpolicy.Definition{
		Properties: &armpolicy.DefinitionProperties{
			Description: to.Ptr("Force resource names to begin with given 'prefix' and/or end with given 'suffix'"),
			DisplayName: to.Ptr("Enforce resource naming convention"),
			Metadata: map[string]any{
				"category": "Naming",
			},
			Mode: to.Ptr("All"),
			Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
				"prefix": {
					Type: to.Ptr(armpolicy.ParameterTypeString),
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("Resource name prefix"),
						DisplayName: to.Ptr("Prefix"),
					},
				},
				"suffix": {
					Type: to.Ptr(armpolicy.ParameterTypeString),
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						Description: to.Ptr("Resource name suffix"),
						DisplayName: to.Ptr("Suffix"),
					},
				},
			},
			PolicyRule: map[string]any{
				"if": map[string]any{
					"not": map[string]any{
						"field": "name",
						"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
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
