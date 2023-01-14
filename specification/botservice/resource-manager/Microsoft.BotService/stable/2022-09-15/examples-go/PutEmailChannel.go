package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/PutEmailChannel.json
func ExampleChannelsClient_Create_createEmailChannel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbotservice.NewChannelsClient("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx, "OneResourceGroupName", "samplebotname", armbotservice.ChannelNameEmailChannel, armbotservice.BotChannel{
		Location: to.Ptr("global"),
		Properties: &armbotservice.EmailChannel{
			ChannelName: to.Ptr("EmailChannel"),
			Properties: &armbotservice.EmailChannelProperties{
				AuthMethod:   to.Ptr(armbotservice.EmailChannelAuthMethodGraph),
				EmailAddress: to.Ptr("a@b.com"),
				IsEnabled:    to.Ptr(true),
				MagicCode:    to.Ptr("000000"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
