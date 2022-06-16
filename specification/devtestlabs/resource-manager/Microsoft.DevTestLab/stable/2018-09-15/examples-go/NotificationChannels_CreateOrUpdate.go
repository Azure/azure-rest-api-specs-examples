package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_CreateOrUpdate.json
func ExampleNotificationChannelsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewNotificationChannelsClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"resourceGroupName",
		"{labName}",
		"{notificationChannelName}",
		armdevtestlabs.NotificationChannel{
			Properties: &armdevtestlabs.NotificationChannelProperties{
				Description:    to.Ptr("Integration configured for auto-shutdown"),
				EmailRecipient: to.Ptr("{email}"),
				Events: []*armdevtestlabs.Event{
					{
						EventName: to.Ptr(armdevtestlabs.NotificationChannelEventTypeAutoShutdown),
					}},
				NotificationLocale: to.Ptr("en"),
				WebHookURL:         to.Ptr("{webhookUrl}"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
