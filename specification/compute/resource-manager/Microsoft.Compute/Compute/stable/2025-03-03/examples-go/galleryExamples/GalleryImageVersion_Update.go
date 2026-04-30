package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/GalleryImageVersion_Update.json
func ExampleGalleryImageVersionsClient_BeginUpdate_updateASimpleGalleryImageVersionManagedImageAsSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImageVersionsClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", armcompute.GalleryImageVersionUpdate{
		Properties: &armcompute.GalleryImageVersionProperties{
			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name:                 to.Ptr("West US"),
						RegionalReplicaCount: to.Ptr[int32](1),
					},
					{
						Name:                 to.Ptr("East US"),
						RegionalReplicaCount: to.Ptr[int32](2),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
					},
				},
			},
			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
				Source: &armcompute.GalleryArtifactVersionFullSource{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryImageVersionsClientUpdateResponse{
	// 	GalleryImageVersion: &armcompute.GalleryImageVersion{
	// 		Properties: &armcompute.GalleryImageVersionProperties{
	// 			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 				TargetRegions: []*armcompute.TargetRegion{
	// 					{
	// 						Name: to.Ptr("West US"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 						Encryption: &armcompute.EncryptionImages{
	// 							OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
	// 							},
	// 							DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 								{
	// 									Lun: to.Ptr[int32](1),
	// 									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
	// 								},
	// 							},
	// 						},
	// 						ExcludeFromLatest: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("East US"),
	// 						RegionalReplicaCount: to.Ptr[int32](2),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardZRS),
	// 						ExcludeFromLatest: to.Ptr(false),
	// 					},
	// 				},
	// 				ReplicaCount: to.Ptr[int32](1),
	// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00Z"); return t}()),
	// 				StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			},
	// 			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
	// 				Source: &armcompute.GalleryArtifactVersionFullSource{
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
	// 				},
	// 				OSDiskImage: &armcompute.GalleryOSDiskImage{
	// 					SizeInGB: to.Ptr[int32](10),
	// 					HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
	// 				},
	// 				DataDiskImages: []*armcompute.GalleryDataDiskImage{
	// 					{
	// 						Lun: to.Ptr[int32](1),
	// 						SizeInGB: to.Ptr[int32](10),
	// 						HostCaching: to.Ptr(armcompute.HostCachingNone),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("1.0.0"),
	// 	},
	// }
}
