package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryImageVersion_Get_WithValidationProfileAndReplicationStatus.json
func ExampleGalleryImageVersionsClient_Get_getAGalleryImageVersionWithValidationProfileAndReplicationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryImageVersionsClient().Get(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", &armcompute.GalleryImageVersionsClientGetOptions{Expand: to.Ptr(armcompute.ReplicationStatusTypes("ValidationProfile,ReplicationStatus"))})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GalleryImageVersion = armcompute.GalleryImageVersion{
	// 	Name: to.Ptr("1.0.0"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName/Versions/1.0.0"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryImageVersionProperties{
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
	// 			ReplicaCount: to.Ptr[int32](1),
	// 			StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			TargetRegions: []*armcompute.TargetRegion{
	// 				{
	// 					Name: to.Ptr("West US"),
	// 					Encryption: &armcompute.EncryptionImages{
	// 						DataDiskImages: []*armcompute.DataDiskImageEncryption{
	// 							{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
	// 								Lun: to.Ptr[int32](0),
	// 							},
	// 							{
	// 								DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
	// 								Lun: to.Ptr[int32](1),
	// 						}},
	// 						OSDiskImage: &armcompute.OSDiskImageEncryption{
	// 							DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
	// 						},
	// 					},
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](1),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 				},
	// 				{
	// 					Name: to.Ptr("East US"),
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](2),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			}},
	// 		},
	// 		ReplicationStatus: &armcompute.ReplicationStatus{
	// 			AggregatedState: to.Ptr(armcompute.AggregatedReplicationStateCompleted),
	// 			Summary: []*armcompute.RegionalReplicationStatus{
	// 				{
	// 					Progress: to.Ptr[int32](100),
	// 					Region: to.Ptr("West US"),
	// 					State: to.Ptr(armcompute.ReplicationStateCompleted),
	// 					Details: to.Ptr(""),
	// 				},
	// 				{
	// 					Progress: to.Ptr[int32](100),
	// 					Region: to.Ptr("East US"),
	// 					State: to.Ptr(armcompute.ReplicationStateCompleted),
	// 					Details: to.Ptr(""),
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
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
	// 			},
	// 		},
	// 		ValidationsProfile: &armcompute.ValidationsProfile{
	// 			ExecutedValidations: []*armcompute.ExecutedValidation{
	// 				{
	// 					Type: to.Ptr("LinuxSecurityBaseline"),
	// 					ExecutionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
	// 					Status: to.Ptr(armcompute.ValidationStatusFailed),
	// 					Version: to.Ptr("beta"),
	// 				},
	// 				{
	// 					Type: to.Ptr("AzCertify"),
	// 					ExecutionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
	// 					Status: to.Ptr(armcompute.ValidationStatusSucceeded),
	// 					Version: to.Ptr("10.0.4"),
	// 			}},
	// 			PlatformAttributes: []*armcompute.PlatformAttribute{
	// 				{
	// 					Name: to.Ptr("source_image_publisher"),
	// 					Value: to.Ptr("publishValue"),
	// 				},
	// 				{
	// 					Name: to.Ptr("source_image_offer"),
	// 					Value: to.Ptr("OfferValue"),
	// 				},
	// 				{
	// 					Name: to.Ptr("source_image_sku"),
	// 					Value: to.Ptr("SkuValue"),
	// 				},
	// 				{
	// 					Name: to.Ptr("source_image_version"),
	// 					Value: to.Ptr("1.0.0"),
	// 			}},
	// 			ValidationEtag: to.Ptr("2018-01-01T00:00:00Z"),
	// 		},
	// 	},
	// }
}
