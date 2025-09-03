package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/MarketplaceGalleryImages_ListAll.json
func ExampleMarketplaceGalleryImagesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMarketplaceGalleryImagesClient().NewListAllPager(nil)
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
		// page = armazurestackhcivm.MarketplaceGalleryImagesClientListAllResponse{
		// 	MarketplaceGalleryImageListResult: armazurestackhcivm.MarketplaceGalleryImageListResult{
		// 		Value: []*armazurestackhcivm.MarketplaceGalleryImage{
		// 			{
		// 				Name: to.Ptr("test-marketplace-gallery-image"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/marketplaceGalleryImages"),
		// 				ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
		// 					Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/test-marketplace-gallery-image"),
		// 				Location: to.Ptr("West US2"),
		// 				Properties: &armazurestackhcivm.MarketplaceGalleryImageProperties{
		// 					CloudInitDataSource: to.Ptr(armazurestackhcivm.CloudInitDataSourceAzure),
		// 					ContainerID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
		// 					HyperVGeneration: to.Ptr(armazurestackhcivm.HyperVGenerationV1),
		// 					Identifier: &armazurestackhcivm.GalleryImageIdentifier{
		// 						Offer: to.Ptr("myOfferName"),
		// 						Publisher: to.Ptr("myPublisherName"),
		// 						SKU: to.Ptr("mySkuName"),
		// 					},
		// 					OSType: to.Ptr(armazurestackhcivm.OperatingSystemTypesWindows),
		// 					ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
		// 					Status: &armazurestackhcivm.MarketplaceGalleryImageStatus{
		// 						DownloadStatus: &armazurestackhcivm.MarketplaceGalleryImageStatusDownloadStatus{
		// 							DownloadSizeInMB: to.Ptr[int64](9383),
		// 						},
		// 						ProgressPercentage: to.Ptr[int64](100),
		// 						ProvisioningStatus: &armazurestackhcivm.MarketplaceGalleryImageStatusProvisioningStatus{
		// 							OperationID: to.Ptr("79cfc696-44f5-4a68-a620-21850f7e9fb0"),
		// 							Status: to.Ptr(armazurestackhcivm.StatusSucceeded),
		// 						},
		// 					},
		// 					Version: &armazurestackhcivm.GalleryImageVersion{
		// 						Name: to.Ptr("1.0.0"),
		// 						Properties: &armazurestackhcivm.GalleryImageVersionProperties{
		// 							StorageProfile: &armazurestackhcivm.GalleryImageVersionStorageProfile{
		// 								OSDiskImage: &armazurestackhcivm.GalleryOSDiskImage{
		// 									SizeInMB: to.Ptr[int64](30270),
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
