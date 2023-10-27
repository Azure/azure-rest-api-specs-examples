package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutMarketplaceGalleryImage.json
func ExampleMarketplaceGalleryImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewMarketplaceGalleryImagesClient().BeginCreateOrUpdate(ctx, "test-rg", "test-marketplace-gallery-image", armazurestackhci.MarketplaceGalleryImages{
		Location: to.Ptr("West US2"),
		ExtendedLocation: &armazurestackhci.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
			Type: to.Ptr(armazurestackhci.ExtendedLocationTypesCustomLocation),
		},
		Properties: &armazurestackhci.MarketplaceGalleryImageProperties{
			CloudInitDataSource: to.Ptr(armazurestackhci.CloudInitDataSourceAzure),
			ContainerID:         to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
			HyperVGeneration:    to.Ptr(armazurestackhci.HyperVGenerationV1),
			Identifier: &armazurestackhci.GalleryImageIdentifier{
				Offer:     to.Ptr("myOfferName"),
				Publisher: to.Ptr("myPublisherName"),
				SKU:       to.Ptr("mySkuName"),
			},
			OSType: to.Ptr(armazurestackhci.OperatingSystemTypesWindows),
			Version: &armazurestackhci.GalleryImageVersion{
				Name: to.Ptr("1.0.0"),
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
	// res.MarketplaceGalleryImages = armazurestackhci.MarketplaceGalleryImages{
	// 	Name: to.Ptr("test-marketplace-gallery-image"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/marketplaceGalleryImages"),
	// 	ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/test-marketplace-gallery-image"),
	// 	Location: to.Ptr("West US2"),
	// 	ExtendedLocation: &armazurestackhci.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 		Type: to.Ptr(armazurestackhci.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurestackhci.MarketplaceGalleryImageProperties{
	// 		CloudInitDataSource: to.Ptr(armazurestackhci.CloudInitDataSourceAzure),
	// 		ContainerID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
	// 		HyperVGeneration: to.Ptr(armazurestackhci.HyperVGenerationV1),
	// 		Identifier: &armazurestackhci.GalleryImageIdentifier{
	// 			Offer: to.Ptr("myOfferName"),
	// 			Publisher: to.Ptr("myPublisherName"),
	// 			SKU: to.Ptr("mySkuName"),
	// 		},
	// 		OSType: to.Ptr(armazurestackhci.OperatingSystemTypesWindows),
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateEnumSucceeded),
	// 		Status: &armazurestackhci.MarketplaceGalleryImageStatus{
	// 			DownloadStatus: &armazurestackhci.MarketplaceGalleryImageStatusDownloadStatus{
	// 				DownloadSizeInMB: to.Ptr[int64](9383),
	// 			},
	// 			ProgressPercentage: to.Ptr[int64](0),
	// 			ProvisioningStatus: &armazurestackhci.MarketplaceGalleryImageStatusProvisioningStatus{
	// 				OperationID: to.Ptr("79cfc696-44f5-4a68-a620-21850f7e9fb0"),
	// 				Status: to.Ptr(armazurestackhci.StatusInProgress),
	// 			},
	// 		},
	// 		Version: &armazurestackhci.GalleryImageVersion{
	// 			Name: to.Ptr("1.0.0"),
	// 			Properties: &armazurestackhci.GalleryImageVersionProperties{
	// 				StorageProfile: &armazurestackhci.GalleryImageVersionStorageProfile{
	// 					OSDiskImage: &armazurestackhci.GalleryOSDiskImage{
	// 						SizeInMB: to.Ptr[int64](30270),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
