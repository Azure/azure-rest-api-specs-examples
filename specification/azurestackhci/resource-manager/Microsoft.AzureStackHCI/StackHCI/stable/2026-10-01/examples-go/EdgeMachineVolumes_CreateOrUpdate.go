package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineVolumes_CreateOrUpdate.json
func ExampleEdgeMachineVolumesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeMachineVolumesClient().BeginCreateOrUpdate(ctx, "test-rg", "EdgeMachine01", "vol-001", armazurestackhci.EdgeMachineVolume{
		Properties: &armazurestackhci.EdgeMachineVolumeProperties{
			VolumeConfiguration: &armazurestackhci.VolumeConfiguration{
				Reserved: to.Ptr(""),
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
	// res = armazurestackhci.EdgeMachineVolumesClientCreateOrUpdateResponse{
	// 	EdgeMachineVolume: armazurestackhci.EdgeMachineVolume{
	// 		Properties: &armazurestackhci.EdgeMachineVolumeProperties{
	// 			VolumeConfiguration: &armazurestackhci.VolumeConfiguration{
	// 				Reserved: to.Ptr(""),
	// 			},
	// 			ReportedProperties: &armazurestackhci.VolumeReportedProperties{
	// 				PartitionID: to.Ptr("partition-001"),
	// 				OffsetInBytes: to.Ptr[int64](1048576),
	// 				Name: to.Ptr("DataVolume"),
	// 				Path: to.Ptr("/data"),
	// 				ParentDiskID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/disks/disk-001"),
	// 				IsBoot: to.Ptr(false),
	// 				SizeInBytes: to.Ptr("107374182400"),
	// 				SizeRemainingInBytes: to.Ptr("85899345920"),
	// 				FileSystem: to.Ptr("ext4"),
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/volumes/vol-001"),
	// 		Name: to.Ptr("vol-001"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/volumes"),
	// 	},
	// }
}
