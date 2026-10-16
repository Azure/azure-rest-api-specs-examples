package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/PutRaiPolicyAcsWithoutTools.json
func ExampleRaiPoliciesClient_CreateOrUpdate_createAnAcsPolicyWithoutAToolCatalog() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiPoliciesClient().CreateOrUpdate(ctx, "resource-group", "safety-account", "input-guard", armcognitiveservices.RaiPolicy{
		Properties: &armcognitiveservices.RaiPolicyProperties{
			Format: to.Ptr(armcognitiveservices.RaiPolicyFormatACS),
			Acs: &armcognitiveservices.RaiAcsManifest{
				AgentControlSpecificationVersion: to.Ptr("0.4.0-alpha.1"),
				Policies: map[string]*armcognitiveservices.RaiAcsRegoPolicyDefinition{
					"input-guard": {
						Type:  to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
						Query: to.Ptr("data.input_guard.verdict"),
					},
				},
				InterventionPoints: &armcognitiveservices.RaiAcsInterventionPoints{
					Input: &armcognitiveservices.RaiAcsInterventionPoint{
						PolicyTarget:     to.Ptr(armcognitiveservices.RaiAcsPolicyTargetInput),
						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindUserInput),
						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
							ID: to.Ptr("input-guard"),
						},
					},
				},
			},
			AcsRegos: []*armcognitiveservices.RaiRegoReference{
				{
					RegoName: to.Ptr("input-guard"),
				},
			},
		},
	}, &armcognitiveservices.RaiPoliciesClientCreateOrUpdateOptions{
		IfNoneMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiPoliciesClientCreateOrUpdateResponse{
	// 	EtagHeader: to.Ptr("\"00000000-0000-0000-0000-000000000005\""),
	// 	RaiPolicy: armcognitiveservices.RaiPolicy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiPolicies/input-guard"),
	// 		Name: to.Ptr("input-guard"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiPolicies"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000005\""),
	// 		Properties: &armcognitiveservices.RaiPolicyProperties{
	// 			Type: to.Ptr(armcognitiveservices.RaiPolicyTypeUserManaged),
	// 			Format: to.Ptr(armcognitiveservices.RaiPolicyFormatACS),
	// 			Acs: &armcognitiveservices.RaiAcsManifest{
	// 				AgentControlSpecificationVersion: to.Ptr("0.4.0-alpha.1"),
	// 				Policies: map[string]*armcognitiveservices.RaiAcsRegoPolicyDefinition{
	// 					"input-guard": &armcognitiveservices.RaiAcsRegoPolicyDefinition{
	// 						Type: to.Ptr(armcognitiveservices.RaiAcsPolicyDefinitionTypeRego),
	// 						Query: to.Ptr("data.input_guard.verdict"),
	// 					},
	// 				},
	// 				InterventionPoints: &armcognitiveservices.RaiAcsInterventionPoints{
	// 					Input: &armcognitiveservices.RaiAcsInterventionPoint{
	// 						PolicyTarget: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetInput),
	// 						PolicyTargetKind: to.Ptr(armcognitiveservices.RaiAcsPolicyTargetKindUserInput),
	// 						Policy: &armcognitiveservices.RaiAcsPolicyBinding{
	// 							ID: to.Ptr("input-guard"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			AcsRegos: []*armcognitiveservices.RaiRegoReference{
	// 				{
	// 					RegoName: to.Ptr("input-guard"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
