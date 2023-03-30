package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_List.json
func ExampleNotificationChannelsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevtestlabs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNotificationChannelsClient().NewListPager("resourceGroupName", "{labName}", &armdevtestlabs.NotificationChannelsClientListOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
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
		// page.NotificationChannelList = armdevtestlabs.NotificationChannelList{
		// 	Value: []*armdevtestlabs.NotificationChannel{
		// 		{
		// 			Name: to.Ptr("autoshutdown"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/notificationChannels"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/notificationChannels/{notificationChannelName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.NotificationChannelProperties{
		// 				Description: to.Ptr("Integration configured for auto-shutdown"),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-05T02:12:10.6566371+00:00"); return t}()),
		// 				EmailRecipient: to.Ptr("{email}"),
		// 				Events: []*armdevtestlabs.Event{
		// 					{
		// 						EventName: to.Ptr(armdevtestlabs.NotificationChannelEventTypeAutoShutdown),
		// 				}},
		// 				NotificationLocale: to.Ptr("en"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				WebHookURL: to.Ptr("{webhookUrl}"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("costThreshold"),
		// 			Type: to.Ptr("Microsoft.DevTestLab/labs/notificationChannels"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/notificationChannels/{notificationChannelName}"),
		// 			Location: to.Ptr("{location}"),
		// 			Tags: map[string]*string{
		// 				"tagName1": to.Ptr("tagValue1"),
		// 			},
		// 			Properties: &armdevtestlabs.NotificationChannelProperties{
		// 				Description: to.Ptr("Integration configured for cost"),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-05T02:14:05.2392034+00:00"); return t}()),
		// 				Events: []*armdevtestlabs.Event{
		// 					{
		// 						EventName: to.Ptr(armdevtestlabs.NotificationChannelEventTypeCost),
		// 				}},
		// 				NotificationLocale: to.Ptr("en"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				UniqueIdentifier: to.Ptr("{uniqueIdentifier}"),
		// 				WebHookURL: to.Ptr("{webhookUrl}"),
		// 			},
		// 	}},
		// }
	}
}
