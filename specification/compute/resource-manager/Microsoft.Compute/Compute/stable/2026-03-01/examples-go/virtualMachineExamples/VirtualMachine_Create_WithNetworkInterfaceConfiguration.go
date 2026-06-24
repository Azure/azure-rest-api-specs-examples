package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineExamples/VirtualMachine_Create_WithNetworkInterfaceConfiguration.json
func ExampleVirtualMachinesClient_BeginCreateOrUpdate_createAVMWithNetworkInterfaceConfiguration() {
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
					SKU:       to.Ptr("2016-Datacenter"),
					Publisher: to.Ptr("MicrosoftWindowsServer"),
					Version:   to.Ptr("latest"),
					Offer:     to.Ptr("WindowsServer"),
				},
				OSDisk: &armcompute.OSDisk{
					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
					ManagedDisk: &armcompute.ManagedDiskParameters{
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
					},
					Name:         to.Ptr("myVMosdisk"),
					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
				},
			},
			NetworkProfile: &armcompute.NetworkProfile{
				NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
				NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
					{
						Name: to.Ptr("{nic-config-name}"),
						Tags: map[string]*string{
							"nicTag": to.Ptr("tag"),
						},
						Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
							Primary:      to.Ptr(true),
							DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
							IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
								{
									Name: to.Ptr("{ip-config-name}"),
									Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
										Primary: to.Ptr(true),
										PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
											Name: to.Ptr("{publicIP-config-name}"),
											Tags: map[string]*string{
												"pipTag": to.Ptr("tag"),
											},
											SKU: &armcompute.PublicIPAddressSKU{
												Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
												Tier: to.Ptr(armcompute.PublicIPAddressSKUTierGlobal),
											},
											Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
												DeleteOption:             to.Ptr(armcompute.DeleteOptionsDetach),
												PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodStatic),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			OSProfile: &armcompute.OSProfile{
				AdminUsername: to.Ptr("{your-username}"),
				ComputerName:  to.Ptr("myVM"),
				AdminPassword: to.Ptr("{your-password}"),
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
	// res = armcompute.VirtualMachinesClientCreateOrUpdateResponse{
	// 	VirtualMachine: armcompute.VirtualMachine{
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
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/toBeCreatedNetworkInterface"),
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
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Name: to.Ptr("myVMosdisk"),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 			},
	// 			VMID: to.Ptr("b7a098cc-b0b8-46e8-a205-62f301a62a8f"),
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
