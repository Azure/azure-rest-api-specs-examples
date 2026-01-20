package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_Create_PlatformImageVmWithUnmanagedOsAndDataDisks.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAPlatformImageVmWithUnmanagedOsAndDataDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vm-name}", armcompute.VirtualMachine{
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
	}, &armcompute.VirtualMachinesClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.VirtualMachine = armcompute.VirtualMachine{
	// 	Name: to.Ptr("myVM"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armcompute.VirtualMachineProperties{
	// 		HardwareProfile: &armcompute.HardwareProfile{
	// 			VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2V2),
	// 		},
	// 		NetworkProfile: &armcompute.NetworkProfile{
	// 			NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/nsgExistingNic"),
	// 					Properties: &armcompute.NetworkInterfaceReferenceProperties{
	// 						Primary: to.Ptr(true),
	// 					},
	// 			}},
	// 		},
	// 		OSProfile: &armcompute.OSProfile{
	// 			AdminUsername: to.Ptr("{your-username}"),
	// 			ComputerName: to.Ptr("myVM"),
	// 			Secrets: []*armcompute.VaultSecretGroup{
	// 			},
	// 			WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 				EnableAutomaticUpdates: to.Ptr(true),
	// 				ProvisionVMAgent: to.Ptr(true),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StorageProfile: &armcompute.StorageProfile{
	// 			DataDisks: []*armcompute.DataDisk{
	// 				{
	// 					Name: to.Ptr("dataDisk0"),
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 					DiskSizeGB: to.Ptr[int32](1023),
	// 					Lun: to.Ptr[int32](0),
	// 					Vhd: &armcompute.VirtualHardDisk{
	// 						URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/vhds/myDisk0.vhd"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("dataDisk1"),
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 					DiskSizeGB: to.Ptr[int32](1023),
	// 					Lun: to.Ptr[int32](1),
	// 					Vhd: &armcompute.VirtualHardDisk{
	// 						URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/vhds/myDisk1.vhd"),
	// 					},
	// 			}},
	// 			ImageReference: &armcompute.ImageReference{
	// 				Offer: to.Ptr("WindowsServer"),
	// 				Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 				SKU: to.Ptr("2016-Datacenter"),
	// 				Version: to.Ptr("latest"),
	// 			},
	// 			OSDisk: &armcompute.OSDisk{
	// 				Name: to.Ptr("myVMosdisk"),
	// 				Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 				Vhd: &armcompute.VirtualHardDisk{
	// 					URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/vhds/myDisk.vhd"),
	// 				},
	// 			},
	// 		},
	// 		VMID: to.Ptr("5230a749-2f68-4830-900b-702182d32e63"),
	// 	},
	// }
}
