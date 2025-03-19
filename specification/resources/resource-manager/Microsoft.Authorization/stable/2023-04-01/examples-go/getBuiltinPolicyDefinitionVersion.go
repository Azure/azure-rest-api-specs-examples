package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/getBuiltinPolicyDefinitionVersion.json
func ExampleDefinitionVersionsClient_GetBuiltIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefinitionVersionsClient().GetBuiltIn(ctx, "7433c107-6db4-4ad1-b57a-a76dce0154a1", "1.2.1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DefinitionVersion = armpolicy.DefinitionVersion{
	// 	Name: to.Ptr("1.2.1"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyDefinitions/versions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1/versions/1.2.1"),
	// 	Properties: &armpolicy.DefinitionVersionProperties{
	// 		Description: to.Ptr("This policy enables you to specify a set of storage account SKUs that your organization can deploy."),
	// 		DisplayName: to.Ptr("Allowed storage account SKUs"),
	// 		Mode: to.Ptr("All"),
	// 		Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 			"listOfAllowedSKUs": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeArray),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("The list of SKUs that can be specified for storage accounts."),
	// 					DisplayName: to.Ptr("Allowed SKUs"),
	// 					StrongType: to.Ptr("StorageSKUs"),
	// 				},
	// 			},
	// 		},
	// 		PolicyRule: map[string]any{
	// 			"if":map[string]any{
	// 				"allOf":[]any{
	// 					map[string]any{
	// 						"equals": "Microsoft.Storage/storageAccounts",
	// 						"field": "type",
	// 					},
	// 					map[string]any{
	// 						"not":map[string]any{
	// 							"field": "Microsoft.Storage/storageAccounts/sku.name",
	// 							"in": "[parameters('listOfAllowedSKUs')]",
	// 						},
	// 					},
	// 				},
	// 			},
	// 			"then":map[string]any{
	// 				"effect": "Deny",
	// 			},
	// 		},
	// 		PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
	// 		Version: to.Ptr("1.2.1"),
	// 	},
	// }
}
