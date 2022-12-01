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
	}
	ctx := context.Background()
	client, err := armeventgrid.NewEventChannelsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "examplerg", "examplePartnerNamespaceName1", "exampleEventChannelName1", armeventgrid.EventChannel{
		Properties: &armeventgrid.EventChannelProperties{
			Destination: &armeventgrid.EventChannelDestination{
				AzureSubscriptionID: to.Ptr("5b4b650e-28b9-4790-b3ab-ddbd88d727c4"),
				PartnerTopicName:    to.Ptr("examplePartnerTopic1"),
				ResourceGroup:       to.Ptr("examplerg2"),
			},
			Source: &armeventgrid.EventChannelSource{
				Source: to.Ptr("ContosoCorp.Accounts.User1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
