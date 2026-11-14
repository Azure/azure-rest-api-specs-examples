package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-01/GetRaiPolicyAcs.json
func ExampleRaiPoliciesClient_Get_getAnAcsPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiPoliciesClient().Get(ctx, "resource-group", "safety-account", "agent-guard", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiPoliciesClientGetResponse{
	// 	RaiPolicy: armcognitiveservices.RaiPolicy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiPolicies/agent-guard"),
	// 		Name: to.Ptr("agent-guard"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiPolicies"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000003\""),
	// 		Properties: &armcognitiveservices.RaiPolicyProperties{
	// 			Type: to.Ptr(armcognitiveservices.RaiPolicyTypeUserManaged),
	// 			CustomBlocklists: []*armcognitiveservices.CustomBlocklistConfig{
	// 				{
	// 					BlocklistName: to.Ptr("blocked-terms"),
	// 					Source: to.Ptr(armcognitiveservices.RaiPolicyContentSourcePrompt),
	// 					Blocking: to.Ptr(true),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
