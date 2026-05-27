package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/listWebhooksByAutomationAccount.json
func ExampleWebhookClient_NewListByAutomationAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebhookClient().NewListByAutomationAccountPager("rg", "myAutomationAccount33", nil)
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
		// page = armautomation.WebhookClientListByAutomationAccountResponse{
		// 	WebhookListResult: armautomation.WebhookListResult{
		// 		Value: []*armautomation.Webhook{
		// 			{
		// 				Name: to.Ptr("TestWebhook"),
		// 				ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/webhooks/TestWebhook"),
		// 				Properties: &armautomation.WebhookProperties{
		// 					CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T22:18:14.6651862+00:00"); return t}()),
		// 					ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-29T22:18:13.7002872+00:00"); return t}()),
		// 					IsEnabled: to.Ptr(true),
		// 					LastInvokedTime: nil,
		// 					LastModifiedBy: to.Ptr(""),
		// 					LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T22:18:14.6651862+00:00"); return t}()),
		// 					Runbook: &armautomation.RunbookAssociationProperty{
		// 						Name: to.Ptr("TestRunbook"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
