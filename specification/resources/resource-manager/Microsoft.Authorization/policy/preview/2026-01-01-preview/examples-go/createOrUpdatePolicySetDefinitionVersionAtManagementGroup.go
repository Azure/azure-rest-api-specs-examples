package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy/v2"
)

// Generated from example definition: 2026-01-01-preview/createOrUpdatePolicySetDefinitionVersionAtManagementGroup.json
func ExampleSetDefinitionVersionsClient_CreateOrUpdateAtManagementGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSetDefinitionVersionsClient().CreateOrUpdateAtManagementGroup(ctx, "MyManagementGroup", "CostManagement", "1.2.1", armpolicy.SetDefinitionVersion{
		Properties: &armpolicy.SetDefinitionVersionProperties{
			Description: to.Ptr("Policies to enforce low cost storage SKUs"),
			DisplayName: to.Ptr("Cost Management"),
			Metadata: map[string]any{
				"category": "Cost Management",
			},
			PolicyDefinitions: []*armpolicy.DefinitionReference{
				{
					Parameters: map[string]*armpolicy.ParameterValuesValue{
						"listOfAllowedSKUs": {
							Value: []any{
								"Standard_GRS",
								"Standard_LRS",
							},
						},
					},
					PolicyDefinitionID:          to.Ptr("/providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
					PolicyDefinitionReferenceID: to.Ptr("Limit_Skus"),
				},
				{
					Parameters: map[string]*armpolicy.ParameterValuesValue{
						"prefix": {
							Value: "DeptA",
						},
						"suffix": {
							Value: "-LC",
						},
					},
					PolicyDefinitionID:          to.Ptr("/providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
					PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
				},
			},
			Version: to.Ptr("1.2.1"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpolicy.SetDefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse{
	// 	SetDefinitionVersion: armpolicy.SetDefinitionVersion{
	// 		Name: to.Ptr("CostManagement"),
	// 		Type: to.Ptr("Microsoft.Authorization/policySetDefinitions"),
	// 		ID: to.Ptr("/providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
	// 		Properties: &armpolicy.SetDefinitionVersionProperties{
	// 			Description: to.Ptr("Policies to enforce low cost storage SKUs"),
	// 			DisplayName: to.Ptr("Cost Management"),
	// 			Metadata: map[string]any{
	// 				"category": "Cost Management",
	// 			},
	// 			PolicyDefinitions: []*armpolicy.DefinitionReference{
	// 				{
	// 					DefinitionVersion: to.Ptr("1.*.*"),
	// 					Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 						"listOfAllowedSKUs": &armpolicy.ParameterValuesValue{
	// 							Value: []any{
	// 								"Standard_GRS",
	// 								"Standard_LRS",
	// 							},
	// 						},
	// 					},
	// 					PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
	// 					PolicyDefinitionReferenceID: to.Ptr("Limit_Skus"),
	// 				},
	// 				{
	// 					DefinitionVersion: to.Ptr("1.*.*"),
	// 					Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 						"prefix": &armpolicy.ParameterValuesValue{
	// 							Value: "DeptA",
	// 						},
	// 						"suffix": &armpolicy.ParameterValuesValue{
	// 							Value: "-LC",
	// 						},
	// 					},
	// 					PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 					PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
