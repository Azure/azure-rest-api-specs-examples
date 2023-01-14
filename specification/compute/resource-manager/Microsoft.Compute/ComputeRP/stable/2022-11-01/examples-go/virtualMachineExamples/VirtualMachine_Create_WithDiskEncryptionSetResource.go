package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineExamples/VirtualMachine_Create_WithDiskEncryptionSetResource.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVmWithDiskEncryptionSetResourceIdInTheOsDiskAndDataDisk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachinesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myVM", armcompute.VirtualMachine{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD1V2),
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
					}},
			},
			OSProfile: &armcompute.OSProfile{
				AdminPassword: to.Ptr("{your-password}"),
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
			},
			StorageProfile: &armcompute.StorageProfile{
				DataDisks: []*armcompute.DataDisk{
					{
						Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
						DiskSizeGB:   to.Ptr[int32](1023),
						Lun:          to.Ptr[int32](0),
						ManagedDisk: &armcompute.ManagedDiskParameters{
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
							},
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
					},
					{
						Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
						DiskSizeGB:   to.Ptr[int32](1023),
						Lun:          to.Ptr[int32](1),
						ManagedDisk: &armcompute.ManagedDiskParameters{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-managed-disk-name}"),
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
							},
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
					}},
				ImageReference: &armcompute.ImageReference{
					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr("myVMosdisk"),
					Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
						},
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
