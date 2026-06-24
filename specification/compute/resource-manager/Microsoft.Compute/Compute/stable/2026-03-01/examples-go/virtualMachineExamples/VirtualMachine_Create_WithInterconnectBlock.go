package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithInterconnectBlock.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createOrUpdateAVMWithInterconnectBlock() {
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
				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_ND128isr_GB300_v6")),
			},
			StorageProfile: &armcompute.StorageProfile{
				ImageReference: &armcompute.ImageReference{
					Publisher: to.Ptr("microsoft-dsvm"),
					Offer:     to.Ptr("ubuntu-hpc"),
					SKU:       to.Ptr("2404-gb"),
					Version:   to.Ptr("latest"),
				},
				OSDisk: &armcompute.OSDisk{
					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
					},
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					Name:         to.Ptr("myVMosdisk"),
				},
			},
			InterconnectBlockProfile: &armcompute.InterconnectBlockProfile{
				InterconnectBlock: &armcompute.APIEntityReference{
					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock"),
				},
			},
			OSProfile: &armcompute.OSProfile{
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
				AdminPassword: to.Ptr("{your-password}"),
				LinuxConfiguration: &armcompute.LinuxConfiguration{
					DisablePasswordAuthentication: to.Ptr(false),
				},
			},
			NetworkProfile: &armcompute.NetworkProfile{
				InterconnectGroupProfile: &armcompute.InterconnectGroupProfile{
					InterconnectGroup: &armcompute.SubResource{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup"),
					},
					Subgroups: []*armcompute.SubResource{
						{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0"),
						},
					},
				},
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
		Zones: []*string{
			to.Ptr("1"),
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
	// res = armcompute.VirtualMachinesClientCreateOrUpdateResponse{
	// 	VirtualMachine: armcompute.VirtualMachine{
	// 		Name: to.Ptr("myVM"),
	// 		Properties: &armcompute.VirtualMachineProperties{
	// 			OSProfile: &armcompute.OSProfile{
	// 				AdminUsername: to.Ptr("{your-username}"),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				ComputerName: to.Ptr("myVM"),
	// 				LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 					DisablePasswordAuthentication: to.Ptr(false),
	// 					ProvisionVMAgent: to.Ptr(true),
	// 				},
	// 			},
	// 			InterconnectBlockProfile: &armcompute.InterconnectBlockProfile{
	// 				InterconnectBlock: &armcompute.APIEntityReference{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock"),
	// 				},
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				InterconnectGroupProfile: &armcompute.InterconnectGroupProfile{
	// 					InterconnectGroup: &armcompute.SubResource{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup"),
	// 					},
	// 					Subgroups: []*armcompute.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0"),
	// 						},
	// 					},
	// 				},
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
	// 					Publisher: to.Ptr("microsoft-dsvm"),
	// 					Offer: to.Ptr("ubuntu-hpc"),
	// 					SKU: to.Ptr("2404-gb"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 					},
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Name: to.Ptr("myVMosdisk"),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 			},
	// 			VMID: to.Ptr("5c0d55a7-c407-4ed6-bf7d-ddb810267c85"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_ND128isr_GB300_v6")),
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 		Location: to.Ptr("westus"),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 	},
	// }
}
