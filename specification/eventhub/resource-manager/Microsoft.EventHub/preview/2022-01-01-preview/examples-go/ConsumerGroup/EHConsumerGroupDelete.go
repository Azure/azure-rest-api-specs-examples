package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2f2b633d940230cbbd5bcf1339a2e1c48674e4a2/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/ConsumerGroup/EHConsumerGroupDelete.json
func ExampleConsumerGroupsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConsumerGroupsClient().Delete(ctx, "ArunMonocle", "sdk-Namespace-2661", "sdk-EventHub-6681", "sdk-ConsumerGroup-5563", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
