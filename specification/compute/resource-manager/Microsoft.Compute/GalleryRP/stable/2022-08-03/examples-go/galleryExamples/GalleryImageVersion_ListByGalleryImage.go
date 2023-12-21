package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/60679ee3db06e93eb73faa0587fed93ed843d6dc/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-08-03/examples/galleryExamples/GalleryImageVersion_ListByGalleryImage.json
func ExampleGalleryImageVersionsClient_NewListByGalleryImagePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.GalleryImageVersionList = armcompute.GalleryImageVersionList{
		// 	Value: []*armcompute.GalleryImageVersion{
		// 		{
		// 			Name: to.Ptr("1.0.0"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName/Versions/1.0.0"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryImageVersionProperties{
		// 				ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
		// 					PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00.000Z"); return t}()),
		// 					ReplicaCount: to.Ptr[int32](1),
		// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 					TargetRegions: []*armcompute.TargetRegion{
		// 						{
		// 							Name: to.Ptr("West US"),
		// 							Encryption: &armcompute.EncryptionImages{
		// 								DataDiskImages: []*armcompute.DataDiskImageEncryption{
		// 									{
		// 										DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet"),
		// 										Lun: to.Ptr[int32](0),
		// 									},
		// 									{
		// 										DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
		// 										Lun: to.Ptr[int32](1),
		// 								}},
		// 								OSDiskImage: &armcompute.OSDiskImageEncryption{
		// 									DiskEncryptionSetID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"),
		// 								},
		// 							},
		// 							ExcludeFromLatest: to.Ptr(false),
		// 							RegionalReplicaCount: to.Ptr[int32](1),
		// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 						},
		// 						{
		// 							Name: to.Ptr("East US"),
		// 							ExcludeFromLatest: to.Ptr(false),
		// 							RegionalReplicaCount: to.Ptr[int32](2),
		// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 					}},
		// 				},
		// 				SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
		// 					AllowDeletionOfReplicatedLocations: to.Ptr(false),
		// 					PolicyViolations: []*armcompute.PolicyViolation{
		// 						{
		// 							Category: to.Ptr(armcompute.PolicyViolationCategoryImageFlaggedUnsafe),
		// 							Details: to.Ptr("This is the policy violation details."),
		// 					}},
		// 					ReportedForPolicyViolation: to.Ptr(true),
		// 				},
		// 				StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
		// 					DataDiskImages: []*armcompute.GalleryDataDiskImage{
		// 						{
		// 							HostCaching: to.Ptr(armcompute.HostCachingNone),
		// 							SizeInGB: to.Ptr[int32](10),
		// 							Lun: to.Ptr[int32](1),
		// 					}},
		// 					OSDiskImage: &armcompute.GalleryOSDiskImage{
		// 						HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
		// 						SizeInGB: to.Ptr[int32](10),
		// 					},
		// 					Source: &armcompute.GalleryArtifactVersionFullSource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
