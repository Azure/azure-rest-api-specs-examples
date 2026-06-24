package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithInterconnectBlock.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createOrUpdateAScaleSetWithInterconnectBlock() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vmss-name}", armcompute.VirtualMachineScaleSet{
		SKU: &armcompute.SKU{
			Tier:     to.Ptr("Standard"),
			Capacity: to.Ptr[int64](3),
			Name:     to.Ptr("Standard_ND128isr_GB300_v6"),
		},
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineScaleSetProperties{
			Overprovision:                  to.Ptr(true),
			HighSpeedInterconnectPlacement: to.Ptr(armcompute.HighSpeedInterconnectPlacementTrunk),
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &armcompute.ImageReference{
						Publisher: to.Ptr("microsoft-dsvm"),
						Offer:     to.Ptr("ubuntu-hpc"),
						SKU:       to.Ptr("2404-gb"),
						Version:   to.Ptr("latest"),
					},
					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
						},
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					},
				},
				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr("{vmss-name}"),
					AdminUsername:      to.Ptr("{your-username}"),
					AdminPassword:      to.Ptr("{your-password}"),
					LinuxConfiguration: &armcompute.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(false),
					},
				},
				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
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
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name: to.Ptr("{vmss-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								Primary:            to.Ptr(true),
								EnableIPForwarding: to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr("{vmss-name}"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
											},
										},
									},
								},
							},
						},
					},
				},
				InterconnectBlockProfile: &armcompute.InterconnectBlockProfile{
					InterconnectBlock: &armcompute.APIEntityReference{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock"),
					},
				},
			},
			UpgradePolicy: &armcompute.UpgradePolicy{
				Mode: to.Ptr(armcompute.UpgradeModeManual),
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
	// res = armcompute.VirtualMachineScaleSetsClientCreateOrUpdateResponse{
	// 	VirtualMachineScaleSet: armcompute.VirtualMachineScaleSet{
	// 		SKU: &armcompute.SKU{
	// 			Tier: to.Ptr("Standard"),
	// 			Capacity: to.Ptr[int64](3),
	// 			Name: to.Ptr("Standard_ND128isr_GB300_v6"),
	// 		},
	// 		Name: to.Ptr("{vmss-name}"),
	// 		Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 			SinglePlacementGroup: to.Ptr(false),
	// 			OrchestrationMode: to.Ptr(armcompute.OrchestrationModeUniform),
	// 			HighSpeedInterconnectPlacement: to.Ptr(armcompute.HighSpeedInterconnectPlacementTrunk),
	// 			Overprovision: to.Ptr(true),
	// 			UniqueID: to.Ptr("d053ec5a-8da6-495f-ab13-38216503c6d7"),
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 					ImageReference: &armcompute.ImageReference{
	// 						Publisher: to.Ptr("microsoft-dsvm"),
	// 						Offer: to.Ptr("ubuntu-hpc"),
	// 						SKU: to.Ptr("2404-gb"),
	// 						Version: to.Ptr("latest"),
	// 					},
	// 					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 						},
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					},
	// 				},
	// 				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 					ComputerNamePrefix: to.Ptr("{vmss-name}"),
	// 					AdminUsername: to.Ptr("{your-username}"),
	// 					Secrets: []*armcompute.VaultSecretGroup{
	// 					},
	// 					LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 						DisablePasswordAuthentication: to.Ptr(false),
	// 						ProvisionVMAgent: to.Ptr(true),
	// 					},
	// 				},
	// 				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 					InterconnectGroupProfile: &armcompute.InterconnectGroupProfile{
	// 						InterconnectGroup: &armcompute.SubResource{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup"),
	// 						},
	// 						Subgroups: []*armcompute.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0"),
	// 							},
	// 						},
	// 					},
	// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 						{
	// 							Name: to.Ptr("{vmss-name}"),
	// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 									DNSServers: []*string{
	// 									},
	// 								},
	// 								Primary: to.Ptr(true),
	// 								EnableIPForwarding: to.Ptr(true),
	// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 									{
	// 										Name: to.Ptr("{vmss-name}"),
	// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 											Subnet: &armcompute.APIEntityReference{
	// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/nsgExistingVnet/subnets/nsgExistingSubnet"),
	// 											},
	// 											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 										},
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				InterconnectBlockProfile: &armcompute.InterconnectBlockProfile{
	// 					InterconnectBlock: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock"),
	// 					},
	// 				},
	// 				SecurityProfile: &armcompute.SecurityProfile{
	// 					SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 				},
	// 			},
	// 			UpgradePolicy: &armcompute.UpgradePolicy{
	// 				Mode: to.Ptr(armcompute.UpgradeModeManual),
	// 			},
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 	},
	// }
}
