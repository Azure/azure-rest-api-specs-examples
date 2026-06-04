package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/listAllPolicyDefinitionVersions.json
func ExampleDefinitionVersionsClient_ListAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("ae640e6b-ba3e-4256-9d62-2993eecfa6f2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefinitionVersionsClient().ListAll(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicy.DefinitionVersionsClientListAllResponse{
	// 	DefinitionVersionListResult: armpolicy.DefinitionVersionListResult{
	// 		Value: []*armpolicy.DefinitionVersion{
	// 			{
	// 				Name: to.Ptr("1.2.1"),
	// 				Type: to.Ptr("Microsoft.Authorization/policyDefinitions/versions"),
	// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming/versions/1.2.1"),
	// 				Properties: &armpolicy.DefinitionVersionProperties{
	// 					Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
	// 					DisplayName: to.Ptr("Naming Convention"),
	// 					Metadata: map[string]any{
	// 						"category": "Naming",
	// 					},
	// 					Mode: to.Ptr("All"),
	// 					Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 						"prefix": &armpolicy.ParameterDefinitionsValue{
	// 							Type: to.Ptr(armpolicy.ParameterTypeString),
	// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 								Description: to.Ptr("Resource name prefix"),
	// 								DisplayName: to.Ptr("Prefix"),
	// 							},
	// 						},
	// 						"suffix": &armpolicy.ParameterDefinitionsValue{
	// 							Type: to.Ptr(armpolicy.ParameterTypeString),
	// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 								Description: to.Ptr("Resource name suffix"),
	// 								DisplayName: to.Ptr("Suffix"),
	// 							},
	// 						},
	// 					},
	// 					PolicyRule: map[string]any{
	// 						"if": map[string]any{
	// 							"not": map[string]any{
	// 								"field": "name",
	// 								"like": "[concat(parameters('prefix'), '*', parameters('suffix'))]",
	// 							},
	// 						},
	// 						"then": map[string]any{
	// 							"effect": "deny",
	// 						},
	// 					},
	// 					PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
	// 					Version: to.Ptr("1.2.1"),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("1.0.0"),
	// 				Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
	// 				ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming/versions/1.0.0"),
	// 				Properties: &armpolicy.DefinitionVersionProperties{
	// 					Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
	// 					DisplayName: to.Ptr("Naming Convention"),
	// 					Metadata: map[string]any{
	// 						"category": "Naming",
	// 					},
	// 					Mode: to.Ptr("All"),
	// 					Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 						"prefix": &armpolicy.ParameterDefinitionsValue{
	// 							Type: to.Ptr(armpolicy.ParameterTypeString),
	// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 								Description: to.Ptr("Resource name prefix"),
	// 								DisplayName: to.Ptr("Prefix"),
	// 							},
	// 						},
	// 						"suffix": &armpolicy.ParameterDefinitionsValue{
	// 							Type: to.Ptr(armpolicy.ParameterTypeString),
	// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 								Description: to.Ptr("Resource name suffix"),
	// 								DisplayName: to.Ptr("Suffix"),
	// 							},
	// 						},
	// 					},
	// 					PolicyRule: map[string]any{
	// 						"if": map[string]any{
	// 							"not": map[string]any{
	// 								"field": "name",
	// 								"like": "[concat(parameters('prefix'), '-*', parameters('suffix'))]",
	// 							},
	// 						},
	// 						"then": map[string]any{
	// 							"effect": "deny",
	// 						},
	// 					},
	// 					PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
	// 					Version: to.Ptr("1.0.0"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
