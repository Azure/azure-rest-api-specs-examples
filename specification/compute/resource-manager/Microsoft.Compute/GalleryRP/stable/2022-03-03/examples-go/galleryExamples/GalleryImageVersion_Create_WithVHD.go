package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/GalleryImageVersion_Create_WithVHD.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryImageVersionUsingVhdAsASource() {
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
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
									Lun:                 to.Ptr[int32](1),
								}},
							OSDiskImage: &armcompute.OSDiskImageEncryption{
								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
							},
						},
						ExcludeFromLatest:    to.Ptr(false),
						RegionalReplicaCount: to.Ptr[int32](1),
					},
					{
						Name:                 to.Ptr("East US"),
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
							ID:  to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
							URI: to.Ptr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
						},
						Lun: to.Ptr[int32](1),
					}},
				OSDiskImage: &armcompute.GalleryOSDiskImage{
					HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
					Source: &armcompute.GalleryDiskImageSource{
						ID:  to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
						URI: to.Ptr("https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"),
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
