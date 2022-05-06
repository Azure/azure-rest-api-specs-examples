Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv0.5.0/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubCreate.json
func ExampleEventHubsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventhub.NewEventHubsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<event-hub-name>",
		armeventhub.Eventhub{
			Properties: &armeventhub.Properties{
				CaptureDescription: &armeventhub.CaptureDescription{
					Destination: &armeventhub.Destination{
						Name: to.Ptr("<name>"),
						Properties: &armeventhub.DestinationProperties{
							ArchiveNameFormat:        to.Ptr("<archive-name-format>"),
							BlobContainer:            to.Ptr("<blob-container>"),
							StorageAccountResourceID: to.Ptr("<storage-account-resource-id>"),
						},
					},
					Enabled:           to.Ptr(true),
					Encoding:          to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
					IntervalInSeconds: to.Ptr[int32](120),
					SizeLimitInBytes:  to.Ptr[int32](10485763),
				},
				MessageRetentionInDays: to.Ptr[int64](4),
				PartitionCount:         to.Ptr[int64](4),
				Status:                 to.Ptr(armeventhub.EntityStatusActive),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
