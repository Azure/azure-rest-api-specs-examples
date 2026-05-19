package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice/v2"
)

// Generated from example definition: 2023-09-15-preview/PutEmailChannel.json
func ExampleChannelsClient_Create_createEmailChannel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("subscription-id", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChannelsClient().Create(ctx, "OneResourceGroupName", "samplebotname", armbotservice.ChannelNameEmailChannel, armbotservice.BotChannel{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbotservice.ChannelsClientCreateResponse{
	// 	BotChannel: &armbotservice.BotChannel{
	// 		ID: to.Ptr("/subscriptions/subscription-id/providers/Microsoft.BotService/botServices"),
	// 		Location: to.Ptr("global"),
	// 		Properties: &armbotservice.EmailChannel{
	// 			ChannelName: to.Ptr("EmailChannel"),
	// 			Properties: &armbotservice.EmailChannelProperties{
	// 				AuthMethod: to.Ptr(			armbotservice.EmailChannelAuthMethodGraph),
	// 				EmailAddress: to.Ptr("a@b.com"),
	// 				IsEnabled: to.Ptr(true),
	// 			},
	// 		},
	// 	},
	// }
}
