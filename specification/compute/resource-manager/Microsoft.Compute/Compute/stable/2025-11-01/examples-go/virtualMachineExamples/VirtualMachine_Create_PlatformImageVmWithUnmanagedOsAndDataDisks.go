package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_PlatformImageVmWithUnmanagedOsAndDataDisks.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAPlatformImageVMWithUnmanagedOSAndDataDisks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vm-name}", armcompute.VirtualMachine{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2V2),
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					SKU:       to.Ptr("2016-Datacenter"),
					Publisher: to.Ptr("MicrosoftWindowsServer"),
					Version:   to.Ptr("latest"),
					Offer:     to.Ptr("WindowsServer"),
				},
				OSDisk: &armcompute.OSDisk{
					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
					Vhd: &armcompute.VirtualHardDisk{
						URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
					},
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					Name:         to.Ptr("myVMosdisk"),
				},
				DataDisks: []*armcompute.DataDisk{
					{
						DiskSizeGB:   to.Ptr[int32](1023),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
						Lun:          to.Ptr[int32](0),
						Vhd: &armcompute.VirtualHardDisk{
							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk0.vhd"),
						},
					},
					{
						DiskSizeGB:   to.Ptr[int32](1023),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
						Lun:          to.Ptr[int32](1),
						Vhd: &armcompute.VirtualHardDisk{
							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk1.vhd"),
						},
					},
				},
			},
			OSProfile: &armcompute.OSProfile{
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
				AdminPassword: to.Ptr("{your-password}"),
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
					{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}"),
						Properties: &armcompute.NetworkInterfaceReferenceProperties{
							Primary: to.Ptr(true),
						},
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachinesClientCreateOrUpdateResponse{
	// 	VirtualMachine: &armcompute.VirtualMachine{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		Properties: &armcompute.VirtualMachineProperties{
	// 			OSProfile: &armcompute.OSProfile{
	// 				AdminUsername: to.Ptr("{your-username}"),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				ComputerName: to.Ptr("myVM"),
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 				},
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/nsgExistingNic"),
	// 						Properties: &armcompute.NetworkInterfaceReferenceProperties{
	// 							Primary: to.Ptr(true),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					SKU: to.Ptr("2016-Datacenter"),
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					Version: to.Ptr("latest"),
	// 					Offer: to.Ptr("WindowsServer"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					Vhd: &armcompute.VirtualHardDisk{
	// 						URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/vhds/myDisk.vhd"),
	// 					},
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Name: to.Ptr("myVMosdisk"),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 					{
	// 						Name: to.Ptr("dataDisk0"),
	// 						DiskSizeGB: to.Ptr[int32](1023),
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 						Caching: to.Ptr(armcompute.CachingTypesNone),
	// 						Vhd: &armcompute.VirtualHardDisk{
	// 							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/vhds/myDisk0.vhd"),
	// 						},
	// 						Lun: to.Ptr[int32](0),
	// 					},
	// 					{
	// 						Name: to.Ptr("dataDisk1"),
	// 						DiskSizeGB: to.Ptr[int32](1023),
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 						Caching: to.Ptr(armcompute.CachingTypesNone),
	// 						Vhd: &armcompute.VirtualHardDisk{
	// 							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/vhds/myDisk1.vhd"),
	// 						},
	// 						Lun: to.Ptr[int32](1),
	// 					},
	// 				},
	// 			},
	// 			VMID: to.Ptr("5230a749-2f68-4830-900b-702182d32e63"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD2V2),
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 		Name: to.Ptr("myVM"),
	// 		Location: to.Ptr("westus"),
	// 	},
	// }
}
