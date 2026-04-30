package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/GalleryImageVersion_ListByGalleryImage.json
func ExampleGalleryImageVersionsClient_NewListByGalleryImagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleryImageVersionsClient().NewListByGalleryImagePager("myResourceGroup", "myGalleryName", "myGalleryImageName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armcompute.GalleryImageVersionsClientListByGalleryImageResponse{
		// 	GalleryImageVersionList: armcompute.GalleryImageVersionList{
		// 		Value: []*armcompute.GalleryImageVersion{
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName/Versions/1.0.0"),
		// 				Properties: &armcompute.GalleryImageVersionProperties{
		// 					PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
		// 						TargetRegions: []*armcompute.TargetRegion{
		// 							{
		// 								Name: to.Ptr("West US"),
		// 								RegionalReplicaCount: to.Ptr[int32](1),
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 								Encryption: &armcompute.EncryptionImages{
		// 									OSDiskImage: &armcompute.OSDiskImageEncryption{
		// 										DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
		// 									},
		// 									DataDiskImages: []*armcompute.DataDiskImageEncryption{
		// 										{
		// 											Lun: to.Ptr[int32](0),
		// 											DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
		// 										},
		// 										{
		// 											Lun: to.Ptr[int32](1),
		// 											DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
		// 										},
		// 									},
		// 								},
		// 								ExcludeFromLatest: to.Ptr(false),
		// 							},
		// 							{
		// 								Name: to.Ptr("East US"),
		// 								RegionalReplicaCount: to.Ptr[int32](2),
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 								ExcludeFromLatest: to.Ptr(false),
		// 							},
		// 						},
		// 						ReplicaCount: to.Ptr[int32](1),
		// 						PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00Z"); return t}()),
		// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 					},
		// 					StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
		// 						Source: &armcompute.GalleryArtifactVersionFullSource{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
		// 						},
		// 						OSDiskImage: &armcompute.GalleryOSDiskImage{
		// 							SizeInGB: to.Ptr[int32](10),
		// 							HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
		// 						},
		// 						DataDiskImages: []*armcompute.GalleryDataDiskImage{
		// 							{
		// 								Lun: to.Ptr[int32](1),
		// 								SizeInGB: to.Ptr[int32](10),
		// 								HostCaching: to.Ptr(armcompute.HostCachingNone),
		// 							},
		// 						},
		// 					},
		// 					SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
		// 						AllowDeletionOfReplicatedLocations: to.Ptr(false),
		// 						ReportedForPolicyViolation: to.Ptr(true),
		// 						PolicyViolations: []*armcompute.PolicyViolation{
		// 							{
		// 								Category: to.Ptr(armcompute.PolicyViolationCategoryImageFlaggedUnsafe),
		// 								Details: to.Ptr("This is the policy violation details."),
		// 							},
		// 						},
		// 						BlockDeletionBeforeEndOfLife: to.Ptr(false),
		// 					},
		// 					ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				Name: to.Ptr("1.0.0"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://svchost:99/subscriptions/subscriptionId/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/images/myGalleryImageName/versions?$skiptoken=token/Subscriptions/subscriptionId/ResourceGroups/myResourceGroup/galleries/myGalleryName/images/myGalleryImageName/versions/myGalleryImageVersionName"),
		// 	},
		// }
	}
}
