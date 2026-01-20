package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_AttachDetachDataDisks_MinimumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginAttachDetachDataDisks_virtualMachineScaleSetVmAttachDetachDataDisksMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMsClient().BeginAttachDetachDataDisks(ctx, "rgcompute", "azure-vmscaleset", "0", armcompute.AttachDetachDataDisksRequest{
		DataDisksToAttach: []*armcompute.DataDisksToAttach{
			{
				DiskID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
			}},
		DataDisksToDetach: []*armcompute.DataDisksToDetach{
			{
				DiskID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_1_disk1_1a4e784bdafa49baa780eb2d128ff65x"),
			}},
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
	// res.StorageProfile = armcompute.StorageProfile{
	// 	DataDisks: []*armcompute.DataDisk{
	// 		{
	// 			Name: to.Ptr("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 			Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
	// 			DiskSizeGB: to.Ptr[int32](30),
	// 			Lun: to.Ptr[int32](0),
	// 			ManagedDisk: &armcompute.ManagedDiskParameters{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 				StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 			},
	// 	}},
	// 	ImageReference: &armcompute.ImageReference{
	// 		Offer: to.Ptr("WindowsServer"),
	// 		Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 		SKU: to.Ptr("2016-Datacenter"),
	// 		Version: to.Ptr("latest"),
	// 	},
	// 	OSDisk: &armcompute.OSDisk{
	// 		Name: to.Ptr("myOsDisk"),
	// 		Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 		CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 		DiskSizeGB: to.Ptr[int32](30),
	// 		ManagedDisk: &armcompute.ManagedDiskParameters{
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myOsDisk"),
	// 			StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 		},
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 	},
	// }
}
