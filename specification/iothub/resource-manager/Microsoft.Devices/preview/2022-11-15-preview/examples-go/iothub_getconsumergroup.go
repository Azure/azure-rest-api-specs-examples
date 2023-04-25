package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1f22d4dbd99b0fe347ad79e79d4eb1ed44a87291/specification/iothub/resource-manager/Microsoft.Devices/preview/2022-11-15-preview/examples/iothub_getconsumergroup.json
func ExampleResourceClient_GetEventHubConsumerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().GetEventHubConsumerGroup(ctx, "myResourceGroup", "testHub", "events", "test", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventHubConsumerGroupInfo = armiothub.EventHubConsumerGroupInfo{
	// 	Name: to.Ptr("test"),
	// 	Type: to.Ptr("Microsoft.Devices/IotHubs/EventHubEndpoints/ConsumerGroups"),
	// 	Etag: to.Ptr("AAAAAAFD6M4="),
	// 	ID: to.Ptr("/subscriptions/cmd-sub-1/resourceGroups/cmd-rg-1/providers/Microsoft.Devices/IotHubs/test-hub-2/eventHubEndpoints/events/ConsumerGroups/%24Default"),
	// 	Properties: map[string]any{
	// 		"created": "Thu, 15 Jun 2017 19:20:58 GMT",
	// 	},
	// }
}
