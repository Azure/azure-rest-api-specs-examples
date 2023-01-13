package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-11-01/examples/virtualMachineExamples/VirtualMachine_Create_PlatformImageVmWithUnmanagedOsAndDataDisks.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAPlatformImageVmWithUnmanagedOsAndDataDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachinesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "{vm-name}", armcompute.VirtualMachine{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2V2),
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
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
						DiskSizeGB:   to.Ptr[int32](1023),
						Lun:          to.Ptr[int32](0),
						Vhd: &armcompute.VirtualHardDisk{
							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd"),
						},
					},
					{
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
						DiskSizeGB:   to.Ptr[int32](1023),
						Lun:          to.Ptr[int32](1),
						Vhd: &armcompute.VirtualHardDisk{
							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd"),
						},
					}},
				ImageReference: &armcompute.ImageReference{
					Offer:     to.Ptr("WindowsServer"),
					Publisher: to.Ptr("MicrosoftWindowsServer"),
					SKU:       to.Ptr("2016-Datacenter"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Name:         to.Ptr("myVMosdisk"),
					Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					Vhd: &armcompute.VirtualHardDisk{
						URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
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
