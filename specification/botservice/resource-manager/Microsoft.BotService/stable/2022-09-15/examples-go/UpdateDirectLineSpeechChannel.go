package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/UpdateDirectLineSpeechChannel.json
func ExampleChannelsClient_Update_updateDirectLineSpeechChannel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewChannelsClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "OneResourceGroupName", "samplebotname", armbotservice.ChannelNameDirectLineSpeechChannel, armbotservice.BotChannel{
		Location: to.Ptr("global"),
		Properties: &armbotservice.DirectLineSpeechChannel{
			ChannelName: to.Ptr("DirectLineSpeechChannel"),
			Properties: &armbotservice.DirectLineSpeechChannelProperties{
				CognitiveServiceRegion:          to.Ptr("XcognitiveServiceRegionX"),
				CognitiveServiceSubscriptionKey: to.Ptr("XcognitiveServiceSubscriptionKeyX"),
				IsEnabled:                       to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
