package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: 2024-10-23/updateWebhook.json
func ExampleWebhookClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebhookClient().Update(ctx, "rg", "myAutomationAccount33", "TestWebhook", armautomation.WebhookUpdateParameters{
		Name: to.Ptr("TestWebhook"),
		Properties: &armautomation.WebhookUpdateProperties{
			Description: to.Ptr("updated webhook"),
			IsEnabled:   to.Ptr(false),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armautomation.WebhookClientUpdateResponse{
	// 	Webhook: armautomation.Webhook{
	// 		Name: to.Ptr("TestWebhook"),
	// 		ID: to.Ptr("/subscriptions/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/webhooks/TestWebhook"),
	// 		Properties: &armautomation.WebhookProperties{
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T21:52:01.272378+00:00"); return t}()),
	// 			ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-30T23:03:14.5752078+00:00"); return t}()),
	// 			IsEnabled: to.Ptr(false),
	// 			LastInvokedTime: nil,
	// 			LastModifiedBy: to.Ptr(""),
	// 			LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-30T23:03:14.9069441+00:00"); return t}()),
	// 			Runbook: &armautomation.RunbookAssociationProperty{
	// 				Name: to.Ptr("TestRunbook"),
	// 			},
	// 			URI: to.Ptr(""),
	// 		},
	// 	},
	// }
}
