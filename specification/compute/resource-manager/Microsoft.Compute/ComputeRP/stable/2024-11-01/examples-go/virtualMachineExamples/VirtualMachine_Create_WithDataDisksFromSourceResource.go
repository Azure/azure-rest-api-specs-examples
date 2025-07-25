package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/virtualMachineExamples/VirtualMachine_Create_WithDataDisksFromSourceResource.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVmWithDataDisksUsingCopyAndRestoreOptions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myVM", armcompute.VirtualMachine{
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
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesCopy),
						DiskSizeGB:   to.Ptr[int32](1023),
						Lun:          to.Ptr[int32](0),
						SourceResource: &armcompute.APIEntityReference{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/{existing-snapshot-name}"),
						},
					},
					{
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesCopy),
						DiskSizeGB:   to.Ptr[int32](1023),
						Lun:          to.Ptr[int32](1),
						SourceResource: &armcompute.APIEntityReference{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-disk-name}"),
						},
					},
					{
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesRestore),
						DiskSizeGB:   to.Ptr[int32](1023),
						Lun:          to.Ptr[int32](2),
						SourceResource: &armcompute.APIEntityReference{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/{existing-rpc-name}/restorePoints/{existing-rp-name}/diskRestorePoints/{existing-disk-restore-point-name}"),
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
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
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
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesCopy),
	// 					DiskSizeGB: to.Ptr[int32](1023),
	// 					Lun: to.Ptr[int32](0),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 					SourceResource: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/{existing-snapshot-name}"),
	// 					},
	// 				},
	// 				{
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesCopy),
	// 					DiskSizeGB: to.Ptr[int32](1023),
	// 					Lun: to.Ptr[int32](1),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 					SourceResource: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-disk-name}"),
	// 					},
	// 				},
	// 				{
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesRestore),
	// 					DiskSizeGB: to.Ptr[int32](1023),
	// 					Lun: to.Ptr[int32](2),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 					SourceResource: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/restorePointCollections/{existing-rpc-name}/restorePoints/{existing-rp-name}/diskRestorePoints/{existing-disk-restore-point-name}"),
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
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 				},
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			},
	// 		},
	// 		VMID: to.Ptr("3906fef9-a1e5-4b83-a8a8-540858b41df0"),
	// 	},
	// }
}
