package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinition.json
func ExampleDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicyDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
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
