package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-06-01/StorageContextCacheContainerCRUD/ContextCacheContainers_ListByContextCache.json
func ExampleContextCacheContainersClient_NewListByContextCachePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContextCacheContainersClient().NewListByContextCachePager("testrg", "testaccount", nil)
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
		// page = armstorage.ContextCacheContainersClientListByContextCacheResponse{
		// 	ContextCacheContainerListResult: armstorage.ContextCacheContainerListResult{
		// 		Value: []*armstorage.ContextCacheContainer{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/contextCaches/testaccount/contextCacheContainers/gpt4-prompts"),
		// 				Name: to.Ptr("gpt4-prompts"),
		// 				Type: to.Ptr("Microsoft.Storage/contextCaches/contextCacheContainers"),
		// 				Properties: &armstorage.ContextCacheContainerProperties{
		// 					Description: to.Ptr("Container for GPT-4 prompt caching"),
		// 					ModelName: to.Ptr("gpt-4"),
		// 					Provider: to.Ptr(armstorage.AiProviderOpenAI),
		// 					ProvisioningState: to.Ptr(armstorage.ContextCacheProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/contextCaches/testaccount/contextCacheContainers/claude3-prompts"),
		// 				Name: to.Ptr("claude3-prompts"),
		// 				Type: to.Ptr("Microsoft.Storage/contextCaches/contextCacheContainers"),
		// 				Properties: &armstorage.ContextCacheContainerProperties{
		// 					Description: to.Ptr("Container for Claude-3 prompt caching"),
		// 					ModelName: to.Ptr("claude-3"),
		// 					Provider: to.Ptr(armstorage.AiProvider("Anthropic")),
		// 					ProvisioningState: to.Ptr(armstorage.ContextCacheProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
