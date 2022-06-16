package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2021-12-01-preview/examples/WebhookUpdate.json
func ExampleWebhooksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerregistry.NewWebhooksClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<webhook-name>",
		armcontainerregistry.WebhookUpdateParameters{
			Properties: &armcontainerregistry.WebhookPropertiesUpdateParameters{
				Actions: []*armcontainerregistry.WebhookAction{
					to.Ptr(armcontainerregistry.WebhookActionPush)},
				CustomHeaders: map[string]*string{
					"Authorization": to.Ptr("Basic 000000000000000000000000000000000000000000000000000"),
				},
				Scope:      to.Ptr("<scope>"),
				ServiceURI: to.Ptr("<service-uri>"),
				Status:     to.Ptr(armcontainerregistry.WebhookStatusEnabled),
			},
			Tags: map[string]*string{
				"key": to.Ptr("value"),
			},
		},
		&armcontainerregistry.WebhooksClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
