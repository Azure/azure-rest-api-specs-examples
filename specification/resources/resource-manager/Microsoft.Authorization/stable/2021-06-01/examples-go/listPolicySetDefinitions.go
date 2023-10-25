package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/resources/resource-manager/Microsoft.Authorization/stable/2021-06-01/examples/listPolicySetDefinitions.json
func ExampleSetDefinitionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSetDefinitionsClient().NewListPager(&armpolicy.SetDefinitionsClientListOptions{Filter: nil,
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
		// page.SetDefinitionListResult = armpolicy.SetDefinitionListResult{
		// 	Value: []*armpolicy.SetDefinition{
		// 		{
		// 			Name: to.Ptr("1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 			Type: to.Ptr("Microsoft.Authorization/policySetDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Authorization/policySetDefinitions/1f3afdf9-d0c9-4c3d-847f-89da613e70a8"),
		// 			Properties: &armpolicy.SetDefinitionProperties{
		// 				Description: to.Ptr("Monitor all the available security recommendations in Azure Security Center. This is the default policy for Azure Security Center."),
		// 				DisplayName: to.Ptr("[Preview]: Enable Monitoring in Azure Security Center"),
		// 				Metadata: map[string]any{
		// 					"category": "Security Center",
		// 				},
		// 				Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
		// 				},
		// 				PolicyDefinitions: []*armpolicy.DefinitionReference{
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/a8bef009-a5c9-4d0f-90d7-6018734e8a16"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId1"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/af8051bf-258b-44e2-a2bf-165330459f9d"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId2"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/86b3d65f-7626-441e-b690-81a8b71cff60"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId3"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/655cb504-bcee-4362-bd4c-402e6aa38759"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId4"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/b0f33259-77d7-4c9e-aac6-3aabcfae693c"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId5"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/47a6b606-51aa-4496-8bb7-64b11cf66adc"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId6"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/44452482-524f-4bf4-b852-0bff7cc4a3ed"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId7"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/e1e5fd5d-3e4c-4ce1-8661-7d1873ae6b15"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId8"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/af6cd1bd-1635-48cb-bde7-5b15693900b9"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId9"),
		// 					},
		// 					{
		// 						PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/0961003e-5a0a-4549-abde-af6a37f2724d"),
		// 						PolicyDefinitionReferenceID: to.Ptr("RefId10"),
		// 				}},
		// 				PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CostManagement"),
		// 			Type: to.Ptr("Microsoft.Authorization/policySetDefinitions"),
		// 			ID: to.Ptr("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
		// 			Properties: &armpolicy.SetDefinitionProperties{
		// 				Description: to.Ptr("Policies to enforce low cost storage SKUs"),
		// 				DisplayName: to.Ptr("Cost Management"),
		// 				Metadata: map[string]any{
		// 					"category": "Cost Management",
		// 				},
		// 				PolicyDefinitions: []*armpolicy.DefinitionReference{
		// 					{
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
		// 			},
		// 	}},
		// }
	}
}
