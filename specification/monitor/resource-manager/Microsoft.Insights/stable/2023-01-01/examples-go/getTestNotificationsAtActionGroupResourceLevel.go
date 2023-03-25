package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/examples/getTestNotificationsAtActionGroupResourceLevel.json
func ExampleActionGroupsClient_GetTestNotificationsAtActionGroupResourceLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewActionGroupsClient().GetTestNotificationsAtActionGroupResourceLevel(ctx, "TestRgName", "TestAgName", "11000222191287", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.TestNotificationDetailsResponse = armmonitor.TestNotificationDetailsResponse{
	// 	ActionDetails: []*armmonitor.ActionDetail{
	// 		{
	// 			MechanismType: to.Ptr("AzureAppPush"),
	// 			Name: to.Ptr("AzureAppPush-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.8620629+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("AzureFunction"),
	// 			Name: to.Ptr("AzureFunction-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.0623319+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("Email"),
	// 			Name: to.Ptr("Email-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:40.7480368+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("LogicApp"),
	// 			Name: to.Ptr("LogicApp-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.2473419+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("Webhook"),
	// 			Name: to.Ptr("Webhook-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.0723479+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("SecureWebhook"),
	// 			Name: to.Ptr("SecureWebhook-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.0723479+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("Sms"),
	// 			Name: to.Ptr("Sms-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:41.353015+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("Voice"),
	// 			Name: to.Ptr("Voice-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:41.6330734+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("EventHub"),
	// 			Name: to.Ptr("EventHub-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.0723479+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("AutomationRunbook"),
	// 			Name: to.Ptr("AutomationRunbook-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.0723479+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 		},
	// 		{
	// 			MechanismType: to.Ptr("Itsm"),
	// 			Name: to.Ptr("Itsm-name"),
	// 			SendTime: to.Ptr("2021-09-21T04:52:42.0723479+00:00"),
	// 			Status: to.Ptr("Completed"),
	// 			SubState: to.Ptr("Default"),
	// 	}},
	// 	CompletedTime: to.Ptr("0001-01-01T00:00:00+00:00"),
	// 	Context: &armmonitor.Context{
	// 		ContextType: to.Ptr("Microsoft.Insights/Budget"),
	// 		NotificationSource: to.Ptr("Microsoft.Insights/TestNotification"),
	// 	},
	// 	CreatedTime: to.Ptr("2021-09-21T04:52:29.5091168+00:00"),
	// 	State: to.Ptr("Completed"),
	// }
}
