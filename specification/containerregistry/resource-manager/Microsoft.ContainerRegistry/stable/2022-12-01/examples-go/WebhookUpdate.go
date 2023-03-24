package armcontainerregistry_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/969fd0c2634fbcc1975d7abe3749330a5145a97c/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2022-12-01/examples/WebhookUpdate.json
func ExampleWebhooksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerregistry.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebhooksClient().BeginUpdate(ctx, "myResourceGroup", "myRegistry", "myWebhook", armcontainerregistry.WebhookUpdateParameters{
		Properties: &armcontainerregistry.WebhookPropertiesUpdateParameters{
			Actions: []*armcontainerregistry.WebhookAction{
				to.Ptr(armcontainerregistry.WebhookActionPush)},
			CustomHeaders: map[string]*string{
				"Authorization": to.Ptr("Basic 000000000000000000000000000000000000000000000000000"),
			},
			Scope:      to.Ptr("myRepository"),
			ServiceURI: to.Ptr("http://myservice.com"),
			Status:     to.Ptr(armcontainerregistry.WebhookStatusEnabled),
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// res.Webhook = armcontainerregistry.Webhook{
	// 	Name: to.Ptr("myWebhook"),
	// 	Type: to.Ptr("Microsoft.ContainerRegistry/registries/webhooks"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/webhooks/myWebhook"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armcontainerregistry.WebhookProperties{
	// 		Actions: []*armcontainerregistry.WebhookAction{
	// 			to.Ptr(armcontainerregistry.WebhookActionPush)},
	// 			ProvisioningState: to.Ptr(armcontainerregistry.ProvisioningStateSucceeded),
	// 			Scope: to.Ptr("myRepository"),
	// 			Status: to.Ptr(armcontainerregistry.WebhookStatusEnabled),
	// 		},
	// 	}
}
