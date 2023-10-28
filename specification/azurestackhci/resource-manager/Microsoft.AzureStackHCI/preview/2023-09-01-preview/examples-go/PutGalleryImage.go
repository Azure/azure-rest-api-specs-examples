package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutGalleryImage.json
func ExampleGalleryImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImagesClient().BeginCreateOrUpdate(ctx, "test-rg", "test-gallery-image", armazurestackhci.GalleryImages{
		Location: to.Ptr("West US2"),
		ExtendedLocation: &armazurestackhci.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
			Type: to.Ptr(armazurestackhci.ExtendedLocationTypesCustomLocation),
		},
		Properties: &armazurestackhci.GalleryImageProperties{
			ContainerID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
			ImagePath:   to.Ptr("C:\\test.vhdx"),
			OSType:      to.Ptr(armazurestackhci.OperatingSystemTypesLinux),
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
	// res.GalleryImages = armazurestackhci.GalleryImages{
	// 	Name: to.Ptr("test-gallery-image"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/galleryImages"),
	// 	ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
	// 	Location: to.Ptr("West US2"),
	// 	ExtendedLocation: &armazurestackhci.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 		Type: to.Ptr(armazurestackhci.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurestackhci.GalleryImageProperties{
	// 		CloudInitDataSource: to.Ptr(armazurestackhci.CloudInitDataSourceNoCloud),
	// 		ContainerID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
	// 		HyperVGeneration: to.Ptr(armazurestackhci.HyperVGenerationV2),
	// 		OSType: to.Ptr(armazurestackhci.OperatingSystemTypesLinux),
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateEnumSucceeded),
	// 		Status: &armazurestackhci.GalleryImageStatus{
	// 			DownloadStatus: &armazurestackhci.GalleryImageStatusDownloadStatus{
	// 				DownloadSizeInMB: to.Ptr[int64](9383),
	// 			},
	// 			ProgressPercentage: to.Ptr[int64](100),
	// 			ProvisioningStatus: &armazurestackhci.GalleryImageStatusProvisioningStatus{
	// 				OperationID: to.Ptr("79cfc696-44f5-4a68-a620-21850f7e9fb0"),
	// 				Status: to.Ptr(armazurestackhci.StatusSucceeded),
	// 			},
	// 		},
	// 		Version: &armazurestackhci.GalleryImageVersion{
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
