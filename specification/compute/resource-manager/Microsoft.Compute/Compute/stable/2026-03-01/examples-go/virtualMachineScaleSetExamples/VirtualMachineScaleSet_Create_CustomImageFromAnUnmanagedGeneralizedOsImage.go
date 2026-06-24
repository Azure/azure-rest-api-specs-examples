package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_CustomImageFromAnUnmanagedGeneralizedOsImage.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createACustomImageScaleSetFromAnUnmanagedGeneralizedOSImage() {
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
					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
						Image: &armcompute.VirtualHardDisk{
							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/{existing-generalized-os-image-blob-name}.vhd"),
						},
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
						Name:         to.Ptr("osDisk"),
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
	// 			UniqueID: to.Ptr("d6e9ab29-f8c9-4792-978c-ae2c07b98f17"),
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 						OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						Image: &armcompute.VirtualHardDisk{
	// 							URI: to.Ptr("https://{existing-storage-account-name}.blob.core.windows.net/system/Microsoft.Compute/Images/vhds/{existing-generalized-os-image-blob-name}.vhd"),
	// 						},
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 						Name: to.Ptr("osDisk"),
	// 					},
	// 				},
	// 				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 					ComputerNamePrefix: to.Ptr("{vmss-name}"),
	// 					AdminUsername: to.Ptr("{your-username}"),
	// 					Secrets: []*armcompute.VaultSecretGroup{
	// 					},
	// 					WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 						ProvisionVMAgent: to.Ptr(true),
	// 						EnableAutomaticUpdates: to.Ptr(true),
	// 					},
	// 				},
	// 				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
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
