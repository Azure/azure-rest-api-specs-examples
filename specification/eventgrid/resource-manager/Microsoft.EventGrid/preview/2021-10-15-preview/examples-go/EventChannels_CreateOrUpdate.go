package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventChannels_CreateOrUpdate.json
func ExampleEventChannelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventgrid.NewEventChannelsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<partner-namespace-name>",
		"<event-channel-name>",
		armeventgrid.EventChannel{
			Properties: &armeventgrid.EventChannelProperties{
				Destination: &armeventgrid.EventChannelDestination{
					AzureSubscriptionID: to.Ptr("<azure-subscription-id>"),
					PartnerTopicName:    to.Ptr("<partner-topic-name>"),
					ResourceGroup:       to.Ptr("<resource-group>"),
				},
				Source: &armeventgrid.EventChannelSource{
					Source: to.Ptr("<source>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
