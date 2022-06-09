```go
package armmediaservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/liveevent-update.json
func ExampleLiveEventsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewLiveEventsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<live-event-name>",
		armmediaservices.LiveEvent{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
				"tag3": to.StringPtr("value3"),
			},
			Properties: &armmediaservices.LiveEventProperties{
				Description: to.StringPtr("<description>"),
				Input: &armmediaservices.LiveEventInput{
					AccessControl: &armmediaservices.LiveEventInputAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.StringPtr("<name>"),
									Address: to.StringPtr("<address>"),
								}},
						},
					},
					KeyFrameIntervalDuration: to.StringPtr("<key-frame-interval-duration>"),
					StreamingProtocol:        armmediaservices.LiveEventInputProtocol("FragmentedMP4").ToPtr(),
				},
				Preview: &armmediaservices.LiveEventPreview{
					AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.StringPtr("<name>"),
									Address: to.StringPtr("<address>"),
								}},
						},
					},
				},
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
	log.Printf("Response result: %#v\n", res.LiveEventsClientUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv0.3.1/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.
