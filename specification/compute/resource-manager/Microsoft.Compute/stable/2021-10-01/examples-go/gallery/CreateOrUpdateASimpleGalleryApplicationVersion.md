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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryApplicationVersion.json
func ExampleGalleryApplicationVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewGalleryApplicationVersionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-application-name>",
		"<gallery-application-version-name>",
		armcompute.GalleryApplicationVersion{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.GalleryApplicationVersionProperties{
				PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
					EndOfLifeDate:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00Z"); return t }()),
					ReplicaCount:       to.Int32Ptr(1),
					StorageAccountType: armcompute.StorageAccountType("Standard_LRS").ToPtr(),
					TargetRegions: []*armcompute.TargetRegion{
						{
							Name:                 to.StringPtr("<name>"),
							RegionalReplicaCount: to.Int32Ptr(1),
							StorageAccountType:   armcompute.StorageAccountType("Standard_LRS").ToPtr(),
						}},
					ManageActions: &armcompute.UserArtifactManage{
						Install: to.StringPtr("<install>"),
						Remove:  to.StringPtr("<remove>"),
					},
					Source: &armcompute.UserArtifactSource{
						MediaLink: to.StringPtr("<media-link>"),
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
	log.Printf("Response result: %#v\n", res.GalleryApplicationVersionsClientCreateOrUpdateResult)
}
```
