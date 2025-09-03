package armazurestackhcivm_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhcivm"
)

// Generated from example definition: 2025-06-01-preview/VirtualHardDisks_ListAll.json
func ExampleVirtualHardDisksClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhcivm.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualHardDisksClient().NewListAllPager(nil)
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
		// page = armazurestackhcivm.VirtualHardDisksClientListAllResponse{
		// 	VirtualHardDiskListResult: armazurestackhcivm.VirtualHardDiskListResult{
		// 		Value: []*armazurestackhcivm.VirtualHardDisk{
		// 			{
		// 				Name: to.Ptr("test-vhd"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/virtualHardDisks"),
		// 				ExtendedLocation: &armazurestackhcivm.ExtendedLocation{
		// 					Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
		// 					Type: to.Ptr(armazurestackhcivm.ExtendedLocationTypesCustomLocation),
		// 				},
		// 				ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
		// 				Location: to.Ptr("West US2"),
		// 				Properties: &armazurestackhcivm.VirtualHardDiskProperties{
		// 					BlockSizeBytes: to.Ptr[int32](0),
		// 					DiskFileFormat: to.Ptr(armazurestackhcivm.DiskFileFormatVhdx),
		// 					DiskSizeGB: to.Ptr[int64](32),
		// 					Dynamic: to.Ptr(true),
		// 					HyperVGeneration: to.Ptr(armazurestackhcivm.HyperVGenerationV2),
		// 					LogicalSectorBytes: to.Ptr[int32](512),
		// 					PhysicalSectorBytes: to.Ptr[int32](512),
		// 					ProvisioningState: to.Ptr(armazurestackhcivm.ProvisioningStateEnumAccepted),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
