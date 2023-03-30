package armengagementfabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/engagementfabric/armengagementfabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/engagementfabric/resource-manager/Microsoft.EngagementFabric/preview/2018-09-01/examples/ChannelsListExample.json
func ExampleChannelsClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armengagementfabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChannelsClient().NewListByAccountPager("ExampleRg", "ExampleAccount", nil)
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
		// page.ChannelList = armengagementfabric.ChannelList{
		// 	Value: []*armengagementfabric.Channel{
		// 		{
		// 			Name: to.Ptr("ExampleChannel"),
		// 			Type: to.Ptr("Microsoft.EngagementFabric/Accounts/Channels"),
		// 			ID: to.Ptr("subscriptions/EDBF0095-A524-4A84-95FB-F72DA41AA6A1/resourceGroups/ExampleRg/providers/Microsoft.EngagementFabric/Accounts/ExampleAccount/Channels/ExampleChannel"),
		// 			Properties: &armengagementfabric.ChannelProperties{
		// 				ChannelFunctions: []*string{
		// 					to.Ptr("MockFunction1"),
		// 					to.Ptr("MockFunction2")},
		// 					ChannelType: to.Ptr("MockChannel"),
		// 					Credentials: map[string]*string{
		// 						"AppId": to.Ptr("exampleApp"),
		// 						"AppKey": to.Ptr(""),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("ExampleChannel2"),
		// 				Type: to.Ptr("Microsoft.EngagementFabric/Accounts/Channels"),
		// 				ID: to.Ptr("subscriptions/EDBF0095-A524-4A84-95FB-F72DA41AA6A1/resourceGroups/ExampleRg/providers/Microsoft.EngagementFabric/Accounts/ExampleAccount/Channels/ExampleChannel2"),
		// 				Properties: &armengagementfabric.ChannelProperties{
		// 					ChannelFunctions: []*string{
		// 						to.Ptr("MockFunction1"),
		// 						to.Ptr("MockFunction3")},
		// 						ChannelType: to.Ptr("MockChannel2"),
		// 						Credentials: map[string]*string{
		// 							"AppId": to.Ptr("exampleApp2"),
		// 							"AppKey": to.Ptr(""),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
