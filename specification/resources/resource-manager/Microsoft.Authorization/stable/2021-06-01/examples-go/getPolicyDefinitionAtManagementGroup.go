package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicyDefinitionAtManagementGroup.json
func ExampleDefinitionsClient_GetAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDefinitionsClient().GetAtManagementGroup(ctx, "ResourceNaming", "MyManagementGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Definition = armpolicy.Definition{
	// 	Name: to.Ptr("ResourceNaming"),
	// 	Type: to.Ptr("Microsoft.Authorization/policyDefinitions"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 	Properties: &armpolicy.DefinitionProperties{
	// 		Description: to.Ptr("Force resource names to begin with 'prefix' and end with 'suffix'"),
	// 		DisplayName: to.Ptr("Naming Convention"),
	// 		Metadata: map[string]any{
	// 			"category": "Naming",
	// 		},
	// 		Mode: to.Ptr("All"),
	// 		Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 			"prefix": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeString),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("Resource name prefix"),
	// 					DisplayName: to.Ptr("Prefix"),
	// 				},
	// 			},
	// 			"suffix": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeString),
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					Description: to.Ptr("Resource name suffix"),
	// 					DisplayName: to.Ptr("Suffix"),
	// 				},
	// 			},
	// 		},
	// 		PolicyRule: map[string]any{
	// 			"if":map[string]any{
	// 				"not":map[string]any{
	// 					"field": "name",
	// 					"like": "[concat(parameters('prefix'), '*', parameters('suffix'))]",
	// 				},
	// 			},
	// 			"then":map[string]any{
	// 				"effect": "deny",
	// 			},
	// 		},
	// 		PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
	// 	},
	// }
}
