package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-10-01-preview/createOrUpdateActionGroup.json
func ExampleActionGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("187f412d-1758-44d9-b052-169e2564721d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewActionGroupsClient().CreateOrUpdate(ctx, "Default-NotificationRules", "SampleActionGroup", armmonitor.ActionGroupResource{
		Identity: &armmonitor.ManagedServiceIdentity{
			Type: to.Ptr(armmonitor.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armmonitor.UserAssignedIdentity{
				"/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ThomasTestManagedIdentity_123": {},
				"/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ThomasTestManagedIdentity_456": {},
			},
		},
		Location: to.Ptr("Global"),
		Properties: &armmonitor.ActionGroup{
			ArmRoleReceivers: []*armmonitor.ArmRoleReceiver{
				{
					Name:                 to.Ptr("Sample armRole"),
					RoleID:               to.Ptr("8e3af657-a8ff-443c-a75c-2fe8c4bcb635"),
					UseCommonAlertSchema: to.Ptr(true),
				},
			},
			AutomationRunbookReceivers: []*armmonitor.AutomationRunbookReceiver{
				{
					Name:                 to.Ptr("testRunbook"),
					AutomationAccountID:  to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest"),
					IsGlobalRunbook:      to.Ptr(false),
					ManagedIdentity:      to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
					RunbookName:          to.Ptr("Sample runbook"),
					ServiceURI:           to.Ptr("<serviceUri>"),
					UseCommonAlertSchema: to.Ptr(true),
					WebhookResourceID:    to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084"),
				},
			},
			AzureAppPushReceivers: []*armmonitor.AzureAppPushReceiver{
				{
					Name:         to.Ptr("Sample azureAppPush"),
					EmailAddress: to.Ptr("johndoe@email.com"),
				},
			},
			AzureFunctionReceivers: []*armmonitor.AzureFunctionReceiver{
				{
					Name:                  to.Ptr("Sample azureFunction"),
					FunctionAppResourceID: to.Ptr("/subscriptions/5def922a-3ed4-49c1-b9fd-05ec533819a3/resourceGroups/aznsTest/providers/Microsoft.Web/sites/testFunctionApp"),
					FunctionName:          to.Ptr("HttpTriggerCSharp1"),
					HTTPTriggerURL:        to.Ptr("http://test.me"),
					ManagedIdentity:       to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
					UseCommonAlertSchema:  to.Ptr(true),
				},
			},
			EmailReceivers: []*armmonitor.EmailReceiver{
				{
					Name:                 to.Ptr("John Doe's email"),
					EmailAddress:         to.Ptr("johndoe@email.com"),
					UseCommonAlertSchema: to.Ptr(false),
				},
				{
					Name:                 to.Ptr("Jane Smith's email"),
					EmailAddress:         to.Ptr("janesmith@email.com"),
					UseCommonAlertSchema: to.Ptr(true),
				},
			},
			Enabled: to.Ptr(true),
			EventHubReceivers: []*armmonitor.EventHubReceiver{
				{
					Name:              to.Ptr("Sample eventHub"),
					EventHubName:      to.Ptr("testEventHub"),
					EventHubNameSpace: to.Ptr("testEventHubNameSpace"),
					ManagedIdentity:   to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
					SubscriptionID:    to.Ptr("187f412d-1758-44d9-b052-169e2564721d"),
					TenantID:          to.Ptr("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
				},
			},
			GroupShortName: to.Ptr("sample"),
			IncidentReceivers: []*armmonitor.IncidentReceiver{
				{
					Name: to.Ptr("IncidentAction"),
					Connection: &armmonitor.IncidentServiceConnection{
						Name: to.Ptr("IncidentConnection"),
						ID:   to.Ptr("8be638e7-1419-42d4-a059-437a5f4f4e4e"),
					},
					IncidentManagementService: to.Ptr(armmonitor.IncidentManagementServiceIcm),
					Mappings: map[string]*string{
						"icm.automitigationenabled":         to.Ptr("true"),
						"icm.correlationid":                 to.Ptr("${data.essentials.signalType}://${data.essentials.originAlertId}"),
						"icm.monitorid":                     to.Ptr("${data.essentials.alertRule}"),
						"icm.occurringlocation.environment": to.Ptr("PROD"),
						"icm.routingid":                     to.Ptr("${data.essentials.monitoringService}://${data.essentials.signalType}"),
						"icm.title":                         to.Ptr("${data.essentials.severity}:${data.essentials.monitorCondition} ${data.essentials.monitoringService}:${data.essentials.signalType} ${data.essentials.alertTargetIds}"),
						"icm.tsgid":                         to.Ptr("https://microsoft.com"),
					},
				},
			},
			ItsmReceivers: []*armmonitor.ItsmReceiver{
				{
					Name:                to.Ptr("Sample itsm"),
					ConnectionID:        to.Ptr("a3b9076c-ce8e-434e-85b4-aff10cb3c8f1"),
					Region:              to.Ptr("westcentralus"),
					TicketConfiguration: to.Ptr("{\"PayloadRevision\":0,\"WorkItemType\":\"Incident\",\"UseTemplate\":false,\"WorkItemData\":\"{}\",\"CreateOneWIPerCI\":false}"),
					WorkspaceID:         to.Ptr("5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c"),
				},
			},
			LogicAppReceivers: []*armmonitor.LogicAppReceiver{
				{
					Name:                 to.Ptr("Sample logicApp"),
					CallbackURL:          to.Ptr("https://prod-27.northcentralus.logic.azure.com/workflows/68e572e818e5457ba898763b7db90877/triggers/manual/paths/invoke/azns/test?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=Abpsb72UYJxPPvmDo937uzofupO5r_vIeWEx7KVHo7w"),
					ManagedIdentity:      to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
					ResourceID:           to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/LogicApp/providers/Microsoft.Logic/workflows/testLogicApp"),
					UseCommonAlertSchema: to.Ptr(false),
				},
			},
			SmsReceivers: []*armmonitor.SmsReceiver{
				{
					Name:        to.Ptr("John Doe's mobile"),
					CountryCode: to.Ptr("1"),
					PhoneNumber: to.Ptr("1234567890"),
				},
				{
					Name:        to.Ptr("Jane Smith's mobile"),
					CountryCode: to.Ptr("1"),
					PhoneNumber: to.Ptr("0987654321"),
				},
			},
			VoiceReceivers: []*armmonitor.VoiceReceiver{
				{
					Name:        to.Ptr("Sample voice"),
					CountryCode: to.Ptr("1"),
					PhoneNumber: to.Ptr("1234567890"),
				},
			},
			WebhookReceivers: []*armmonitor.WebhookReceiver{
				{
					Name:                 to.Ptr("Sample webhook 1"),
					ServiceURI:           to.Ptr("http://www.example.com/webhook1"),
					UseCommonAlertSchema: to.Ptr(true),
				},
				{
					Name:                 to.Ptr("Sample webhook 2"),
					IdentifierURI:        to.Ptr("http://someidentifier/d7811ba3-7996-4a93-99b6-6b2f3f355f8a"),
					ManagedIdentity:      to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
					ObjectID:             to.Ptr("d3bb868c-fe44-452c-aa26-769a6538c808"),
					ServiceURI:           to.Ptr("http://www.example.com/webhook2"),
					TenantID:             to.Ptr("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
					UseAADAuth:           to.Ptr(true),
					UseCommonAlertSchema: to.Ptr(true),
				},
			},
		},
		Tags: map[string]*string{},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.ActionGroupsClientCreateOrUpdateResponse{
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
	// 				{
	// 					Name: to.Ptr("Sample armRole"),
	// 					RoleID: to.Ptr("8e3af657-a8ff-443c-a75c-2fe8c4bcb635"),
	// 					UseCommonAlertSchema: to.Ptr(true),
	// 				},
	// 			},
	// 			AutomationRunbookReceivers: []*armmonitor.AutomationRunbookReceiver{
	// 				{
	// 					Name: to.Ptr("testRunbook"),
	// 					AutomationAccountID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest"),
	// 					IsGlobalRunbook: to.Ptr(false),
	// 					ManagedIdentity: to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
	// 					RunbookName: to.Ptr("Sample runbook"),
	// 					ServiceURI: to.Ptr("<serviceUri>"),
	// 					UseCommonAlertSchema: to.Ptr(true),
	// 					WebhookResourceID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/runbookTest/providers/Microsoft.Automation/automationAccounts/runbooktest/webhooks/Alert1510184037084"),
	// 				},
	// 			},
	// 			AzureAppPushReceivers: []*armmonitor.AzureAppPushReceiver{
	// 				{
	// 					Name: to.Ptr("Sample azureAppPush"),
	// 					EmailAddress: to.Ptr("johndoe@email.com"),
	// 				},
	// 			},
	// 			AzureFunctionReceivers: []*armmonitor.AzureFunctionReceiver{
	// 				{
	// 					Name: to.Ptr("Sample azureFunction"),
	// 					FunctionAppResourceID: to.Ptr("/subscriptions/5def922a-3ed4-49c1-b9fd-05ec533819a3/resourceGroups/aznsTest/providers/Microsoft.Web/sites/testFunctionApp"),
	// 					FunctionName: to.Ptr("HttpTriggerCSharp1"),
	// 					HTTPTriggerURL: to.Ptr("<httpTriggerUrl>"),
	// 					ManagedIdentity: to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
	// 					UseCommonAlertSchema: to.Ptr(true),
	// 				},
	// 			},
	// 			EmailReceivers: []*armmonitor.EmailReceiver{
	// 				{
	// 					Name: to.Ptr("John Doe's email"),
	// 					EmailAddress: to.Ptr("johndoe@email.com"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
	// 					UseCommonAlertSchema: to.Ptr(false),
	// 				},
	// 				{
	// 					Name: to.Ptr("Jane Smith's email"),
	// 					EmailAddress: to.Ptr("janesmith@email.com"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
	// 					UseCommonAlertSchema: to.Ptr(true),
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			EventHubReceivers: []*armmonitor.EventHubReceiver{
	// 				{
	// 					Name: to.Ptr("Sample eventHub"),
	// 					EventHubName: to.Ptr("testEventHub"),
	// 					EventHubNameSpace: to.Ptr("testEventHubNameSpace"),
	// 					ManagedIdentity: to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
	// 					SubscriptionID: to.Ptr("187f412d-1758-44d9-b052-169e2564721d"),
	// 					TenantID: to.Ptr("68a4459a-ccb8-493c-b9da-dd30457d1b84"),
	// 					UseCommonAlertSchema: to.Ptr(false),
	// 				},
	// 			},
	// 			GroupShortName: to.Ptr("sample"),
	// 			IncidentReceivers: []*armmonitor.IncidentReceiver{
	// 				{
	// 					Name: to.Ptr("IncidentAction"),
	// 					Connection: &armmonitor.IncidentServiceConnection{
	// 						Name: to.Ptr("IncidentConnection"),
	// 						ID: to.Ptr("8be638e7-1419-42d4-a059-437a5f4f4e4e"),
	// 					},
	// 					IncidentManagementService: to.Ptr(armmonitor.IncidentManagementServiceIcm),
	// 					Mappings: map[string]*string{
	// 						"icm.automitigationenabled": to.Ptr("true"),
	// 						"icm.correlationid": to.Ptr("${data.essentials.signalType}://${data.essentials.originAlertId}"),
	// 						"icm.monitorid": to.Ptr("${data.essentials.alertRule}"),
	// 						"icm.occurringlocation.environment": to.Ptr("PROD"),
	// 						"icm.routingid": to.Ptr("${data.essentials.monitoringService}://${data.essentials.signalType}"),
	// 						"icm.title": to.Ptr("${data.essentials.severity}:${data.essentials.monitorCondition} ${data.essentials.monitoringService}:${data.essentials.signalType} ${data.essentials.alertTargetIds}"),
	// 						"icm.tsgid": to.Ptr("https://microsoft.com"),
	// 					},
	// 				},
	// 			},
	// 			ItsmReceivers: []*armmonitor.ItsmReceiver{
	// 				{
	// 					Name: to.Ptr("Sample itsm"),
	// 					ConnectionID: to.Ptr("a3b9076c-ce8e-434e-85b4-aff10cb3c8f1"),
	// 					Region: to.Ptr("westcentralus"),
	// 					TicketConfiguration: to.Ptr("{\"PayloadRevision\":0,\"WorkItemType\":\"Incident\",\"UseTemplate\":false,\"WorkItemData\":\"{}\",\"CreateOneWIPerCI\":false}"),
	// 					WorkspaceID: to.Ptr("5def922a-3ed4-49c1-b9fd-05ec533819a3|55dfd1f8-7e59-4f89-bf56-4c82f5ace23c"),
	// 				},
	// 			},
	// 			LogicAppReceivers: []*armmonitor.LogicAppReceiver{
	// 				{
	// 					Name: to.Ptr("Sample logicApp"),
	// 					CallbackURL: to.Ptr("https://prod-27.northcentralus.logic.azure.com/workflows/68e572e818e5457ba898763b7db90877/triggers/manual/paths/invoke/azns/test?api-version=2016-10-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=Abpsb72UYJxPPvmDo937uzofupO5r_vIeWEx7KVHo7w"),
	// 					ManagedIdentity: to.Ptr("30fe7a91-cd31-4edf-96ab-52883b3199cd"),
	// 					ResourceID: to.Ptr("/subscriptions/187f412d-1758-44d9-b052-169e2564721d/resourceGroups/LogicApp/providers/Microsoft.Logic/workflows/testLogicApp"),
	// 					UseCommonAlertSchema: to.Ptr(false),
	// 				},
	// 			},
	// 			SmsReceivers: []*armmonitor.SmsReceiver{
	// 				{
	// 					Name: to.Ptr("John Doe's mobile"),
	// 					CountryCode: to.Ptr("1"),
	// 					PhoneNumber: to.Ptr("1234567890"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
	// 				},
	// 				{
	// 					Name: to.Ptr("Jane Smith's mobile"),
	// 					CountryCode: to.Ptr("1"),
	// 					PhoneNumber: to.Ptr("0987654321"),
	// 					Status: to.Ptr(armmonitor.ReceiverStatusEnabled),
	// 				},
	// 			},
	// 			VoiceReceivers: []*armmonitor.VoiceReceiver{
	// 				{
	// 					Name: to.Ptr("Sample voice"),
	// 					CountryCode: to.Ptr("1"),
	// 					PhoneNumber: to.Ptr("1234567890"),
	// 				},
	// 			},
	// 			WebhookReceivers: []*armmonitor.WebhookReceiver{
	// 				{
	// 					Name: to.Ptr("Sample webhook 1"),
	// 					ServiceURI: to.Ptr("http://www.example.com/webhook1"),
	// 					UseCommonAlertSchema: to.Ptr(true),
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
