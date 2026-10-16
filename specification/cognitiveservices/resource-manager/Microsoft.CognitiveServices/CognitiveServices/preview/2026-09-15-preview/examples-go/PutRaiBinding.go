package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/PutRaiBinding.json
func ExampleRaiBindingsClient_CreateOrUpdate_bindAnAzureOpenAiDeploymentToAnAcsPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiBindingsClient().CreateOrUpdate(ctx, "resource-group", "safety-account", "chat-guard", armcognitiveservices.RaiBinding{
		Properties: &armcognitiveservices.RaiBindingProperties{
			BoundResourceID:  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/aoai-account/deployments/chat-prod"),
			TargetPolicyName: to.Ptr("agent-guard"),
		},
	}, &armcognitiveservices.RaiBindingsClientCreateOrUpdateOptions{
		IfNoneMatch: to.Ptr("*")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.RaiBindingsClientCreateOrUpdateResponse{
	// 	EtagHeader: to.Ptr("\"00000000-0000-0000-0000-000000000002\""),
	// 	RaiBinding: armcognitiveservices.RaiBinding{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/safety-account/raiBindings/chat-guard"),
	// 		Name: to.Ptr("chat-guard"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/raiBindings"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000002\""),
	// 		Properties: &armcognitiveservices.RaiBindingProperties{
	// 			BoundResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resource-group/providers/Microsoft.CognitiveServices/accounts/aoai-account/deployments/chat-prod"),
	// 			TargetPolicyName: to.Ptr("agent-guard"),
	// 		},
	// 	},
	// }
}
