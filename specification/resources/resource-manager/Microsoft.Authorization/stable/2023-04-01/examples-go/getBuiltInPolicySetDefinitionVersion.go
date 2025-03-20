package armpolicy_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/219b2e3ef270f18149774eb2793b48baacde982f/specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/getBuiltInPolicySetDefinitionVersion.json
func ExampleSetDefinitionVersionsClient_GetBuiltIn() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpolicy.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSetDefinitionVersionsClient().GetBuiltIn(ctx, "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", "1.2.1", &armpolicy.SetDefinitionVersionsClientGetBuiltInOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SetDefinitionVersion = armpolicy.SetDefinitionVersion{
	// 	Name: to.Ptr("1.2.1"),
	// 	Type: to.Ptr("Microsoft.Authorization/policySetDefinitions/versions"),
	// 	ID: to.Ptr("/providers/Microsoft.Authorization/policySetDefinitions/1f3afdf9-d0c9-4c3d-847f-89da613e70a8/versions/1.2.1"),
	// 	Properties: &armpolicy.SetDefinitionVersionProperties{
	// 		Description: to.Ptr("Monitor all the available security recommendations in Azure Security Center. This is the default policy for Azure Security Center."),
	// 		DisplayName: to.Ptr("[Preview]: Enable Monitoring in Azure Security Center"),
	// 		Metadata: map[string]any{
	// 			"category": "Security Center",
	// 		},
	// 		Parameters: map[string]*armpolicy.ParameterDefinitionsValue{
	// 		},
	// 		PolicyDefinitions: []*armpolicy.DefinitionReference{
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/a8bef009-a5c9-4d0f-90d7-6018734e8a16"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId1"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/af8051bf-258b-44e2-a2bf-165330459f9d"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId2"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/86b3d65f-7626-441e-b690-81a8b71cff60"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId3"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/655cb504-bcee-4362-bd4c-402e6aa38759"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId4"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/b0f33259-77d7-4c9e-aac6-3aabcfae693c"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId5"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/47a6b606-51aa-4496-8bb7-64b11cf66adc"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId6"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/44452482-524f-4bf4-b852-0bff7cc4a3ed"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId7"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/e1e5fd5d-3e4c-4ce1-8661-7d1873ae6b15"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId8"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/af6cd1bd-1635-48cb-bde7-5b15693900b9"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId9"),
	// 			},
	// 			{
	// 				DefinitionVersion: to.Ptr("1.*.*"),
	// 				PolicyDefinitionID: to.Ptr("/providers/Microsoft.Authorization/policyDefinitions/0961003e-5a0a-4549-abde-af6a37f2724d"),
	// 				PolicyDefinitionReferenceID: to.Ptr("RefId10"),
	// 		}},
	// 		PolicyType: to.Ptr(armpolicy.PolicyTypeBuiltIn),
	// 		Version: to.Ptr("1.2.1"),
	// 	},
	// }
}
