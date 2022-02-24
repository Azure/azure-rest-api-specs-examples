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

// x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2021-12-01/examples/EventSubscriptions_CreateOrUpdateForCustomTopic_EventHubDestination.json
func ExampleEventSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventgrid.NewEventSubscriptionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<scope>",
		"<event-subscription-name>",
		armeventgrid.EventSubscription{
			Properties: &armeventgrid.EventSubscriptionProperties{
				DeadLetterDestination: &armeventgrid.StorageBlobDeadLetterDestination{
					EndpointType: armeventgrid.DeadLetterEndPointType("StorageBlob").ToPtr(),
					Properties: &armeventgrid.StorageBlobDeadLetterDestinationProperties{
						BlobContainerName: to.StringPtr("<blob-container-name>"),
						ResourceID:        to.StringPtr("<resource-id>"),
					},
				},
				Destination: &armeventgrid.EventHubEventSubscriptionDestination{
					EndpointType: armeventgrid.EndpointType("EventHub").ToPtr(),
					Properties: &armeventgrid.EventHubEventSubscriptionDestinationProperties{
						ResourceID: to.StringPtr("<resource-id>"),
					},
				},
				Filter: &armeventgrid.EventSubscriptionFilter{
					IsSubjectCaseSensitive: to.BoolPtr(false),
					SubjectBeginsWith:      to.StringPtr("<subject-begins-with>"),
					SubjectEndsWith:        to.StringPtr("<subject-ends-with>"),
				},
			},
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
