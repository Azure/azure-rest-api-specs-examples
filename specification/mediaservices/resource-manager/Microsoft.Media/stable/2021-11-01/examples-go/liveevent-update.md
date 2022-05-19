Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmediaservices%2Farmmediaservices%2Fv1.0.0/sdk/resourcemanager/mediaservices/armmediaservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-update.json
func ExampleLiveEventsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		armmediaservices.LiveEvent{
			Location: to.Ptr("West US"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
				"tag3": to.Ptr("value3"),
			},
			Properties: &armmediaservices.LiveEventProperties{
				Description: to.Ptr("test event updated"),
				Input: &armmediaservices.LiveEventInput{
					AccessControl: &armmediaservices.LiveEventInputAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.Ptr("AllowOne"),
									Address: to.Ptr("192.1.1.0"),
								}},
						},
					},
					KeyFrameIntervalDuration: to.Ptr("PT6S"),
					StreamingProtocol:        to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
				},
				Preview: &armmediaservices.LiveEventPreview{
					AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.Ptr("AllowOne"),
									Address: to.Ptr("192.1.1.0"),
								}},
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
