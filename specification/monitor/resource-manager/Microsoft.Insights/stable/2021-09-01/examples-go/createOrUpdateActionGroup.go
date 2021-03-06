package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/monitor/resource-manager/Microsoft.Insights/stable/2021-09-01/examples/createOrUpdateActionGroup.json
func ExampleActionGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmonitor.NewActionGroupsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<action-group-name>",
		armmonitor.ActionGroupResource{
			Location: to.Ptr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.ActionGroup{
				ArmRoleReceivers: []*armmonitor.ArmRoleReceiver{
					{
						Name:                 to.Ptr("<name>"),
						RoleID:               to.Ptr("<role-id>"),
						UseCommonAlertSchema: to.Ptr(true),
					}},
				AutomationRunbookReceivers: []*armmonitor.AutomationRunbookReceiver{
					{
						Name:                 to.Ptr("<name>"),
						AutomationAccountID:  to.Ptr("<automation-account-id>"),
						IsGlobalRunbook:      to.Ptr(false),
						RunbookName:          to.Ptr("<runbook-name>"),
						ServiceURI:           to.Ptr("<service-uri>"),
						UseCommonAlertSchema: to.Ptr(true),
						WebhookResourceID:    to.Ptr("<webhook-resource-id>"),
					}},
				AzureAppPushReceivers: []*armmonitor.AzureAppPushReceiver{
					{
						Name:         to.Ptr("<name>"),
						EmailAddress: to.Ptr("<email-address>"),
					}},
				AzureFunctionReceivers: []*armmonitor.AzureFunctionReceiver{
					{
						Name:                  to.Ptr("<name>"),
						FunctionAppResourceID: to.Ptr("<function-app-resource-id>"),
						FunctionName:          to.Ptr("<function-name>"),
						HTTPTriggerURL:        to.Ptr("<httptrigger-url>"),
						UseCommonAlertSchema:  to.Ptr(true),
					}},
				EmailReceivers: []*armmonitor.EmailReceiver{
					{
						Name:                 to.Ptr("<name>"),
						EmailAddress:         to.Ptr("<email-address>"),
						UseCommonAlertSchema: to.Ptr(false),
					},
					{
						Name:                 to.Ptr("<name>"),
						EmailAddress:         to.Ptr("<email-address>"),
						UseCommonAlertSchema: to.Ptr(true),
					}},
				Enabled: to.Ptr(true),
				EventHubReceivers: []*armmonitor.EventHubReceiver{
					{
						Name:              to.Ptr("<name>"),
						EventHubName:      to.Ptr("<event-hub-name>"),
						EventHubNameSpace: to.Ptr("<event-hub-name-space>"),
						SubscriptionID:    to.Ptr("<subscription-id>"),
						TenantID:          to.Ptr("<tenant-id>"),
					}},
				GroupShortName: to.Ptr("<group-short-name>"),
				ItsmReceivers: []*armmonitor.ItsmReceiver{
					{
						Name:                to.Ptr("<name>"),
						ConnectionID:        to.Ptr("<connection-id>"),
						Region:              to.Ptr("<region>"),
						TicketConfiguration: to.Ptr("<ticket-configuration>"),
						WorkspaceID:         to.Ptr("<workspace-id>"),
					}},
				LogicAppReceivers: []*armmonitor.LogicAppReceiver{
					{
						Name:                 to.Ptr("<name>"),
						CallbackURL:          to.Ptr("<callback-url>"),
						ResourceID:           to.Ptr("<resource-id>"),
						UseCommonAlertSchema: to.Ptr(false),
					}},
				SmsReceivers: []*armmonitor.SmsReceiver{
					{
						Name:        to.Ptr("<name>"),
						CountryCode: to.Ptr("<country-code>"),
						PhoneNumber: to.Ptr("<phone-number>"),
					},
					{
						Name:        to.Ptr("<name>"),
						CountryCode: to.Ptr("<country-code>"),
						PhoneNumber: to.Ptr("<phone-number>"),
					}},
				VoiceReceivers: []*armmonitor.VoiceReceiver{
					{
						Name:        to.Ptr("<name>"),
						CountryCode: to.Ptr("<country-code>"),
						PhoneNumber: to.Ptr("<phone-number>"),
					}},
				WebhookReceivers: []*armmonitor.WebhookReceiver{
					{
						Name:                 to.Ptr("<name>"),
						ServiceURI:           to.Ptr("<service-uri>"),
						UseCommonAlertSchema: to.Ptr(true),
					},
					{
						Name:                 to.Ptr("<name>"),
						IdentifierURI:        to.Ptr("<identifier-uri>"),
						ObjectID:             to.Ptr("<object-id>"),
						ServiceURI:           to.Ptr("<service-uri>"),
						TenantID:             to.Ptr("<tenant-id>"),
						UseAADAuth:           to.Ptr(true),
						UseCommonAlertSchema: to.Ptr(true),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
