package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/GetAlexaChannel.json
func ExampleChannelsClient_Get_getAlexaChannel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChannelsClient().Get(ctx, "OneResourceGroupName", "samplebotname", "AlexaChannel", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BotChannel = armbotservice.BotChannel{
	// 	ID: to.Ptr("/subscriptions/subscription-id/providers/Microsoft.BotService/botServices"),
	// 	Location: to.Ptr("global"),
	// 	Properties: &armbotservice.AlexaChannel{
	// 		ChannelName: to.Ptr("AlexaChannel"),
	// 		Properties: &armbotservice.AlexaChannelProperties{
	// 			AlexaSkillID: to.Ptr("alexa skill id"),
	// 			IsEnabled: to.Ptr(true),
	// 			ServiceEndpointURI: to.Ptr("https://domain/XUrlFragmentX/botId"),
	// 			URLFragment: to.Ptr("XUrlFragmentX"),
	// 		},
	// 	},
	// }
}
