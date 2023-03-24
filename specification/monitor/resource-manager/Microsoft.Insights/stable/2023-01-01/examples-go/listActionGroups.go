package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/monitor/resource-manager/Microsoft.Insights/stable/2023-01-01/examples/listActionGroups.json
func ExampleActionGroupsClient_NewListBySubscriptionIDPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewActionGroupsClient().NewListBySubscriptionIDPager(nil)
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
		// page.ActionGroupList = armmonitor.ActionGroupList{
		// 	Value: []*armmonitor.ActionGroupResource{
		// 		{
		// 			Name: to.Ptr("SampleActionGroup"),
		// 			Type: to.Ptr("Microsoft.Insights/ActionGroups"),
		// 			ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/Default-NotificationRules/providers/microsoft.insights/actionGroups/SampleActionGroup"),
		// 			Location: to.Ptr("Global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armmonitor.ActionGroup{
		// 				ArmRoleReceivers: []*armmonitor.ArmRoleReceiver{
		// 				},
		// 				AutomationRunbookReceivers: []*armmonitor.AutomationRunbookReceiver{
		// 				},
		// 				AzureAppPushReceivers: []*armmonitor.AzureAppPushReceiver{
		// 				},
		// 				AzureFunctionReceivers: []*armmonitor.AzureFunctionReceiver{
		// 				},
		// 				EmailReceivers: []*armmonitor.EmailReceiver{
		// 					{
		// 						Name: to.Ptr("John Doe's email"),
		// 						EmailAddress: to.Ptr("johndoe@email.com"),
		// 						Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
		// 						UseCommonAlertSchema: to.Ptr(true),
		// 					},
		// 					{
		// 						Name: to.Ptr("Jane Smith's email"),
		// 						EmailAddress: to.Ptr("janesmith@email.com"),
		// 						Status: to.Ptr(armmonitor.ReceiverStatusDisabled),
		// 						UseCommonAlertSchema: to.Ptr(true),
		// 				}},
		// 				Enabled: to.Ptr(true),
		// 				EventHubReceivers: []*armmonitor.EventHubReceiver{
		// 				},
		// 				GroupShortName: to.Ptr("sample"),
		// 				ItsmReceivers: []*armmonitor.ItsmReceiver{
		// 				},
		// 				LogicAppReceivers: []*armmonitor.LogicAppReceiver{
		// 				},
		// 				SmsReceivers: []*armmonitor.SmsReceiver{
		// 					{
		// 						Name: to.Ptr("John Doe's mobile"),
		// 						CountryCode: to.Ptr("1"),
		// 						PhoneNumber: to.Ptr("1234567890"),
		// 						Status: to.Ptr(armmonitor.ReceiverStatusDisabled),
		// 					},
		// 					{
		// 						Name: to.Ptr("Jane Smith's mobile"),
		// 						CountryCode: to.Ptr("1"),
		// 						PhoneNumber: to.Ptr("0987654321"),
		// 						Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
		// 				}},
		// 				VoiceReceivers: []*armmonitor.VoiceReceiver{
		// 				},
		// 				WebhookReceivers: []*armmonitor.WebhookReceiver{
		// 					{
		// 						Name: to.Ptr("Sample webhook"),
		// 						ServiceURI: to.Ptr("http://www.example.com/webhook"),
		// 						UseCommonAlertSchema: to.Ptr(false),
		// 					},
		// 					{
		// 						Name: to.Ptr("Sample webhook 2"),
		// 						IdentifierURI: to.Ptr("http://someidentifier/d7811ba3-7996-4a93-99b6-6b2f3f355f8a"),
		// 						ObjectID: to.Ptr("d3bb868c-fe44-452c-aa26-769a6538c808"),
		// 						ServiceURI: to.Ptr("http://www.example.com/webhook2"),
		// 						TenantID: to.Ptr("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
		// 						UseAADAuth: to.Ptr(true),
		// 						UseCommonAlertSchema: to.Ptr(true),
		// 				}},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("SampleActionGroup2"),
		// 			Type: to.Ptr("Microsoft.Insights/ActionGroups"),
		// 			ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/Default-NotificationRules/providers/microsoft.insights/actionGroups/SampleActionGroup2"),
		// 			Location: to.Ptr("Global"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armmonitor.ActionGroup{
		// 				ArmRoleReceivers: []*armmonitor.ArmRoleReceiver{
		// 				},
		// 				AutomationRunbookReceivers: []*armmonitor.AutomationRunbookReceiver{
		// 				},
		// 				AzureAppPushReceivers: []*armmonitor.AzureAppPushReceiver{
		// 				},
		// 				AzureFunctionReceivers: []*armmonitor.AzureFunctionReceiver{
		// 				},
		// 				EmailReceivers: []*armmonitor.EmailReceiver{
		// 					{
		// 						Name: to.Ptr("John Doe's email"),
		// 						EmailAddress: to.Ptr("johndoe@email.com"),
		// 						Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
		// 						UseCommonAlertSchema: to.Ptr(true),
		// 				}},
		// 				Enabled: to.Ptr(false),
		// 				EventHubReceivers: []*armmonitor.EventHubReceiver{
		// 				},
		// 				GroupShortName: to.Ptr("sample2"),
		// 				ItsmReceivers: []*armmonitor.ItsmReceiver{
		// 				},
		// 				LogicAppReceivers: []*armmonitor.LogicAppReceiver{
		// 				},
		// 				SmsReceivers: []*armmonitor.SmsReceiver{
		// 					{
		// 						Name: to.Ptr("Jane Smith's mobile"),
		// 						CountryCode: to.Ptr("1"),
		// 						PhoneNumber: to.Ptr("0987654321"),
		// 						Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
		// 				}},
		// 				VoiceReceivers: []*armmonitor.VoiceReceiver{
		// 				},
		// 				WebhookReceivers: []*armmonitor.WebhookReceiver{
		// 				},
		// 			},
		// 	}},
		// }
	}
}
