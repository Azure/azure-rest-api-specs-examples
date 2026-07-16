package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2023-05-01-preview/getTestNotificationsAtTenantActionGroupResourceLevel.json
func ExampleCombineClient_GetTestNotificationsAtTenantActionGroupResourceLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCombineClient().GetTestNotificationsAtTenantActionGroupResourceLevel(ctx, "11111111-1111-1111-1111-111111111111", "testTenantActionGroup", "11000222191287", "72f988bf-86f1-41af-91ab-2d7cd011db47", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.CombineClientGetTestNotificationsAtTenantActionGroupResourceLevelResponse{
	// 	TenantActionGroupTestNotificationDetailsResponse: armmonitor.TenantActionGroupTestNotificationDetailsResponse{
	// 		ActionDetails: []*armmonitor.TenantActionGroupActionDetail{
	// 			{
	// 				Name: to.Ptr("AzureAppPush-name"),
	// 				MechanismType: to.Ptr("AzureAppPush"),
	// 				SendTime: to.Ptr("2021-09-21T04:52:42.8620629+00:00"),
	// 				Status: to.Ptr("Completed"),
	// 				SubState: to.Ptr("Default"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Email-name"),
	// 				MechanismType: to.Ptr("Email"),
	// 				SendTime: to.Ptr("2021-09-21T04:52:40.7480368+00:00"),
	// 				Status: to.Ptr("Completed"),
	// 				SubState: to.Ptr("Default"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Webhook-name"),
	// 				MechanismType: to.Ptr("Webhook"),
	// 				SendTime: to.Ptr("2021-09-21T04:52:42.0723479+00:00"),
	// 				Status: to.Ptr("Completed"),
	// 				SubState: to.Ptr("Default"),
	// 			},
	// 			{
	// 				Name: to.Ptr("SecureWebhook-name"),
	// 				MechanismType: to.Ptr("SecureWebhook"),
	// 				SendTime: to.Ptr("2021-09-21T04:52:42.0723479+00:00"),
	// 				Status: to.Ptr("Completed"),
	// 				SubState: to.Ptr("Default"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Sms-name"),
	// 				MechanismType: to.Ptr("Sms"),
	// 				SendTime: to.Ptr("2021-09-21T04:52:41.353015+00:00"),
	// 				Status: to.Ptr("Completed"),
	// 				SubState: to.Ptr("Default"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Voice-name"),
	// 				MechanismType: to.Ptr("Voice"),
	// 				SendTime: to.Ptr("2021-09-21T04:52:41.6330734+00:00"),
	// 				Status: to.Ptr("Completed"),
	// 				SubState: to.Ptr("Default"),
	// 			},
	// 		},
	// 		CompletedTime: to.Ptr("0001-01-01T00:00:00+00:00"),
	// 		Context: &armmonitor.Context{
	// 			ContextType: to.Ptr("Microsoft.Insights/ServiceHealth"),
	// 			NotificationSource: to.Ptr("Microsoft.Insights/TestNotification"),
	// 		},
	// 		CreatedTime: to.Ptr("2021-09-21T04:52:29.5091168+00:00"),
	// 		State: to.Ptr("Completed"),
	// 	},
	// }
}
