Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv0.3.1/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubCreate.json
func ExampleEventHubsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventhub.NewEventHubsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		"<event-hub-name>",
		armeventhub.Eventhub{
			Properties: &armeventhub.Properties{
				CaptureDescription: &armeventhub.CaptureDescription{
					Destination: &armeventhub.Destination{
						Name: to.StringPtr("<name>"),
						Properties: &armeventhub.DestinationProperties{
							ArchiveNameFormat:        to.StringPtr("<archive-name-format>"),
							BlobContainer:            to.StringPtr("<blob-container>"),
							StorageAccountResourceID: to.StringPtr("<storage-account-resource-id>"),
						},
					},
					Enabled:           to.BoolPtr(true),
					Encoding:          armeventhub.EncodingCaptureDescriptionAvro.ToPtr(),
					IntervalInSeconds: to.Int32Ptr(120),
					SizeLimitInBytes:  to.Int32Ptr(10485763),
				},
				MessageRetentionInDays: to.Int64Ptr(4),
				PartitionCount:         to.Int64Ptr(4),
				Status:                 armeventhub.EntityStatusActive.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EventHubsClientCreateOrUpdateResult)
}
```
