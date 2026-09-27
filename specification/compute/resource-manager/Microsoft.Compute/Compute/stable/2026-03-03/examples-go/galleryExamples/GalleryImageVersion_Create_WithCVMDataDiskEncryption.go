package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-03/galleryExamples/GalleryImageVersion_Create_WithCVMDataDiskEncryption.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate_createOrUpdateAGalleryImageVersionWithCvmDataDiskEncryptionUsingCustomerManagedKey() {
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
		Location: to.Ptr("eastus"),
		Properties: &armcompute.GalleryImageVersionProperties{
			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name:                 to.Ptr("eastus"),
						RegionalReplicaCount: to.Ptr[int32](1),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardZRS),
						Encryption: &armcompute.EncryptionImages{
							OSDiskImage: &armcompute.OSDiskImageEncryption{
								SecurityProfile: &armcompute.OSDiskImageSecurityProfile{
									ConfidentialVMEncryptionType: to.Ptr(armcompute.ConfidentialVMEncryptionTypeEncryptedWithPmk),
								},
							},
							DataDiskImages: []*armcompute.DataDiskImageEncryption{
								{
									Lun: to.Ptr[int32](0),
									SecurityProfile: &armcompute.DataDiskImageSecurityProfile{
										ConfidentialVMEncryptionType: to.Ptr(armcompute.ConfidentialVMEncryptionTypeDataDiskEncryptedWithCmk),
										SecureVMDiskEncryptionSetID:  to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
									},
								},
							},
						},
						ExcludeFromLatest: to.Ptr(false),
					},
				},
				ReplicaCount:      to.Ptr[int32](1),
				ExcludeFromLatest: to.Ptr(false),
				ReplicationMode:   to.Ptr(armcompute.ReplicationModeFull),
			},
			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
				Source: &armcompute.GalleryArtifactVersionFullSource{
					VirtualMachineID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
				},
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
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/images/myGalleryImageName/versions/1.0.0"),
	// 		Properties: &armcompute.GalleryImageVersionProperties{
	// 			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 				TargetRegions: []*armcompute.TargetRegion{
	// 					{
	// 						Name: to.Ptr("eastus"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardZRS),
	// 						Encryption: &armcompute.EncryptionImages{
	// 							OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 								SecurityProfile: &armcompute.OSDiskImageSecurityProfile{
	// 									ConfidentialVMEncryptionType: to.Ptr(armcompute.ConfidentialVMEncryptionTypeEncryptedWithPmk),
	// 								},
	// 							},
	// 							DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 								{
	// 									Lun: to.Ptr[int32](0),
	// 									SecurityProfile: &armcompute.DataDiskImageSecurityProfile{
	// 										ConfidentialVMEncryptionType: to.Ptr(armcompute.ConfidentialVMEncryptionTypeDataDiskEncryptedWithCmk),
	// 										SecureVMDiskEncryptionSetID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 									},
	// 								},
	// 							},
	// 						},
	// 						ExcludeFromLatest: to.Ptr(false),
	// 					},
	// 				},
	// 				ReplicaCount: to.Ptr[int32](1),
	// 				PublishedDate: to.Ptr(time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 				ExcludeFromLatest: to.Ptr(false),
	// 				ReplicationMode: to.Ptr(armcompute.ReplicationModeFull),
	// 			},
	// 			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
	// 				Source: &armcompute.GalleryArtifactVersionFullSource{
	// 					VirtualMachineID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 				},
	// 			},
	// 			SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
	// 				ReportedForPolicyViolation: to.Ptr(false),
	// 				AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Name: to.Ptr("1.0.0"),
	// 	},
	// }
}
