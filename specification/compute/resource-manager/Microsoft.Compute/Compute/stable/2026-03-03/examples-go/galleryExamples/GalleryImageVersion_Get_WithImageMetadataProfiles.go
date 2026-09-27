package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-03/galleryExamples/GalleryImageVersion_Get_WithImageMetadataProfiles.json
func ExampleGalleryImageVersionsClient_Get_getAGalleryImageVersionWithImageMetadataProfiles() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryImageVersionsClient().Get(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryImageVersionsClientGetResponse{
	// 	GalleryImageVersion: armcompute.GalleryImageVersion{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/images/myGalleryImageName/versions/1.0.0"),
	// 		Properties: &armcompute.GalleryImageVersionProperties{
	// 			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 				TargetRegions: []*armcompute.TargetRegion{
	// 					{
	// 						Name: to.Ptr("West US"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 						ExcludeFromLatest: to.Ptr(false),
	// 					},
	// 				},
	// 				ReplicaCount: to.Ptr[int32](1),
	// 				PublishedDate: to.Ptr(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)),
	// 				StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			},
	// 			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
	// 				OSDiskImage: &armcompute.GalleryOSDiskImage{
	// 					SizeInGB: to.Ptr[int32](30),
	// 					HostCaching: to.Ptr(armcompute.HostCachingReadOnly),
	// 					Source: &armcompute.GalleryDiskImageSource{
	// 						StorageAccountID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"),
	// 						URI: to.Ptr("https://gallerysourcencus.blob.core.windows.net/myvhds/Linux-VM-2024.vhd"),
	// 					},
	// 				},
	// 			},
	// 			SecurityProfile: &armcompute.ImageVersionSecurityProfile{
	// 				SecretsProvisioningSettings: &armcompute.SecretsProvisioningSettings{
	// 					IsSupported: to.Ptr(true),
	// 					OSName: to.Ptr("mariner"),
	// 					Components: []*armcompute.SecretsProvisioningComponent{
	// 						{
	// 							Name: to.Ptr(armcompute.SecretsProvisioningComponentNameAzureGuestAgent),
	// 							Version: to.Ptr("2.7.0"),
	// 						},
	// 						{
	// 							Name: to.Ptr(armcompute.SecretsProvisioningComponentNameSecretsProvisioningLibrary),
	// 							Version: to.Ptr("1.0.0"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ImageMetadataProfiles: []*armcompute.ImageMetadataProfile{
	// 				{
	// 					Type: to.Ptr(armcompute.MetadataTypeUserProvidedSecretsProvisioningMetadata),
	// 					PublicMetadataList: []*armcompute.MetadataKeyValue{
	// 						{
	// 							MetadataKey: to.Ptr("Linux.AzureSecretsProvisioning.Enabled"),
	// 							MetadataValue: to.Ptr("true"),
	// 						},
	// 						{
	// 							MetadataKey: to.Ptr("OS.Name"),
	// 							MetadataValue: to.Ptr("mariner"),
	// 						},
	// 						{
	// 							MetadataKey: to.Ptr("AzureGuestAgent.Version"),
	// 							MetadataValue: to.Ptr("2.7.0"),
	// 						},
	// 						{
	// 							MetadataKey: to.Ptr("SecretsProvisioningLibrary.Version"),
	// 							MetadataValue: to.Ptr("1.0.0"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("1.0.0"),
	// 	},
	// }
}
