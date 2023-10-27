package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9ec0fcc278aa2128c4fbb2b8a1aa93432d72cce0/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/GetVirtualHardDisk.json
func ExampleVirtualHardDisksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.VirtualHardDisks = armazurestackhci.VirtualHardDisks{
	// 	Name: to.Ptr("test-vhd"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/virtualHardDisks"),
	// 	ID: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.AzureStackHCI/virtualHardDisks/test-vhd"),
	// 	Location: to.Ptr("West US2"),
	// 	ExtendedLocation: &armazurestackhci.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
	// 		Type: to.Ptr(armazurestackhci.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armazurestackhci.VirtualHardDiskProperties{
	// 		BlockSizeBytes: to.Ptr[int32](0),
	// 		DiskFileFormat: to.Ptr(armazurestackhci.DiskFileFormatVhdx),
	// 		DiskSizeGB: to.Ptr[int64](32),
	// 		Dynamic: to.Ptr(true),
	// 		HyperVGeneration: to.Ptr(armazurestackhci.HyperVGenerationV2),
	// 		LogicalSectorBytes: to.Ptr[int32](512),
	// 		PhysicalSectorBytes: to.Ptr[int32](512),
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateEnumSucceeded),
	// 	},
	// }
}
