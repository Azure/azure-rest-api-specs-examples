Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmonitor%2Farmmonitor%2Fv0.4.0/sdk/resourcemanager/monitor/armmonitor/README.md) on how to add the SDK to your project and authenticate.

```go
package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2021-09-01/examples/createOrUpdateActionGroup.json
func ExampleActionGroupsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmonitor.NewActionGroupsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<action-group-name>",
		armmonitor.ActionGroupResource{
			Location: to.StringPtr("<location>"),
			Tags:     map[string]*string{},
			Properties: &armmonitor.ActionGroup{
				ArmRoleReceivers: []*armmonitor.ArmRoleReceiver{
					{
						Name:                 to.StringPtr("<name>"),
						RoleID:               to.StringPtr("<role-id>"),
						UseCommonAlertSchema: to.BoolPtr(true),
					}},
				AutomationRunbookReceivers: []*armmonitor.AutomationRunbookReceiver{
					{
						Name:                 to.StringPtr("<name>"),
						AutomationAccountID:  to.StringPtr("<automation-account-id>"),
						IsGlobalRunbook:      to.BoolPtr(false),
						RunbookName:          to.StringPtr("<runbook-name>"),
						ServiceURI:           to.StringPtr("<service-uri>"),
						UseCommonAlertSchema: to.BoolPtr(true),
						WebhookResourceID:    to.StringPtr("<webhook-resource-id>"),
					}},
				AzureAppPushReceivers: []*armmonitor.AzureAppPushReceiver{
					{
						Name:         to.StringPtr("<name>"),
						EmailAddress: to.StringPtr("<email-address>"),
					}},
				AzureFunctionReceivers: []*armmonitor.AzureFunctionReceiver{
					{
						Name:                  to.StringPtr("<name>"),
						FunctionAppResourceID: to.StringPtr("<function-app-resource-id>"),
						FunctionName:          to.StringPtr("<function-name>"),
						HTTPTriggerURL:        to.StringPtr("<httptrigger-url>"),
						UseCommonAlertSchema:  to.BoolPtr(true),
					}},
				EmailReceivers: []*armmonitor.EmailReceiver{
					{
						Name:                 to.StringPtr("<name>"),
						EmailAddress:         to.StringPtr("<email-address>"),
						UseCommonAlertSchema: to.BoolPtr(false),
					},
					{
						Name:                 to.StringPtr("<name>"),
						EmailAddress:         to.StringPtr("<email-address>"),
						UseCommonAlertSchema: to.BoolPtr(true),
					}},
				Enabled: to.BoolPtr(true),
				EventHubReceivers: []*armmonitor.EventHubReceiver{
					{
						Name:              to.StringPtr("<name>"),
						EventHubName:      to.StringPtr("<event-hub-name>"),
						EventHubNameSpace: to.StringPtr("<event-hub-name-space>"),
						SubscriptionID:    to.StringPtr("<subscription-id>"),
						TenantID:          to.StringPtr("<tenant-id>"),
					}},
				GroupShortName: to.StringPtr("<group-short-name>"),
				ItsmReceivers: []*armmonitor.ItsmReceiver{
					{
						Name:                to.StringPtr("<name>"),
						ConnectionID:        to.StringPtr("<connection-id>"),
						Region:              to.StringPtr("<region>"),
						TicketConfiguration: to.StringPtr("<ticket-configuration>"),
						WorkspaceID:         to.StringPtr("<workspace-id>"),
					}},
				LogicAppReceivers: []*armmonitor.LogicAppReceiver{
					{
						Name:                 to.StringPtr("<name>"),
						CallbackURL:          to.StringPtr("<callback-url>"),
						ResourceID:           to.StringPtr("<resource-id>"),
						UseCommonAlertSchema: to.BoolPtr(false),
					}},
				SmsReceivers: []*armmonitor.SmsReceiver{
					{
						Name:        to.StringPtr("<name>"),
						CountryCode: to.StringPtr("<country-code>"),
						PhoneNumber: to.StringPtr("<phone-number>"),
					},
					{
						Name:        to.StringPtr("<name>"),
						CountryCode: to.StringPtr("<country-code>"),
						PhoneNumber: to.StringPtr("<phone-number>"),
					}},
				VoiceReceivers: []*armmonitor.VoiceReceiver{
					{
						Name:        to.StringPtr("<name>"),
						CountryCode: to.StringPtr("<country-code>"),
						PhoneNumber: to.StringPtr("<phone-number>"),
					}},
				WebhookReceivers: []*armmonitor.WebhookReceiver{
					{
						Name:                 to.StringPtr("<name>"),
						ServiceURI:           to.StringPtr("<service-uri>"),
						UseCommonAlertSchema: to.BoolPtr(true),
					},
					{
						Name:                 to.StringPtr("<name>"),
						IdentifierURI:        to.StringPtr("<identifier-uri>"),
						ObjectID:             to.StringPtr("<object-id>"),
						ServiceURI:           to.StringPtr("<service-uri>"),
						TenantID:             to.StringPtr("<tenant-id>"),
						UseAADAuth:           to.BoolPtr(true),
						UseCommonAlertSchema: to.BoolPtr(true),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ActionGroupsClientCreateOrUpdateResult)
}
```
