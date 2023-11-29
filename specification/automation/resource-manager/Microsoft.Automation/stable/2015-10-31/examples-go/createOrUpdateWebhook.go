package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/432872fac1d0f8edcae98a0e8504afc0ee302710/specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/createOrUpdateWebhook.json
func ExampleWebhookClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebhookClient().CreateOrUpdate(ctx, "rg", "myAutomationAccount33", "TestWebhook", armautomation.WebhookCreateOrUpdateParameters{
		Name: to.Ptr("TestWebhook"),
		Properties: &armautomation.WebhookCreateOrUpdateProperties{
			ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-29T22:18:13.700Z"); return t }()),
			IsEnabled:  to.Ptr(true),
			Runbook: &armautomation.RunbookAssociationProperty{
				Name: to.Ptr("TestRunbook"),
			},
			URI: to.Ptr("<uri>"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Webhook = armautomation.Webhook{
	// 	Name: to.Ptr("TestWebhook"),
	// 	Type: to.Ptr("Microsoft.Automation/AutomationAccounts/Webhooks"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Automation/automationAccounts/myAutomationAccount33/webhooks/TestWebhook"),
	// 	Properties: &armautomation.WebhookProperties{
	// 		CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T22:18:14.665Z"); return t}()),
	// 		ExpiryTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-29T22:18:13.700Z"); return t}()),
	// 		IsEnabled: to.Ptr(true),
	// 		LastModifiedBy: to.Ptr(""),
	// 		LastModifiedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-29T22:18:14.665Z"); return t}()),
	// 		Runbook: &armautomation.RunbookAssociationProperty{
	// 			Name: to.Ptr("TestRunbook"),
	// 		},
	// 		URI: to.Ptr(""),
	// 	},
	// }
}
