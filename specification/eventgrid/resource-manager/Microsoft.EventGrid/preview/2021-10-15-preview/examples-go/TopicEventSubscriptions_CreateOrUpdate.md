Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventgrid%2Farmeventgrid%2Fv0.5.0/sdk/resourcemanager/eventgrid/armeventgrid/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/TopicEventSubscriptions_CreateOrUpdate.json
func ExampleTopicEventSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventgrid.NewTopicEventSubscriptionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<topic-name>",
		"<event-subscription-name>",
		armeventgrid.EventSubscription{
			Properties: &armeventgrid.EventSubscriptionProperties{
				Destination: &armeventgrid.WebHookEventSubscriptionDestination{
					EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
					Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
						EndpointURL: to.Ptr("<endpoint-url>"),
					},
				},
				Filter: &armeventgrid.EventSubscriptionFilter{
					IsSubjectCaseSensitive: to.Ptr(false),
					SubjectBeginsWith:      to.Ptr("<subject-begins-with>"),
					SubjectEndsWith:        to.Ptr("<subject-ends-with>"),
				},
			},
		},
		&armeventgrid.TopicEventSubscriptionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
