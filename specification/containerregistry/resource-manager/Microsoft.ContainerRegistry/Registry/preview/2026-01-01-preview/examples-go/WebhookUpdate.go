package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry/v3"
)

// Generated from example definition: 2026-01-01-preview/WebhookUpdate.json
func ExampleWebhooksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebhooksClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myWebhook", armcontainerregistry.WebhookUpdateParameters{
		Tags: map[string]*string{
			"key": to.Ptr("value"),
		},
		Properties: &armcontainerregistry.WebhookPropertiesUpdateParameters{
			ServiceURI: to.Ptr("http://myservice.com"),
			CustomHeaders: map[string]*string{
				"Authorization": to.Ptr("******"),
			},
			Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
			Scope:  to.Ptr("myRepository"),
			Actions: []*armcontainerregistry.WebhookAction{
				to.Ptr(armcontainerregistry.WebhookActionPush),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcontainerregistry.WebhooksClientUpdateResponse{
	// 	Webhook: &armcontainerregistry.Webhook{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/webhooks/myWebhook"),
	// 		Name: to.Ptr("myWebhook"),
	// 		Type: to.Ptr("Microsoft.ContainerRegistry/registries/webhooks"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"key": to.Ptr("value"),
	// 		},
	// 		Properties: &armcontainerregistry.WebhookProperties{
	// 			Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
	// 			Scope: to.Ptr("myRepository"),
	// 			Actions: []*armcontainerregistry.WebhookAction{
	// 				to.Ptr(armcontainerregistry.WebhookActionPush),
	// 			},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
