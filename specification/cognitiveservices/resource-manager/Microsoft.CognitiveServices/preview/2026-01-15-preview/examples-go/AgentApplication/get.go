package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-01-15-preview/AgentApplication/get.json
func ExampleAgentApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentApplicationsClient().Get(ctx, "test-rg", "my-cognitive-services-account", "my-project", "agent-app-1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AgentApplicationsClientGetResponse{
	// 	AgentApplication: &armcognitiveservices.AgentApplication{
	// 		Name: to.Ptr("agent-app-1"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/applications"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/test-rg/providers/Microsoft.CognitiveServices/accounts/my-cognitive-services-account/projects/my-project/applications/agent-app-1"),
	// 		Properties: &armcognitiveservices.AgenticApplicationProperties{
	// 			Description: to.Ptr("Sample agent application for customer support"),
	// 			DisplayName: to.Ptr("Customer Support Agent"),
	// 			Tags: map[string]*string{
	// 				"environment": to.Ptr("production"),
	// 				"team": to.Ptr("ai-platform"),
	// 			},
	// 		},
	// 		SystemData: &armcognitiveservices.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-29T12:34:56.999Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-29T12:34:56.999Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armcognitiveservices.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
