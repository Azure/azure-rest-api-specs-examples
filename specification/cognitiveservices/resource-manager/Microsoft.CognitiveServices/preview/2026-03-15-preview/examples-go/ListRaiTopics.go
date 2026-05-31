package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-03-15-preview/ListRaiTopics.json
func ExampleRaiTopicsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRaiTopicsClient().NewListPager("resourceGroupName", "accountName", nil)
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
		// page = armcognitiveservices.RaiTopicsClientListResponse{
		// 	RaiTopicResult: armcognitiveservices.RaiTopicResult{
		// 		Value: []*armcognitiveservices.RaiTopic{
		// 			{
		// 				Name: to.Ptr("raiTopicName"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiTopics/raiTopicName"),
		// 				Properties: &armcognitiveservices.RaiTopicProperties{
		// 					Description: to.Ptr("This is a sample topic."),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-01T00:00:00Z"); return t}()),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-02T00:00:00Z"); return t}()),
		// 					SampleBlobURL: to.Ptr("https://example.blob.core.windows.net/sampleblob"),
		// 					Status: to.Ptr("Succeeded"),
		// 					TopicID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 					TopicName: to.Ptr("raiTopicName"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
