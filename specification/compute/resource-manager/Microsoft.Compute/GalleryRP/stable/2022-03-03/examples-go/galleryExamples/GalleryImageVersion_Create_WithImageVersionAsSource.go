package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/GalleryImageVersion_Create_WithImageVersionAsSource.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryImageVersionUsingSharedImageAsSource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImageVersionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", armcompute.GalleryImageVersion{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryImageVersionProperties{
			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name: to.Ptr("West US"),
						Encryption: &armcompute.EncryptionImages{
							DataDiskImages: []*armcompute.DataDiskImageEncryption{
								{
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet"),
									Lun:                 to.Ptr[int32](0),
								},
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
									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet"),
									Lun:                 to.Ptr[int32](0),
								},
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
				Source: &armcompute.GalleryArtifactVersionFullSource{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageDefinitionName}/versions/{versionName}"),
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
	// res.GalleryImageVersion = armcompute.GalleryImageVersion{
	// 	Name: to.Ptr("1.0.0"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryImageVersionProperties{
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00Z"); return t}()),
	// 			ReplicaCount: to.Ptr[int32](1),
	// 			StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			TargetRegions: []*armcompute.TargetRegion{
	// 				{
	// 					Name: to.Ptr("West US"),
	// 					Encryption: &armcompute.EncryptionImages{
	// 						DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 							{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet"),
	// 								Lun: to.Ptr[int32](0),
	// 							},
	// 							{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
	// 								Lun: to.Ptr[int32](1),
	// 						}},
	// 						OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 							DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"),
	// 						},
	// 					},
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](1),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 				},
	// 				{
	// 					Name: to.Ptr("East US"),
	// 					Encryption: &armcompute.EncryptionImages{
	// 						DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 							{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet"),
	// 								Lun: to.Ptr[int32](0),
	// 							},
	// 							{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
	// 								Lun: to.Ptr[int32](1),
	// 						}},
	// 						OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 							DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"),
	// 						},
	// 					},
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](2),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardZRS),
	// 			}},
	// 		},
	// 		SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
	// 			AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 			PolicyViolations: []*armcompute.PolicyViolation{
	// 				{
	// 					Category: to.Ptr(armcompute.PolicyViolationCategoryImageFlaggedUnsafe),
	// 					Details: to.Ptr("This is the policy violation details."),
	// 			}},
	// 			ReportedForPolicyViolation: to.Ptr(true),
	// 		},
	// 		StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
	// 			DataDiskImages: []*armcompute.GalleryDataDiskImage{
	// 				{
	// 					HostCaching: to.Ptr(armcompute.HostCachingNone),
	// 					SizeInGB: to.Ptr[int32](10),
	// 					Lun: to.Ptr[int32](1),
	// 			}},
	// 			OSDiskImage: &armcompute.GalleryOSDiskImage{
	// 				HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
	// 				SizeInGB: to.Ptr[int32](10),
	// 			},
	// 			Source: &armcompute.GalleryArtifactVersionFullSource{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/galleries/{galleryName}/images/{imageDefinitionName}/versions/{versionName}"),
	// 			},
	// 		},
	// 	},
	// }
}
