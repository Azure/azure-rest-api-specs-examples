package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/VirtualHardDisks_Get.json
func ExampleVirtualHardDisksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualHardDisksClient().Get(ctx, "test-rg", "test-vhd", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhcivm.VirtualHardDisksClientGetResponse{
	// 	VirtualHardDisk: &armazurestackhcivm.VirtualHardDisk{
	// 		Name: to.Ptr("test-vhd"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/virtualHardDisks"),
	// 		ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
	// 			Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 			Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
	// 		},
	// 		ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
	// 		Location: to.Ptr("West US2"),
	// 		Properties: &armazurestackhcivm.VirtualHardDiskProperties{
	// 			BlockSizeBytes: to.Ptr[int32](0),
	// 			DiskFileFormat: to.Ptr(armazurestackhcivm.DiskFileFormatVhdx),
	// 			DiskSizeGB: to.Ptr[int64](32),
	// 			Dynamic: to.Ptr(true),
	// 			HyperVGeneration: to.Ptr(armazurestackhcivm.HyperVGenerationV2),
	// 			LogicalSectorBytes: to.Ptr[int32](512),
	// 			PhysicalSectorBytes: to.Ptr[int32](512),
	// 			ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
	// 			Status: &armazurestackhcivm.VirtualHardDiskStatus{
	// 				UploadStatus: &armazurestackhcivm.VirtualHardDiskUploadStatus{
	// 					UploadedSizeInMB: to.Ptr[int64](10240),
	// 					ProgressPercentage: to.Ptr[int64](55),
	// 					Status: to.Ptr(armazurestackhcivm.StatusInProgress),
	// 				},
	// 				DownloadStatus: &armazurestackhcivm.VirtualHardDiskDownloadStatus{
	// 					DownloadedSizeInMB: to.Ptr[int64](10240),
	// 					ProgressPercentage: to.Ptr[int64](100),
	// 					Status: to.Ptr(armazurestackhcivm.StatusSucceeded),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
