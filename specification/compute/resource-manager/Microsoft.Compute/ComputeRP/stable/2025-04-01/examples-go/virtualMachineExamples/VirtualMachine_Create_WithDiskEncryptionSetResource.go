package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_Create_WithDiskEncryptionSetResource.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVmWithDiskEncryptionSetResourceIdInTheOsDiskAndDataDisk() {
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
	// 			VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD1V2),
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
	// 			LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 				DisablePasswordAuthentication: to.Ptr(false),
	// 			},
	// 			Secrets: []*armcompute.VaultSecretGroup{
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StorageProfile: &armcompute.StorageProfile{
	// 			DataDisks: []*armcompute.DataDisk{
	// 				{
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 					DiskSizeGB: to.Ptr[int32](1023),
	// 					Lun: to.Ptr[int32](0),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
	// 						},
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 				},
	// 				{
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
	// 					DiskSizeGB: to.Ptr[int32](1023),
	// 					Lun: to.Ptr[int32](1),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-managed-disk-name}"),
	// 						DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
	// 						},
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 			}},
	// 			ImageReference: &armcompute.ImageReference{
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/nsgcustom"),
	// 			},
	// 			OSDisk: &armcompute.OSDisk{
	// 				Name: to.Ptr("myVMosdisk"),
	// 				Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 				DiskSizeGB: to.Ptr[int32](30),
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskencryptionset-name}"),
	// 					},
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 				},
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 			},
	// 		},
	// 		VMID: to.Ptr("71aa3d5a-d73d-4970-9182-8580433b2865"),
	// 	},
	// }
}
