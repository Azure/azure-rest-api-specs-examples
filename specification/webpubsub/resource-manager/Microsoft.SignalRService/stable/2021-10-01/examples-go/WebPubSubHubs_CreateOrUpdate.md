Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwebpubsub%2Farmwebpubsub%2Fv0.4.0/sdk/resourcemanager/webpubsub/armwebpubsub/README.md) on how to add the SDK to your project and authenticate.

```go
package armwebpubsub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubHubs_CreateOrUpdate.json
func ExampleHubsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armwebpubsub.NewHubsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<hub-name>",
		"<resource-group-name>",
		"<resource-name>",
		armwebpubsub.Hub{
			Properties: &armwebpubsub.HubProperties{
				EventHandlers: []*armwebpubsub.EventHandler{
					{
						Auth: &armwebpubsub.UpstreamAuthSettings{
							Type: to.Ptr(armwebpubsub.UpstreamAuthTypeManagedIdentity),
							ManagedIdentity: &armwebpubsub.ManagedIdentitySettings{
								Resource: to.Ptr("<resource>"),
							},
						},
						SystemEvents: []*string{
							to.Ptr("connect"),
							to.Ptr("connected")},
						URLTemplate:      to.Ptr("<urltemplate>"),
						UserEventPattern: to.Ptr("<user-event-pattern>"),
					}},
			},
		},
		&armwebpubsub.HubsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
```
