package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/ConsumerGroup/EHConsumerGroupDelete.json
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
