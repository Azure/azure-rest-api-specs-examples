package armbotservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/botservice/armbotservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/ListChannel.json
func ExampleChannelsClient_ListWithKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbotservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewChannelsClient().ListWithKeys(ctx, "OneResourceGroupName", "samplebotname", armbotservice.ChannelNameEmailChannel, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListChannelWithKeysResponse = armbotservice.ListChannelWithKeysResponse{
	// 	Location: to.Ptr("global"),
	// 	Properties: &armbotservice.EmailChannel{
	// 		ChannelName: to.Ptr("EmailChannel"),
	// 		Properties: &armbotservice.EmailChannelProperties{
	// 			EmailAddress: to.Ptr("a@b.com"),
	// 			IsEnabled: to.Ptr(true),
	// 			Password: to.Ptr("pwd"),
	// 		},
	// 	},
	// 	Resource: &armbotservice.DirectLineChannel{
	// 		ChannelName: to.Ptr("DirectLineChannel"),
	// 		Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		Location: to.Ptr("aaaaaa"),
	// 		Properties: &armbotservice.DirectLineChannelProperties{
	// 			Sites: []*armbotservice.DirectLineSite{
	// 				{
	// 					IsBlockUserUploadEnabled: to.Ptr(false),
	// 					IsEnabled: to.Ptr(true),
	// 					IsSecureSiteEnabled: to.Ptr(false),
	// 					IsV1Enabled: to.Ptr(true),
	// 					IsV3Enabled: to.Ptr(true),
	// 					Key: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					Key2: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					SiteID: to.Ptr("aaaaaaaaaaa"),
	// 					SiteName: to.Ptr("aaaaaaaaaaaa"),
	// 					TrustedOrigins: []*string{
	// 						to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 				}},
	// 			},
	// 		},
	// 		Setting: &armbotservice.ChannelSettings{
	// 			BotIconURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			BotID: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 			ChannelDisplayName: to.Ptr("aaaaaaaaaaa"),
	// 			ChannelID: to.Ptr("aaaaaaaaaa"),
	// 			DisableLocalAuth: to.Ptr(false),
	// 			ExtensionKey1: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			ExtensionKey2: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			IsEnabled: to.Ptr(true),
	// 			Sites: []*armbotservice.Site{
	// 				{
	// 					ETag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					IsBlockUserUploadEnabled: to.Ptr(false),
	// 					IsEnabled: to.Ptr(true),
	// 					IsSecureSiteEnabled: to.Ptr(false),
	// 					IsTokenEnabled: to.Ptr(false),
	// 					IsV1Enabled: to.Ptr(true),
	// 					IsV3Enabled: to.Ptr(true),
	// 					IsWebchatPreviewEnabled: to.Ptr(false),
	// 					Key: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					Key2: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					SiteID: to.Ptr("aaaaaaaaaaa"),
	// 					SiteName: to.Ptr("aaaaaaaaaaaa"),
	// 					TrustedOrigins: []*string{
	// 						to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 				}},
	// 			},
	// 		}
}
