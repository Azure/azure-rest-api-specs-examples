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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/UpdateASimpleGalleryImageVersion.json
func ExampleGalleryImageVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		armcompute.GalleryImageVersionUpdate{
			Properties: &armcompute.GalleryImageVersionProperties{
				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
					TargetRegions: []*armcompute.TargetRegion{
						{
							Name:                 to.StringPtr("<name>"),
							RegionalReplicaCount: to.Int32Ptr(1),
						},
						{
							Name:                 to.StringPtr("<name>"),
							RegionalReplicaCount: to.Int32Ptr(2),
							StorageAccountType:   armcompute.StorageAccountType("Standard_ZRS").ToPtr(),
						}},
				},
				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
					Source: &armcompute.GalleryArtifactVersionSource{
						ID: to.StringPtr("<id>"),
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
	log.Printf("Response result: %#v\n", res.GalleryImageVersionsClientUpdateResult)
}
```
