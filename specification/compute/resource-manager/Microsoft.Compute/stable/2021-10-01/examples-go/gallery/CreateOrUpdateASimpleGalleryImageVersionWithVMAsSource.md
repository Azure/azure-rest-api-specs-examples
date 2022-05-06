Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryImageVersionWithVMAsSource.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewGalleryImageVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<gallery-name>",
		"<gallery-image-name>",
		"<gallery-image-version-name>",
		armcompute.GalleryImageVersion{
			Location: to.Ptr("<location>"),
			Properties: &armcompute.GalleryImageVersionProperties{
				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
					TargetRegions: []*armcompute.TargetRegion{
						{
							Name: to.Ptr("<name>"),
							Encryption: &armcompute.EncryptionImages{
								DataDiskImages: []*armcompute.DataDiskImageEncryption{
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](0),
									},
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](1),
									}},
								OSDiskImage: &armcompute.OSDiskImageEncryption{
									DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
								},
							},
							RegionalReplicaCount: to.Ptr[int32](1),
						},
						{
							Name: to.Ptr("<name>"),
							Encryption: &armcompute.EncryptionImages{
								DataDiskImages: []*armcompute.DataDiskImageEncryption{
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](0),
									},
									{
										DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
										Lun:                 to.Ptr[int32](1),
									}},
								OSDiskImage: &armcompute.OSDiskImageEncryption{
									DiskEncryptionSetID: to.Ptr("<disk-encryption-set-id>"),
								},
							},
							RegionalReplicaCount: to.Ptr[int32](2),
							StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
						}},
				},
				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
					Source: &armcompute.GalleryArtifactVersionSource{
						ID: to.Ptr("<id>"),
					},
				},
			},
		},
		&armcompute.GalleryImageVersionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
