package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-12-03/galleryExamples/GalleryImageVersion_Create_WithImageVersionAsSource.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryImageVersionUsingSharedImageAsSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImageVersionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", armcompute.GalleryImageVersion{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryImageVersionProperties{
			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name:                 to.Ptr("West US"),
						RegionalReplicaCount: to.Ptr[int32](1),
						Encryption: &armcompute.EncryptionImages{
							OSDiskImage: &armcompute.OSDiskImageEncryption{
								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
							},
							DataDiskImages: []*armcompute.DataDiskImageEncryption{
								{
									Lun:                 to.Ptr[int32](0),
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet"),
								},
								{
									Lun:                 to.Ptr[int32](1),
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
								},
							},
						},
						ExcludeFromLatest: to.Ptr(false),
					},
					{
						Name:                 to.Ptr("East US"),
						RegionalReplicaCount: to.Ptr[int32](2),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
						Encryption: &armcompute.EncryptionImages{
							OSDiskImage: &armcompute.OSDiskImageEncryption{
								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
							},
							DataDiskImages: []*armcompute.DataDiskImageEncryption{
								{
									Lun:                 to.Ptr[int32](0),
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet"),
								},
								{
									Lun:                 to.Ptr[int32](1),
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
								},
							},
						},
						ExcludeFromLatest: to.Ptr(false),
					},
				},
			},
			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
				Source: &armcompute.GalleryArtifactVersionFullSource{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageDefinitionName}/versions/{versionName}"),
				},
			},
			SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
				AllowDeletionOfReplicatedLocations: to.Ptr(false),
				BlockDeletionBeforeEndOfLife:       to.Ptr(false),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryImageVersionsClientCreateOrUpdateResponse{
	// 	GalleryImageVersion: armcompute.GalleryImageVersion{
	// 		ID: to.Ptr("/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName/Versions/1.0.0"),
	// 		Properties: &armcompute.GalleryImageVersionProperties{
	// 			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 				TargetRegions: []*armcompute.TargetRegion{
	// 					{
	// 						Name: to.Ptr("West US"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 						Encryption: &armcompute.EncryptionImages{
	// 							OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
	// 							},
	// 							DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 								{
	// 									Lun: to.Ptr[int32](0),
	// 									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet"),
	// 								},
	// 								{
	// 									Lun: to.Ptr[int32](1),
	// 									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
	// 								},
	// 							},
	// 						},
	// 						ExcludeFromLatest: to.Ptr(false),
	// 					},
	// 					{
	// 						Name: to.Ptr("East US"),
	// 						RegionalReplicaCount: to.Ptr[int32](2),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardZRS),
	// 						Encryption: &armcompute.EncryptionImages{
	// 							OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
	// 							},
	// 							DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 								{
	// 									Lun: to.Ptr[int32](0),
	// 									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet"),
	// 								},
	// 								{
	// 									Lun: to.Ptr[int32](1),
	// 									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
	// 								},
	// 							},
	// 						},
	// 						ExcludeFromLatest: to.Ptr(false),
	// 					},
	// 				},
	// 				ReplicaCount: to.Ptr[int32](1),
	// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00Z"); return t}()),
	// 				StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			},
	// 			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
	// 				Source: &armcompute.GalleryArtifactVersionFullSource{
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageDefinitionName}/versions/{versionName}"),
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
	// 			SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
	// 				AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 				ReportedForPolicyViolation: to.Ptr(true),
	// 				PolicyViolations: []*armcompute.PolicyViolation{
	// 					{
	// 						Category: to.Ptr(armcompute.PolicyViolationCategoryImageFlaggedUnsafe),
	// 						Details: to.Ptr("This is the policy violation details."),
	// 					},
	// 				},
	// 				BlockDeletionBeforeEndOfLife: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("1.0.0"),
	// 	},
	// }
}
