package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListChannelsByBotService.json
func ExampleChannelsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChannelsClient().NewListByResourceGroupPager("OneResourceGroupName", "samplebotname", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ChannelResponseList = armbotservice.ChannelResponseList{
		// 	Value: []*armbotservice.BotChannel{
		// 		{
		// 			ID: to.Ptr("/subscriptions/subscription-id/providers/Microsoft.BotService/botServices"),
		// 			Location: to.Ptr("global"),
		// 			Properties: &armbotservice.EmailChannel{
		// 				ChannelName: to.Ptr("EmailChannel"),
		// 				Properties: &armbotservice.EmailChannelProperties{
		// 					EmailAddress: to.Ptr("a@b.com"),
		// 					IsEnabled: to.Ptr(true),
		// 					Password: to.Ptr("pwd"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/subscription-id/providers/Microsoft.BotService/botServices"),
		// 			Location: to.Ptr("global"),
		// 			Properties: &armbotservice.FacebookChannel{
		// 				ChannelName: to.Ptr("FacebookChannel"),
		// 				Properties: &armbotservice.FacebookChannelProperties{
		// 					AppID: to.Ptr("id"),
		// 					CallbackURL: to.Ptr("appid"),
		// 					IsEnabled: to.Ptr(true),
		// 					Pages: []*armbotservice.FacebookPage{
		// 						{
		// 							ID: to.Ptr("id"),
		// 					}},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
