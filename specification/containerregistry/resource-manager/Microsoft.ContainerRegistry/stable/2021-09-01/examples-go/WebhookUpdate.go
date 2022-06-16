package armcontainerregistry_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
)

// x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2021-09-01/examples/WebhookUpdate.json
func ExampleWebhooksClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerregistry.NewWebhooksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<registry-name>",
		"<webhook-name>",
		armcontainerregistry.WebhookUpdateParameters{
			Properties: &armcontainerregistry.WebhookPropertiesUpdateParameters{
				Actions: []*armcontainerregistry.WebhookAction{
					armcontainerregistry.WebhookAction("push").ToPtr()},
				CustomHeaders: map[string]*string{
					"Authorization": to.StringPtr("Basic 000000000000000000000000000000000000000000000000000"),
				},
				Scope:      to.StringPtr("<scope>"),
				ServiceURI: to.StringPtr("<service-uri>"),
				Status:     armcontainerregistry.WebhookStatus("enabled").ToPtr(),
			},
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.WebhooksClientUpdateResult)
}
