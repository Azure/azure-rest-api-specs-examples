package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicyDefinitionsByManagementGroup.json
func ExampleDefinitionsClient_NewListByManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefinitionsClient().NewListByManagementGroupPager("MyManagementGroup", &armpolicy.DefinitionsClientListByManagementGroupOptions{Filter: nil,
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
		// 			Name: to.Ptr("7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("This policy enables you to specify a set of storage account SKUs that your organization can deploy."),
		// 				DisplayName: to.Ptr("Allowed storage account SKUs"),
		// 				Mode: to.Ptr("All"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"listOfAllowedSKUs": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeArray),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("The list of SKUs that can be specified for storage accounts."),
		// 							DisplayName: to.Ptr("Allowed SKUs"),
		// 							StrongType: to.Ptr("StorageSKUs"),
		// 						},
		// 					},
		// 				},
		// 				PolicyRule: map[string]any{
		// 					"if":map[string]any{
		// 						"allOf":[]any{
		// 							map[string]any{
		// 								"equals": "Microsoft.Storage/storageAccounts",
		// 								"field": "type",
		// 							},
		// 							map[string]any{
		// 								"not":map[string]any{
		// 									"field": "Microsoft.Storage/storageAccounts/sku.name",
		// 									"in": "[parameters('listOfAllowedSKUs')]",
		// 								},
		// 							},
		// 						},
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "Deny",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ResourceNaming"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
		// 			Properties: &armpolicy.DefinitionProperties{
		// 				Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
		// 				DisplayName: to.Ptr("Naming Convention"),
		// 				Metadata: map[string]any{
		// 					"category": "Naming",
		// 				},
		// 				Mode: to.Ptr("All"),
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 					"prefix": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeString),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("Resource name prefix"),
		// 							DisplayName: to.Ptr("Prefix"),
		// 						},
		// 					},
		// 					"suffix": &armpolicy.ParameterDefinitionsValue{
		// 						Type: to.Ptr(armpolicy.ParameterTypeString),
		// 						Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
		// 							Description: to.Ptr("Resource name suffix"),
		// 							DisplayName: to.Ptr("Suffix"),
		// 						},
		// 					},
		// 				},
		// 				PolicyRule: map[string]any{
		// 					"if":map[string]any{
		// 						"not":map[string]any{
		// 							"field": "name",
		// 							"like": "[concat(parameters('prefix'), '*', parameters('suffix'))]",
		// 						},
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "deny",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
		// 			},
		// 	}},
		// }
	}
}
