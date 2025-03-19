package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listPolicyDefinitionVersionsByManagementGroup.json
func ExampleDefinitionVersionsClient_NewListByManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDefinitionVersionsClient().NewListByManagementGroupPager("MyManagementGroup", "ResourceNaming", &armpolicy.DefinitionVersionsClientListByManagementGroupOptions{Top: nil})
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
		// page.DefinitionVersionListResult = armpolicy.DefinitionVersionListResult{
		// 	Value: []*armpolicy.DefinitionVersion{
		// 		{
		// 			Name: to.Ptr("1.2.1"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions/versions"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming/versions/1.2.1"),
		// 			Properties: &armpolicy.DefinitionVersionProperties{
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
		// 				Version: to.Ptr("1.2.1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1.0.0"),
		// 			Type: to.Ptr("Microsoft.Authorization/policyDefinitions/versions"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming/versions/1.0.0"),
		// 			Properties: &armpolicy.DefinitionVersionProperties{
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
		// 							"like": "[concat(parameters('prefix'), '-*', parameters('suffix'))]",
		// 						},
		// 					},
		// 					"then":map[string]any{
		// 						"effect": "deny",
		// 					},
		// 				},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeCustom),
		// 				Version: to.Ptr("1.2.1"),
		// 			},
		// 	}},
		// }
	}
}
