package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/GalleryImages_CreateOrUpdate.json
func ExampleGalleryImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImagesClient().BeginCreateOrUpdate(ctx, "test-rg", "test-gallery-image", armazurestackhcivm.GalleryImage{
		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
		},
		Location: to.Ptr("West US2"),
		Properties: &armazurestackhcivm.GalleryImageProperties{
			ContainerID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
			ImagePath:   to.Ptr("C:\\test.vhdx"),
			OSType:      to.Ptr(armazurestackhcivm.OperatingSystemTypesLinux),
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
	// res = armazurestackhcivm.GalleryImagesClientCreateOrUpdateResponse{
	// 	GalleryImage: &armazurestackhcivm.GalleryImage{
	// 		Name: to.Ptr("test-gallery-image"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/galleryImages"),
	// 		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/galleryImages/test-gallery-image"),
	// 		Location: to.Ptr("West US2"),
	// 		Properties: &armazurestackhcivm.GalleryImageProperties{
	// 			CloudInitDataSource: to.Ptr(armazurestackhcivm.CloudInitDataSourceNoCloud),
	// 			ContainerID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/storageContainers/test-storage-container"),
	// 			HyperVGeneration: to.Ptr(armazurestackhcivm.HyperVGenerationV2),
	// 			OSType: to.Ptr(armazurestackhcivm.OperatingSystemTypesLinux),
	// 			ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
	// 			Status: &armazurestackhcivm.GalleryImageStatus{
	// 				DownloadStatus: &armazurestackhcivm.GalleryImageStatusDownloadStatus{
	// 					DownloadSizeInMB: to.Ptr[int64](9383),
	// 				},
	// 				ProgressPercentage: to.Ptr[int64](100),
	// 				ProvisioningStatus: &armazurestackhcivm.GalleryImageStatusProvisioningStatus{
	// 					OperationID: to.Ptr("79cfc696-44f5-4a68-a620-21850f7e9fb0"),
	// 					Status: to.Ptr(armazurestackhcivm.StatusSucceeded),
	// 				},
	// 			},
	// 			Version: &armazurestackhcivm.GalleryImageVersion{
	// 				Properties: &armazurestackhcivm.GalleryImageVersionProperties{
	// 					StorageProfile: &armazurestackhcivm.GalleryImageVersionStorageProfile{
	// 						OSDiskImage: &armazurestackhcivm.GalleryOSDiskImage{
	// 							SizeInMB: to.Ptr[int64](30270),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
