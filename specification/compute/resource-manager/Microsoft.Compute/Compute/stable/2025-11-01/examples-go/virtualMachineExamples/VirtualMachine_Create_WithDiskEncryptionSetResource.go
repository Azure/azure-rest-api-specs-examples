package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithDiskEncryptionSetResource.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVMWithDiskEncryptionSetResourceIdInTheOSDiskAndDataDisk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachinesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myVM", armcompute.VirtualMachine{
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineProperties{
			HardwareProfile: &armcompute.HardwareProfile{
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD1V2),
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
				},
				OSDisk: &armcompute.OSDisk{
					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
						},
					},
					Name:         to.Ptr("myVMosdisk"),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
				},
				DataDisks: []*armcompute.DataDisk{
					{
						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
						ManagedDisk: &armcompute.ManagedDiskParameters{
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
							},
						},
						DiskSizeGB:   to.Ptr[int32](1023),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
						Lun:          to.Ptr[int32](0),
					},
					{
						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
						ManagedDisk: &armcompute.ManagedDiskParameters{
							ID:                 to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-managed-disk-name}"),
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
							},
						},
						DiskSizeGB:   to.Ptr[int32](1023),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
						Lun:          to.Ptr[int32](1),
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
	// 				LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 					DisablePasswordAuthentication: to.Ptr(false),
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
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/nsgcustom"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					Name: to.Ptr("myVMosdisk"),
	// 					DiskSizeGB: to.Ptr[int32](30),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 						DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskencryptionset-name}"),
	// 						},
	// 					},
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 					{
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
	// 							},
	// 						},
	// 						DiskSizeGB: to.Ptr[int32](1023),
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 						Lun: to.Ptr[int32](0),
	// 					},
	// 					{
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{existing-managed-disk-name}"),
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"),
	// 							},
	// 						},
	// 						DiskSizeGB: to.Ptr[int32](1023),
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesAttach),
	// 						Lun: to.Ptr[int32](1),
	// 					},
	// 				},
	// 			},
	// 			VMID: to.Ptr("71aa3d5a-d73d-4970-9182-8580433b2865"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardD1V2),
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
