Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fwebpubsub%2Farmwebpubsub%2Fv0.2.0/sdk/resourcemanager/webpubsub/armwebpubsub/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubHubs_CreateOrUpdate.json
func ExampleHubsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armwebpubsub.NewHubsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<hub-name>",
		"<resource-group-name>",
		"<resource-name>",
		armwebpubsub.Hub{
			Properties: &armwebpubsub.HubProperties{
				EventHandlers: []*armwebpubsub.EventHandler{
					{
						Auth: &armwebpubsub.UpstreamAuthSettings{
							Type: armwebpubsub.UpstreamAuthType("ManagedIdentity").ToPtr(),
							ManagedIdentity: &armwebpubsub.ManagedIdentitySettings{
								Resource: to.StringPtr("<resource>"),
							},
						},
						SystemEvents: []*string{
							to.StringPtr("connect"),
							to.StringPtr("connected")},
						URLTemplate:      to.StringPtr("<urltemplate>"),
						UserEventPattern: to.StringPtr("<user-event-pattern>"),
					}},
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
	log.Printf("Response result: %#v\n", res.HubsClientCreateOrUpdateResult)
}
```
