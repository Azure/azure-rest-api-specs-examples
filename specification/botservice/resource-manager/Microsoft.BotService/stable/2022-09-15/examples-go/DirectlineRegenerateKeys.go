package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/DirectlineRegenerateKeys.json
func ExampleDirectLineClient_RegenerateKeys_regenerateKeysForDirectLineChannelSite() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDirectLineClient().RegenerateKeys(ctx, "OneResourceGroupName", "samplebotname", armbotservice.RegenerateKeysChannelNameDirectLineChannel, armbotservice.SiteInfo{
		Key:      to.Ptr(armbotservice.KeyKey1),
		SiteName: to.Ptr("testSiteName"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BotChannel = armbotservice.BotChannel{
	// 	Location: to.Ptr("global"),
	// 	Properties: &armbotservice.DirectLineChannel{
	// 		ChannelName: to.Ptr("DirectLineChannel"),
	// 		Properties: &armbotservice.DirectLineChannelProperties{
	// 			Sites: []*armbotservice.DirectLineSite{
	// 				{
	// 					IsEnabled: to.Ptr(true),
	// 					IsV1Enabled: to.Ptr(true),
	// 					IsV3Enabled: to.Ptr(true),
	// 					Key: to.Ptr("key1"),
	// 					Key2: to.Ptr("key2"),
	// 					SiteID: to.Ptr("abcd"),
	// 					SiteName: to.Ptr("Default Site"),
	// 			}},
	// 		},
	// 	},
	// }
}
