package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/GalleryImageVersion_Create_WithSnapshotsAsSource.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryImageVersionUsingMixOfDisksAndSnapshotsAsASource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGalleryImageVersionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", armcompute.GalleryImageVersion{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryImageVersionProperties{
			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name: to.Ptr("West US"),
						Encryption: &armcompute.EncryptionImages{
							DataDiskImages: []*armcompute.DataDiskImageEncryption{
								{
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
									Lun:                 to.Ptr[int32](1),
								}},
							OSDiskImage: &armcompute.OSDiskImageEncryption{
								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
							},
						},
						ExcludeFromLatest:    to.Ptr(false),
						RegionalReplicaCount: to.Ptr[int32](1),
					},
					{
						Name: to.Ptr("East US"),
						Encryption: &armcompute.EncryptionImages{
							DataDiskImages: []*armcompute.DataDiskImageEncryption{
								{
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
									Lun:                 to.Ptr[int32](1),
								}},
							OSDiskImage: &armcompute.OSDiskImageEncryption{
								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
							},
						},
						ExcludeFromLatest:    to.Ptr(false),
						RegionalReplicaCount: to.Ptr[int32](2),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
					}},
			},
			SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
				AllowDeletionOfReplicatedLocations: to.Ptr(false),
			},
			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
				DataDiskImages: []*armcompute.GalleryDataDiskImage{
					{
						HostCaching: to.Ptr(armcompute.HostCachingNone),
						Source: &armcompute.GalleryDiskImageSource{
							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/disks/{dataDiskName}"),
						},
						Lun: to.Ptr[int32](1),
					}},
				OSDiskImage: &armcompute.GalleryOSDiskImage{
					HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
					Source: &armcompute.GalleryDiskImageSource{
						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/snapshots/{osSnapshotName}"),
					},
				},
			},
		},
	}, nil)
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
