package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_createconsumergroup.json
func ExampleResourceClient_CreateEventHubConsumerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armiothub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().CreateEventHubConsumerGroup(ctx, "myResourceGroup", "testHub", "events", "test", armiothub.EventHubConsumerGroupBodyDescription{
		Properties: &armiothub.EventHubConsumerGroupName{
			Name: to.Ptr("test"),
		},
	}, nil)
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
