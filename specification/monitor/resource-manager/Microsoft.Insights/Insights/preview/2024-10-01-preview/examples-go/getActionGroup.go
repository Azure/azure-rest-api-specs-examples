package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-10-01-preview/getActionGroup.json
func ExampleActionGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewActionGroupsClient().Get(ctx, "Default-NotificationRules", "SampleActionGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.ActionGroupsClientGetResponse{
	// 	ActionGroupResource: armmonitor.ActionGroupResource{
	// 		Name: to.Ptr("SampleActionGroup"),
	// 		Type: to.Ptr("Microsoft.Insights/ActionGroups"),
	// 		ID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/Default-NotificationRules/providers/microsoft.insights/actionGroups/SampleActionGroup"),
	// 		Identity: &armmonitor.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmonitor.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armmonitor.UserAssignedIdentity{
	// 				"/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ThomasTestManagedIdentity_123": &armmonitor.UserAssignedIdentity{
	// 					ClientID: to.Ptr("73525f7b-16f5-439f-b150-e6f737b8cb16"),
	// 					PrincipalID: to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
	// 				},
	// 				"/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ThomasTestManagedIdentity_456": &armmonitor.UserAssignedIdentity{
	// 					ClientID: to.Ptr("4dfa770d-290d-465b-ad06-4a787b65c641"),
	// 					PrincipalID: to.Ptr("92b72e3e-dc08-4697-8643-32426387ebce"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("Global"),
	// 		Properties: &armmonitor.ActionGroup{
	// 			ArmRoleReceivers: []*armmonitor.ArmRoleReceiver{
	// 			},
	// 			AutomationRunbookReceivers: []*armmonitor.AutomationRunbookReceiver{
	// 			},
	// 			AzureAppPushReceivers: []*armmonitor.AzureAppPushReceiver{
	// 			},
	// 			AzureFunctionReceivers: []*armmonitor.AzureFunctionReceiver{
	// 			},
	// 			EmailReceivers: []*armmonitor.EmailReceiver{
	// 				{
	// 					Name: to.Ptr("John Doe's email"),
	// 					EmailAddress: to.Ptr("johndoe@email.com"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
	// 					UseCommonAlertSchema: to.Ptr(true),
	// 				},
	// 				{
	// 					Name: to.Ptr("Jane Smith's email"),
	// 					EmailAddress: to.Ptr("janesmith@email.com"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusDisabled),
	// 					UseCommonAlertSchema: to.Ptr(true),
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			EventHubReceivers: []*armmonitor.EventHubReceiver{
	// 			},
	// 			GroupShortName: to.Ptr("sample"),
	// 			IncidentReceivers: []*armmonitor.IncidentReceiver{
	// 			},
	// 			ItsmReceivers: []*armmonitor.ItsmReceiver{
	// 			},
	// 			LogicAppReceivers: []*armmonitor.LogicAppReceiver{
	// 			},
	// 			SmsReceivers: []*armmonitor.SmsReceiver{
	// 				{
	// 					Name: to.Ptr("John Doe's mobile"),
	// 					CountryCode: to.Ptr("1"),
	// 					PhoneNumber: to.Ptr("1234567890"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusDisabled),
	// 				},
	// 				{
	// 					Name: to.Ptr("Jane Smith's mobile"),
	// 					CountryCode: to.Ptr("1"),
	// 					PhoneNumber: to.Ptr("0987654321"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
	// 				},
	// 			},
	// 			VoiceReceivers: []*armmonitor.VoiceReceiver{
	// 			},
	// 			WebhookReceivers: []*armmonitor.WebhookReceiver{
	// 				{
	// 					Name: to.Ptr("Sample webhook"),
	// 					ServiceURI: to.Ptr("http://www.example.com/webhook"),
	// 					UseCommonAlertSchema: to.Ptr(false),
	// 				},
	// 				{
	// 					Name: to.Ptr("Sample webhook 2"),
	// 					IdentifierURI: to.Ptr("http://someidentifier/d7811ba3-7996-4a93-99b6-6b2f3f355f8a"),
	// 					ManagedIdentity: to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
	// 					ObjectID: to.Ptr("d3bb868c-fe44-452c-aa26-769a6538c808"),
	// 					ServiceURI: to.Ptr("http://www.example.com/webhook2"),
	// 					TenantID: to.Ptr("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
	// 					UseAADAuth: to.Ptr(true),
	// 					UseCommonAlertSchema: to.Ptr(true),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
