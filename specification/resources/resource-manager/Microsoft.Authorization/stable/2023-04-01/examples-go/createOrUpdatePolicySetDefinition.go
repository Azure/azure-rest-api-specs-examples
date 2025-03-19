package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createOrUpdatePolicySetDefinition.json
func ExampleSetDefinitionsClient_CreateOrUpdate_createOrUpdateAPolicySetDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSetDefinitionsClient().CreateOrUpdate(ctx, "CostManagement", armpolicy.SetDefinition{
		Properties: &armpolicy.SetDefinitionProperties{
			Description: to.Ptr("Policies to enforce low cost storage SKUs"),
			DisplayName: to.Ptr("Cost Management"),
			Metadata: map[string]any{
				"category": "Cost Management",
			},
			Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
				"namePrefix": {
					Type:         to.Ptr(armpolicy.ParameterTypeString),
					DefaultValue: "myPrefix",
					Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
						DisplayName: to.Ptr("Prefix to enforce on resource names"),
					},
				},
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
					PolicyDefinitionID:          to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
					PolicyDefinitionReferenceID: to.Ptr("Limit_Skus"),
				},
				{
					Parameters: map[string]*armpolicy.ParameterValuesValue{
						"prefix": {
							Value: "[parameters('namePrefix')]",
						},
						"suffix": {
							Value: "-LC",
						},
					},
					PolicyDefinitionID:          to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
					PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SetDefinition = armpolicy.SetDefinition{
	// 	Name: to.Ptr("CostManagement"),
	// 	Type: to.Ptr("Microsoft.Authorization/policySetDefinitions"),
	// 	ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
	// 	Properties: &armpolicy.SetDefinitionProperties{
	// 		Description: to.Ptr("Policies to enforce low cost storage SKUs"),
	// 		DisplayName: to.Ptr("Cost Management"),
	// 		Metadata: map[string]any{
	// 			"category": "Cost Management",
	// 		},
	// 		Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 			"namePrefix": &armpolicy.ParameterDefinitionsValue{
	// 				Type: to.Ptr(armpolicy.ParameterTypeString),
	// 				DefaultValue: "myPrefix",
	// 				Metadata: &armpolicy.ParameterDefinitionsValueMetadata{
	// 					DisplayName: to.Ptr("Prefix to enforce on resource names"),
	// 				},
	// 			},
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
	// 				PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
	// 				PolicyDefinitionReferenceID: to.Ptr("Limit_Skus"),
	// 			},
	// 			{
	// 				Parameters: map[string]*armpolicy.ParameterValuesValue{
	// 					"prefix": &armpolicy.ParameterValuesValue{
	// 						Value: "[parameters('namePrefix')]",
	// 					},
	// 					"suffix": &armpolicy.ParameterValuesValue{
	// 						Value: "-LC",
	// 					},
	// 				},
	// 				PolicyDefinitionID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
	// 				PolicyDefinitionReferenceID: to.Ptr("Resource_Naming"),
	// 		}},
	// 	},
	// }
}
