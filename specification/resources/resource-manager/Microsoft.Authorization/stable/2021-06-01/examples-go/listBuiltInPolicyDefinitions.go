package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listBuiltInPolicyDefinitions.json
func ExampleDefinitionsClient_NewListBuiltInPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefinitionsClient().NewListBuiltInPager(&armpolicy.DefinitionsClientListBuiltInOptions{Filter: nil,
		Top: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DefinitionListResult = armpolicy.DefinitionListResult{
		// 	Value: []*armpolicy.Definition{
		// 		{
		// 			Name: to.Ptr("06a78e20-9358-41c9-923c-fb736d382a12"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/06a78e20-9358-41c9-923c-fb736d382a12"),
		// 			Properties: &armpolicy.DefinitionProperties{
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
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 				Properties: &armpolicy.DefinitionProperties{
		// 					Description: to.Ptr("This policy enables you to specify a set of storage account SKUs that your organization can deploy."),
		// 					DisplayName: to.Ptr("Allowed storage account SKUs"),
		// 					Mode: to.Ptr("All"),
		// 					Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 						"listOfAllowedSKUs": &armpolicy.ParameterDefinitionsValue{
		// 							Type: to.Ptr(armpolicy.ParameterTypeArray),
		// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 								Description: to.Ptr("The list of SKUs that can be specified for storage accounts."),
		// 								DisplayName: to.Ptr("Allowed SKUs"),
		// 								StrongType: to.Ptr("StorageSKUs"),
		// 							},
		// 						},
		// 					},
		// 					PolicyRule: map[string]any{
		// 						"if":map[string]any{
		// 							"allOf":[]any{
		// 								map[string]any{
		// 									"equals": "Microsoft.Storage/storageAccounts",
		// 									"field": "type",
		// 								},
		// 								map[string]any{
		// 									"not":map[string]any{
		// 										"field": "Microsoft.Storage/storageAccounts/sku.name",
		// 										"in": "[parameters('listOfAllowedSKUs')]",
		// 									},
		// 								},
		// 							},
		// 						},
		// 						"then":map[string]any{
		// 							"effect": "Deny",
		// 						},
		// 					},
		// 					PolicyType: to.Ptr(armpolicy.PolicyTypeStatic),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("abeed54a-73c5-441d-8a8c-6b5e7a0c299e"),
		// 				Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/abeed54a-73c5-441d-8a8c-6b5e7a0c299e"),
		// 				Properties: &armpolicy.DefinitionProperties{
		// 					Description: to.Ptr("Audit certificates that are stored in Azure Key Vault, that expire within 'X' number of days."),
		// 					DisplayName: to.Ptr("Audit KeyVault certificates that expire within specified number of days"),
		// 					Metadata: map[string]any{
		// 						"category": "KeyVault DataPlane",
		// 					},
		// 					Mode: to.Ptr("Microsoft.KeyVault.Data"),
		// 					Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 						"daysToExpire": &armpolicy.ParameterDefinitionsValue{
		// 							Type: to.Ptr(armpolicy.ParameterTypeInteger),
		// 							Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 								Description: to.Ptr("The number of days for a certificate to expire."),
		// 								DisplayName: to.Ptr("Days to expire"),
		// 							},
		// 						},
		// 					},
		// 					PolicyRule: map[string]any{
		// 						"if":map[string]any{
		// 							"field": "Microsoft.KeyVault.Data/vaults/certificates/attributes/expiresOn",
		// 							"lessOrEquals": "[addDays(utcNow(), parameters('daysToExpire'))]",
		// 						},
		// 						"then":map[string]any{
		// 							"effect": "audit",
		// 						},
		// 					},
		// 					PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
		// 				},
		// 		}},
		// 	}
	}
}
