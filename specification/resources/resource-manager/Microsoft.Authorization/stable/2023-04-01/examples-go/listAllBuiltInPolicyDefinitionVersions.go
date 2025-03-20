package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listAllBuiltInPolicyDefinitionVersions.json
func ExampleDefinitionVersionsClient_ListAllBuiltins() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefinitionVersionsClient().ListAllBuiltins(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefinitionVersionListResult = armpolicy.DefinitionVersionListResult{
	// 	Value: []*armpolicy.DefinitionVersion{
	// 		{
	// 			Name: to.Ptr("1.2.1"),
	// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions/versions"),
	// 			ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/06a78e20-9358-41c9-923c-fb736d382a12/versions/1.2.1"),
	// 			Properties: &armpolicy.DefinitionVersionProperties{
	// 				Description: to.Ptr("Audit DB level audit setting for SQL databases"),
	// 				DisplayName: to.Ptr("Audit SQL DB Level Audit Setting"),
	// 				Mode: to.Ptr("All"),
	// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 					"setting": &armpolicy.ParameterDefinitionsValue{
	// 						Type: to.Ptr(armpolicy.ParameterTypeString),
	// 						AllowedValues: []any{
	// 							"enabled",
	// 							"disabled"},
	// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 								DisplayName: to.Ptr("Audit Setting"),
	// 							},
	// 						},
	// 					},
	// 					PolicyRule: map[string]any{
	// 						"if":map[string]any{
	// 							"equals": "Microsoft.Sql/servers/databases",
	// 							"field": "type",
	// 						},
	// 						"then":map[string]any{
	// 							"effect": "AuditIfNotExists",
	// 							"details":map[string]any{
	// 								"name": "default",
	// 								"type": "Microsoft.Sql/servers/databases/auditingSettings",
	// 								"existenceCondition":map[string]any{
	// 									"allOf":[]any{
	// 										map[string]any{
	// 											"equals": "[parameters('setting')]",
	// 											"field": "Microsoft.Sql/auditingSettings.state",
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 					PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
	// 					Version: to.Ptr("1.2.1"),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("1.0.0"),
	// 				Type: to.Ptr("Microsoft.Authorization/policyDefinitions/versions"),
	// 				ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/06a78e20-9358-41c9-923c-fb736d382a12/versions/1.0.0"),
	// 				Properties: &armpolicy.DefinitionVersionProperties{
	// 					Description: to.Ptr("Audit DB level audit setting for SQL databases"),
	// 					DisplayName: to.Ptr("Audit SQL DB Level Audit Setting"),
	// 					Mode: to.Ptr("All"),
	// 					Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 						"setting": &armpolicy.ParameterDefinitionsValue{
	// 							Type: to.Ptr(armpolicy.ParameterTypeString),
	// 							AllowedValues: []any{
	// 								"enabled",
	// 								"disabled",
	// 								"default"},
	// 								Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 									DisplayName: to.Ptr("Audit Setting"),
	// 								},
	// 							},
	// 						},
	// 						PolicyRule: map[string]any{
	// 							"if":map[string]any{
	// 								"equals": "Microsoft.Sql/servers/databases",
	// 								"field": "type",
	// 							},
	// 							"then":map[string]any{
	// 								"effect": "AuditIfNotExists",
	// 								"details":map[string]any{
	// 									"name": "default",
	// 									"type": "Microsoft.Sql/servers/databases/auditingSettings",
	// 									"existenceCondition":map[string]any{
	// 										"allOf":[]any{
	// 											map[string]any{
	// 												"equals": "[parameters('setting')]",
	// 												"field": "Microsoft.Sql/auditingSettings.state",
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 						PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
	// 						Version: to.Ptr("1.0.0"),
	// 					},
	// 			}},
	// 		}
}
