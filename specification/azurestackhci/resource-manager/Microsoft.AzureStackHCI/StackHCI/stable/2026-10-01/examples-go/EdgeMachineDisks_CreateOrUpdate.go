package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineDisks_CreateOrUpdate.json
func ExampleEdgeMachineDisksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeMachineDisksClient().BeginCreateOrUpdate(ctx, "test-rg", "EdgeMachine01", "disk-001", armazurestackhci.EdgeMachineDisk{
		Properties: &armazurestackhci.EdgeMachineDiskProperties{
			DiskConfiguration: &armazurestackhci.DiskConfiguration{
				Volumes: []*armazurestackhci.DiskVolumeConfiguration{
					{
						SizeInBytes: to.Ptr("107374182400"),
						Path:        to.Ptr("/data"),
						FileSystem:  to.Ptr(armazurestackhci.DiskFileSystemTypeExt4),
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeMachineDisksClientCreateOrUpdateResponse{
	// 	EdgeMachineDisk: armazurestackhci.EdgeMachineDisk{
	// 		Properties: &armazurestackhci.EdgeMachineDiskProperties{
	// 			DiskConfiguration: &armazurestackhci.DiskConfiguration{
	// 				Volumes: []*armazurestackhci.DiskVolumeConfiguration{
	// 					{
	// 						SizeInBytes: to.Ptr("107374182400"),
	// 						Path: to.Ptr("/data"),
	// 						FileSystem: to.Ptr(armazurestackhci.DiskFileSystemTypeExt4),
	// 					},
	// 				},
	// 			},
	// 			ReportedProperties: &armazurestackhci.DiskReportedProperties{
	// 				DiskName: to.Ptr("disk-001"),
	// 				Model: to.Ptr("Samsung SSD 980 PRO"),
	// 				Manufacturer: to.Ptr("Samsung"),
	// 				SerialNumber: to.Ptr("S5GXNF0R123456"),
	// 				FirmwareVersion: to.Ptr("5B2QGXA7"),
	// 				SizeInBytes: to.Ptr("1000204886016"),
	// 				DiskType: to.Ptr(armazurestackhci.DiskTypeNVMe),
	// 				BusLocation: to.Ptr("0000:03:00.0"),
	// 				UnallocatedSizeInBytes: to.Ptr("892830703616"),
	// 				State: to.Ptr(armazurestackhci.DiskStateOnline),
	// 				Volumes: []*armazurestackhci.DiskVolumeReportedProperties{
	// 					{
	// 						SizeInBytes: to.Ptr("107374182400"),
	// 						Path: to.Ptr("/data"),
	// 						FileSystem: to.Ptr("ext4"),
	// 						PartitionID: to.Ptr("vol-001"),
	// 					},
	// 				},
	// 				IsSupported: to.Ptr(true),
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/disks/disk-001"),
	// 		Name: to.Ptr("disk-001"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/disks"),
	// 	},
	// }
}
