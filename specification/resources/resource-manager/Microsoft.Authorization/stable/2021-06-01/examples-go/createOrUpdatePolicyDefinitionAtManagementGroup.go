package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/createOrUpdatePolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_CreateOrUpdateAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armpolicy.NewDefinitionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.CreateOrUpdateAtManagementGroup(ctx,
		"ResourceNaming",
		"MyManagementGroup",
		armpolicy.Definition{
			Properties: &armpolicy.DefinitionProperties{
				Description: to.Ptr("Force resource names to begin with given 'prefix' and/or end with given 'suffix'"),
				DisplayName: to.Ptr("Enforce resource naming convention"),
				Metadata: map[string]interface{}{
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
				PolicyRule: map[string]interface{}{
					"if": map[string]interface{}{
						"not": map[string]interface{}{
							"field": "name",
							"like":  "[concat(parameters('prefix'), '*', parameters('suffix'))]",
						},
					},
					"then": map[string]interface{}{
						"effect": "deny",
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
