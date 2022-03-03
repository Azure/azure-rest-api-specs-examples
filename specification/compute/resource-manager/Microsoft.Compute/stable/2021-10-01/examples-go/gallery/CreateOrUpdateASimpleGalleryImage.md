Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.4.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryImage.json
func ExampleGalleryImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImagesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		armcompute.GalleryImage{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.GalleryImageProperties{
				HyperVGeneration: armcompute.HyperVGeneration("V1").ToPtr(),
				Identifier: &armcompute.GalleryImageIdentifier{
					Offer:     to.StringPtr("<offer>"),
					Publisher: to.StringPtr("<publisher>"),
					SKU:       to.StringPtr("<sku>"),
				},
				OSState: armcompute.OperatingSystemStateTypesGeneralized.ToPtr(),
				OSType:  armcompute.OperatingSystemTypesWindows.ToPtr(),
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
	log.Printf("Response result: %#v\n", res.GalleryImagesClientCreateOrUpdateResult)
}
```
