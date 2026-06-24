package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_FromWithNetworkInterfaceWithDnsSettings.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWithNetworkInterfacesWithPublicIPAddressDnsSettings() {
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
			Name:     to.Ptr("Standard_D1_v2"),
		},
		Location: to.Ptr("westus"),
		Properties: &armcompute.VirtualMachineScaleSetProperties{
			Overprovision: to.Ptr(true),
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &armcompute.ImageReference{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/{existing-custom-image-name}"),
					},
					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
					},
				},
				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
					ComputerNamePrefix: to.Ptr("{vmss-name}"),
					AdminUsername:      to.Ptr("{your-username}"),
					AdminPassword:      to.Ptr("{your-password}"),
				},
				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name: to.Ptr("{nicConfig1-name}"),
							Tags: map[string]*string{
								"nicTag": to.Ptr("tag"),
							},
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								Primary:                     to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
								DisableTCPStateTracking:     to.Ptr(true),
								EnableAcceleratedNetworking: to.Ptr(true),
								AuxiliaryMode:               to.Ptr(armcompute.NetworkInterfaceAuxiliaryModeAcceleratedConnections),
								AuxiliarySKU:                to.Ptr(armcompute.NetworkInterfaceAuxiliarySKUA1),
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
						{
							Name: to.Ptr("{nicConfig2-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								Primary:                     to.Ptr(false),
								EnableAcceleratedNetworking: to.Ptr(false),
								EnableIPForwarding:          to.Ptr(false),
								DisableTCPStateTracking:     to.Ptr(false),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr("{nicConfig2-name}"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											Primary: to.Ptr(true),
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-fpga-subnet-name2}"),
											},
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
												Name: to.Ptr("publicip"),
												Tags: map[string]*string{
													"pipTag": to.Ptr("tag"),
												},
												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
													IdleTimeoutInMinutes: to.Ptr[int32](10),
													DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
														DomainNameLabel:      to.Ptr("vmsstestlabel01"),
														DomainNameLabelScope: to.Ptr(armcompute.DomainNameLabelScopeTypesNoReuse),
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			UpgradePolicy: &armcompute.UpgradePolicy{
				Mode: to.Ptr(armcompute.UpgradeModeManual),
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
	// res = armcompute.VirtualMachineScaleSetsClientCreateOrUpdateResponse{
	// 	VirtualMachineScaleSet: armcompute.VirtualMachineScaleSet{
	// 		SKU: &armcompute.SKU{
	// 			Tier: to.Ptr("Standard"),
	// 			Capacity: to.Ptr[int64](3),
	// 			Name: to.Ptr("Standard_D1_v2"),
	// 		},
	// 		Name: to.Ptr("{vmss-name}"),
	// 		Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 			SinglePlacementGroup: to.Ptr(true),
	// 			Overprovision: to.Ptr(true),
	// 			UniqueID: to.Ptr("afa2afa8-9e49-48fb-9d18-c86323b5d064"),
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 					ImageReference: &armcompute.ImageReference{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/images/nsgcustom"),
	// 					},
	// 					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
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
	// 					},
	// 				},
	// 				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 						{
	// 							Name: to.Ptr("{nicConfig1-name}"),
	// 							Tags: map[string]*string{
	// 								"nicTag": to.Ptr("tag"),
	// 							},
	// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 									DNSServers: []*string{
	// 									},
	// 								},
	// 								Primary: to.Ptr(true),
	// 								EnableIPForwarding: to.Ptr(true),
	// 								DisableTCPStateTracking: to.Ptr(true),
	// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 									{
	// 										Name: to.Ptr("{nicConfig1-name}"),
	// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 											Subnet: &armcompute.APIEntityReference{
	// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/nsgExistingVnet/subnets/nsgExistingSubnet"),
	// 											},
	// 											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
	// 												Name: to.Ptr("publicip"),
	// 												Tags: map[string]*string{
	// 													"pipTag": to.Ptr("tag"),
	// 												},
	// 												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
	// 													IdleTimeoutInMinutes: to.Ptr[int32](10),
	// 													DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
	// 														DomainNameLabel: to.Ptr("vmsstestlabel01"),
	// 														DomainNameLabelScope: to.Ptr(armcompute.DomainNameLabelScopeTypesTenantReuse),
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 								EnableAcceleratedNetworking: to.Ptr(true),
	// 								AuxiliaryMode: to.Ptr(armcompute.NetworkInterfaceAuxiliaryModeAcceleratedConnections),
	// 								AuxiliarySKU: to.Ptr(armcompute.NetworkInterfaceAuxiliarySKUA1),
	// 							},
	// 						},
	// 						{
	// 							Name: to.Ptr("{nicConfig2-name}"),
	// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 									DNSServers: []*string{
	// 									},
	// 								},
	// 								Primary: to.Ptr(false),
	// 								EnableFpga: to.Ptr(false),
	// 								DisableTCPStateTracking: to.Ptr(false),
	// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 									{
	// 										Name: to.Ptr("{nicConfig2-name}"),
	// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 											Primary: to.Ptr(true),
	// 											Subnet: &armcompute.APIEntityReference{
	// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-fpga-subnet-name2}"),
	// 											},
	// 											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 										},
	// 									},
	// 								},
	// 								EnableAcceleratedNetworking: to.Ptr(false),
	// 							},
	// 						},
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
	// 	},
	// }
}
