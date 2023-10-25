package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/getPolicySetDefinitionAtManagementGroup.json
func ExampleSetDefinitionsClient_GetAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSetDefinitionsClient().GetAtManagementGroup(ctx, "CostManagement", "MyManagementGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SetDefinition = armpolicy.SetDefinition{
	// 	Name: to.Ptr("CostManagement"),
	// 	Type: to.Ptr("Microsoft.Authorization/policySetDefinitions"),
	// 	ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
	// 	Properties: &armpolicy.SetDefinitionProperties{
	// 		Description: to.Ptr("Policies to enforce low cost storage SKUs"),
	// 		DisplayName: to.Ptr("Cost Management"),
	// 		Metadata: map[string]any{
	// 			"category": "Cost Management",
	// 		},
	// 		PolicyDefinitions: []*armpolicy.DefinitionReference{
	// 			{
	// 				Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 					"listOfAllowedSKUs": &armpolicy.ParameterValuesValue{
	// 						Value: []any{
	// 							"Standard_GRS",
	// 							"Standard_LRS",
	// 						},
	// 					},
	// 				},
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
	// 				PolicyDefinitionReferenceID: to.Ptr("Limit_Skus"),
	// 			},
	// 			{
	// 				Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 					"prefix": &armpolicy.ParameterValuesValue{
	// 						Value: "DeptA",
	// 					},
	// 					"suffix": &armpolicy.ParameterValuesValue{
	// 						Value: "-LC",
	// 					},
	// 				},
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 				PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
	// 		}},
	// 	},
	// }
}
