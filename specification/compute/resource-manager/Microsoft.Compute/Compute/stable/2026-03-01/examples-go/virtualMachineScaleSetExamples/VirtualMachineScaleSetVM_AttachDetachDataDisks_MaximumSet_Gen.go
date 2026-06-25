package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_AttachDetachDataDisks_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginAttachDetachDataDisks_virtualMachineScaleSetVMAttachDetachDataDisksMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetVMsClient().BeginAttachDetachDataDisks(ctx, "rgcompute", "azure-vmscaleset", "0", armcompute.AttachDetachDataDisksRequest{
		DataDisksToAttach: []*armcompute.DataDisksToAttach{
			{
				Lun:    to.Ptr[int32](1),
				DiskID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
				DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
				},
				Caching:                 to.Ptr(armcompute.CachingTypesReadOnly),
				WriteAcceleratorEnabled: to.Ptr(true),
			},
			{
				Lun:    to.Ptr[int32](2),
				DiskID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_2_disk3_7d5e664bdafa49baa780eb2d128ff38e"),
				DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
				},
				Caching:                 to.Ptr(armcompute.CachingTypesReadWrite),
				WriteAcceleratorEnabled: to.Ptr(false),
			},
		},
		DataDisksToDetach: []*armcompute.DataDisksToDetach{
			{
				DiskID:       to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_1_disk1_1a4e784bdafa49baa780eb2d128ff65x"),
				DetachOption: to.Ptr(armcompute.DiskDetachOptionTypesForceDetach),
			},
			{
				DiskID:       to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_4_disk4_4d4e784bdafa49baa780eb2d256ff41z"),
				DetachOption: to.Ptr(armcompute.DiskDetachOptionTypesForceDetach),
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
	// res = armcompute.VirtualMachineScaleSetVMsClientAttachDetachDataDisksResponse{
	// 	StorageProfile: armcompute.StorageProfile{
	// 		ImageReference: &armcompute.ImageReference{
	// 			Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 			Offer: to.Ptr("WindowsServer"),
	// 			SKU: to.Ptr("2016-Datacenter"),
	// 			Version: to.Ptr("latest"),
	// 		},
	// 		OSDisk: &armcompute.OSDisk{
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			Name: to.Ptr("myOsDisk"),
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 			Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 			ManagedDisk: &armcompute.ManagedDiskParameters{
	// 				StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myOsDisk"),
	// 			},
	// 			DiskSizeGB: to.Ptr[int32](30),
	// 		},
	// 		DataDisks: []*armcompute.DataDisk{
	// 			{
	// 				Lun: to.Ptr[int32](1),
	// 				Name: to.Ptr("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
	// 				Caching: to.Ptr(armcompute.CachingTypesReadOnly),
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 					DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
	// 					},
	// 				},
	// 				DiskSizeGB: to.Ptr[int32](30),
	// 				WriteAcceleratorEnabled: to.Ptr(true),
	// 			},
	// 			{
	// 				Lun: to.Ptr[int32](2),
	// 				Name: to.Ptr("vmss3176_vmss3176_2_disk3_7d5e664bdafa49baa780eb2d128ff38e"),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
	// 				Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_2_disk3_7d5e664bdafa49baa780eb2d128ff38e"),
	// 					DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
	// 					},
	// 				},
	// 				DiskSizeGB: to.Ptr[int32](100),
	// 				WriteAcceleratorEnabled: to.Ptr(false),
	// 			},
	// 		},
	// 	},
	// }
}
