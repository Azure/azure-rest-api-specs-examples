package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6bbb6daca7175b2cab69b20b2dd01094e3c6a534/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/PutRaiTopic.json
func ExampleRaiTopicsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRaiTopicsClient().CreateOrUpdate(ctx, "resourceGroupName", "accountName", "raiTopicName", armcognitiveservices.RaiTopic{
		Properties: &armcognitiveservices.RaiTopicProperties{
			Description:   to.Ptr("This is a sample topic."),
			SampleBlobURL: to.Ptr("https://example.blob.core.windows.net/sampleblob"),
			TopicName:     to.Ptr("raiTopicName"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RaiTopic = armcognitiveservices.RaiTopic{
	// 	Name: to.Ptr("raiTopicName"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/raiTopics/raiTopicName"),
	// 	Properties: &armcognitiveservices.RaiTopicProperties{
	// 		Description: to.Ptr("This is a sample topic."),
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-01T00:00:00.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-07-02T00:00:00.000Z"); return t}()),
	// 		SampleBlobURL: to.Ptr("https://example.blob.core.windows.net/sampleblob"),
	// 		Status: to.Ptr("Training"),
	// 		TopicID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TopicName: to.Ptr("raiTopicName"),
	// 	},
	// }
}
