package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/listAllPolicySetDefinitionVersions.json
func ExampleSetDefinitionVersionsClient_ListAll() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSetDefinitionVersionsClient().ListAll(ctx, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SetDefinitionVersionListResult = armpolicy.SetDefinitionVersionListResult{
	// 	Value: []*armpolicy.SetDefinitionVersion{
	// 		{
	// 			Name: to.Ptr("1.2.1"),
	// 			Type: to.Ptr("Microsoft.Authorization/policySetDefinitions/versions"),
	// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement/versions/1.2.1"),
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
	// 						PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
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
	// 						PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 						PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
	// 				}},
	// 				Version: to.Ptr("1.2.1"),
	// 			},
	// 	}},
	// }
}
