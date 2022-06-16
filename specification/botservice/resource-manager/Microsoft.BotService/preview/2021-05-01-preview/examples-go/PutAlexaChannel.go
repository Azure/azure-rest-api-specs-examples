package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutAlexaChannel.json
func ExampleChannelsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewChannelsClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"OneResourceGroupName",
		"samplebotname",
		armbotservice.ChannelNameAlexaChannel,
		armbotservice.BotChannel{
			Location: to.Ptr("global"),
			Properties: &armbotservice.AlexaChannel{
				ChannelName: to.Ptr("AlexaChannel"),
				Properties: &armbotservice.AlexaChannelProperties{
					AlexaSkillID: to.Ptr("XAlexaSkillIdX"),
					IsEnabled:    to.Ptr(true),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
