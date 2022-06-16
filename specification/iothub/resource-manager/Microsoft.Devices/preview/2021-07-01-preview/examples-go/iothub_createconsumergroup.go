package armiothub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iothub/armiothub"
)

// x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2021-07-01-preview/examples/iothub_createconsumergroup.json
func ExampleResourceClient_CreateEventHubConsumerGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armiothub.NewResourceClient("<subscription-id>", cred, nil)
	res, err := client.CreateEventHubConsumerGroup(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<event-hub-endpoint-name>",
		"<name>",
		armiothub.EventHubConsumerGroupBodyDescription{
			Properties: &armiothub.EventHubConsumerGroupName{
				Name: to.StringPtr("<name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ResourceClientCreateEventHubConsumerGroupResult)
}
