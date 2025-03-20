package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listPolicySetDefinitionVersionsByManagementGroup.json
func ExampleSetDefinitionVersionsClient_NewListByManagementGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSetDefinitionVersionsClient().NewListByManagementGroupPager("MyManagementGroup", "CostManagement", &armpolicy.SetDefinitionVersionsClientListByManagementGroupOptions{Expand: nil,
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
		// page.SetDefinitionVersionListResult = armpolicy.SetDefinitionVersionListResult{
		// 	Value: []*armpolicy.SetDefinitionVersion{
		// 		{
		// 			Name: to.Ptr("1.2.1"),
		// 			Type: to.Ptr("Microsoft.Authorization/policySetDefinitions/versions"),
		// 			ID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policySetDefinitions/CostManagement/versions/1.2.1"),
		// 			Properties: &armpolicy.SetDefinitionVersionProperties{
		// 				Description: to.Ptr("Policies to enforce low cost storage SKUs"),
		// 				DisplayName: to.Ptr("Cost Management"),
		// 				Metadata: map[string]any{
		// 					"category": "Cost Management",
		// 				},
		// 				PolicyDefinitions: []*armpolicy.DefinitionReference{
		// 					{
		// 						DefinitionVersion: to.Ptr("1.*.*"),
		// 						Parameters: map[string]*armpolicy.ParameterValuesValue{
		// 							"listOfAllowedSKUs": &armpolicy.ParameterValuesValue{
		// 								Value: []any{
		// 									"Standard_GRS",
		// 									"Standard_LRS",
		// 								},
		// 							},
		// 						},
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
		// 						PolicyDefinitionReferenceID: to.Ptr("Limit_Skus"),
		// 					},
		// 					{
		// 						DefinitionVersion: to.Ptr("1.*.*"),
		// 						Parameters: map[string]*armpolicy.ParameterValuesValue{
		// 							"prefix": &armpolicy.ParameterValuesValue{
		// 								Value: "DeptA",
		// 							},
		// 							"suffix": &armpolicy.ParameterValuesValue{
		// 								Value: "-LC",
		// 							},
		// 						},
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
		// 						PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
		// 				}},
		// 				Version: to.Ptr("1.2.1"),
		// 			},
		// 	}},
		// }
	}
}
