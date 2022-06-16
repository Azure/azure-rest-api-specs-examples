package armdevtestlabs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_Notify.json
func ExampleNotificationChannelsClient_Notify() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevtestlabs.NewNotificationChannelsClient("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Notify(ctx,
		"resourceGroupName",
		"{labName}",
		"{notificationChannelName}",
		armdevtestlabs.NotifyParameters{
			EventName:   to.Ptr(armdevtestlabs.NotificationChannelEventTypeAutoShutdown),
			JSONPayload: to.Ptr("{\"eventType\":\"AutoShutdown\",\"subscriptionId\":\"{subscriptionId}\",\"resourceGroupName\":\"resourceGroupName\",\"labName\":\"{labName}\"}"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
