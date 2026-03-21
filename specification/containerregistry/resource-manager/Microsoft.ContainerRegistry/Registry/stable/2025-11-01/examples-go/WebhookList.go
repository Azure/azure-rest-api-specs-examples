package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2025-11-01/WebhookList.json
func ExampleWebhooksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebhooksClient().NewListPager("myResourceGroup", "myRegistry", nil)
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
		// page = armcontainerregistry.WebhooksClientListResponse{
		// 	WebhookListResult: armcontainerregistry.WebhookListResult{
		// 		Value: []*armcontainerregistry.Webhook{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/webhooks/myWebhook"),
		// 				Name: to.Ptr("myWebhook"),
		// 				Type: to.Ptr("Microsoft.ContainerRegistry/registries/webhooks"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"key": to.Ptr("value"),
		// 				},
		// 				Properties: &armcontainerregistry.WebhookProperties{
		// 					Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
		// 					Scope: to.Ptr("myRepository"),
		// 					Actions: []*armcontainerregistry.WebhookAction{
		// 						to.Ptr(armcontainerregistry.WebhookActionPush),
		// 					},
		// 					ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
