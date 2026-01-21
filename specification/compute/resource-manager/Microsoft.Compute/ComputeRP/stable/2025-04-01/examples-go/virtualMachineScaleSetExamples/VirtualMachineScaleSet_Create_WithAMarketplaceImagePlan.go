package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithAMarketplaceImagePlan.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWithAMarketplaceImagePlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "{vmss-name}", armcompute.VirtualMachineScaleSet{
		Location: to.Ptr("westus"),
		Plan: &armcompute.Plan{
			Name:      to.Ptr("windows2016"),
			Product:   to.Ptr("windows-data-science-vm"),
			Publisher: to.Ptr("microsoft-ads"),
		},
		Properties: &armcompute.VirtualMachineScaleSetProperties{
			Overprovision: to.Ptr(true),
			UpgradePolicy: &armcompute.UpgradePolicy{
				Mode: to.Ptr(armcompute.UpgradeModeManual),
			},
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							Name: to.Ptr("{vmss-name}"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								EnableIPForwarding: to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										Name: to.Ptr("{vmss-name}"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
											},
										},
									}},
								Primary: to.Ptr(true),
							},
						}},
				},
				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
					AdminPassword:      to.Ptr("{your-password}"),
					AdminUsername:      to.Ptr("{your-username}"),
					ComputerNamePrefix: to.Ptr("{vmss-name}"),
				},
				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
					ImageReference: &armcompute.ImageReference{
						Offer:     to.Ptr("windows-data-science-vm"),
						Publisher: to.Ptr("microsoft-ads"),
						SKU:       to.Ptr("windows2016"),
						Version:   to.Ptr("latest"),
					},
					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
						Caching:      to.Ptr(armcompute.CachingTypesReadWrite),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
					},
				},
			},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("Standard_D1_v2"),
			Capacity: to.Ptr[int64](3),
			Tier:     to.Ptr("Standard"),
		},
	}, &armcompute.VirtualMachineScaleSetsClientBeginCreateOrUpdateOptions{IfMatch: nil,
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
	// res.VirtualMachineScaleSet = armcompute.VirtualMachineScaleSet{
	// 	Name: to.Ptr("{vmss-name}"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 	Location: to.Ptr("westus"),
	// 	Plan: &armcompute.Plan{
	// 		Name: to.Ptr("standard-data-science-vm"),
	// 		Product: to.Ptr("standard-data-science-vm"),
	// 		Publisher: to.Ptr("microsoft-ads"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 		Overprovision: to.Ptr(true),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		SinglePlacementGroup: to.Ptr(true),
	// 		UniqueID: to.Ptr("b9e23088-6ffc-46e0-9e02-b0a6eeef47db"),
	// 		UpgradePolicy: &armcompute.UpgradePolicy{
	// 			Mode: to.Ptr(armcompute.UpgradeModeManual),
	// 		},
	// 		VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 			NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 				NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 					{
	// 						Name: to.Ptr("{vmss-name}"),
	// 						Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 							DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 								DNSServers: []*string{
	// 								},
	// 							},
	// 							EnableAcceleratedNetworking: to.Ptr(false),
	// 							EnableIPForwarding: to.Ptr(true),
	// 							IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 								{
	// 									Name: to.Ptr("{vmss-name}"),
	// 									Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 										PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 										Subnet: &armcompute.APIEntityReference{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/nsgExistingVnet/subnets/nsgExistingSubnet"),
	// 										},
	// 									},
	// 							}},
	// 							Primary: to.Ptr(true),
	// 						},
	// 				}},
	// 			},
	// 			OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 				AdminUsername: to.Ptr("{your-username}"),
	// 				ComputerNamePrefix: to.Ptr("{vmss-name}"),
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 					ProvisionVMAgent: to.Ptr(true),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Offer: to.Ptr("standard-data-science-vm"),
	// 					Publisher: to.Ptr("microsoft-ads"),
	// 					SKU: to.Ptr("standard-data-science-vm"),
	// 					Version: to.Ptr("latest"),
	// 				},
	// 				OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// 	SKU: &armcompute.SKU{
	// 		Name: to.Ptr("Standard_D1_v2"),
	// 		Capacity: to.Ptr[int64](3),
	// 		Tier: to.Ptr("Standard"),
	// 	},
	// }
}
