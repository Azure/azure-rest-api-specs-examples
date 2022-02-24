Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventgrid%2Farmeventgrid%2Fv0.3.1/sdk/resourcemanager/eventgrid/armeventgrid/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventgrid_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/SystemTopicEventSubscriptions_Update.json
func ExampleSystemTopicEventSubscriptionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventgrid.NewSystemTopicEventSubscriptionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<system-topic-name>",
		"<event-subscription-name>",
		armeventgrid.EventSubscriptionUpdateParameters{
			Destination: &armeventgrid.WebHookEventSubscriptionDestination{
				EndpointType: armeventgrid.EndpointType("WebHook").ToPtr(),
				Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
					EndpointURL: to.StringPtr("<endpoint-url>"),
				},
			},
			Filter: &armeventgrid.EventSubscriptionFilter{
				IsSubjectCaseSensitive: to.BoolPtr(true),
				SubjectBeginsWith:      to.StringPtr("<subject-begins-with>"),
				SubjectEndsWith:        to.StringPtr("<subject-ends-with>"),
			},
			Labels: []*string{
				to.StringPtr("label1"),
				to.StringPtr("label2")},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
