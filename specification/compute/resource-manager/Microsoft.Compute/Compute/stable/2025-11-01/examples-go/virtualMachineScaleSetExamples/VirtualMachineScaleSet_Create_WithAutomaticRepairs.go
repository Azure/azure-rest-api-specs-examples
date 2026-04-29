package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithAutomaticRepairs.json
func ExampleVirtualMachineScaleSetsClient_BeginCreateOrUpdate_createAScaleSetWithAutomaticRepairsEnabled() {
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
						SKU:       to.Ptr("2016-Datacenter"),
						Publisher: to.Ptr("MicrosoftWindowsServer"),
						Version:   to.Ptr("latest"),
						Offer:     to.Ptr("WindowsServer"),
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
			ScheduledEventsPolicy: &armcompute.ScheduledEventsPolicy{
				ScheduledEventsAdditionalPublishingTargets: &armcompute.ScheduledEventsAdditionalPublishingTargets{
					EventGridAndResourceGraph: &armcompute.EventGridAndResourceGraph{
						Enable:                    to.Ptr(true),
						ScheduledEventsAPIVersion: to.Ptr("2020-07-01"),
					},
				},
				UserInitiatedRedeploy: &armcompute.UserInitiatedRedeploy{
					AutomaticallyApprove: to.Ptr(true),
				},
				UserInitiatedReboot: &armcompute.UserInitiatedReboot{
					AutomaticallyApprove: to.Ptr(true),
				},
				AllInstancesDown: &armcompute.AllInstancesDown{
					AutomaticallyApprove: to.Ptr(true),
				},
			},
			AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
				Enabled:     to.Ptr(true),
				GracePeriod: to.Ptr("PT10M"),
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
	// res = armcompute.VirtualMachineScaleSetsClientCreateOrUpdateResponse{
	// 	VirtualMachineScaleSet: &armcompute.VirtualMachineScaleSet{
	// 		SKU: &armcompute.SKU{
	// 			Tier: to.Ptr("Standard"),
	// 			Capacity: to.Ptr[int64](3),
	// 			Name: to.Ptr("Standard_D1_v2"),
	// 		},
	// 		Name: to.Ptr("{vmss-name}"),
	// 		Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 			SinglePlacementGroup: to.Ptr(true),
	// 			Overprovision: to.Ptr(true),
	// 			UniqueID: to.Ptr("d053ec5a-8da6-495f-ab13-38216503c6d7"),
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 					ImageReference: &armcompute.ImageReference{
	// 						SKU: to.Ptr("2016-Datacenter"),
	// 						Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 						Version: to.Ptr("latest"),
	// 						Offer: to.Ptr("WindowsServer"),
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
	// 			ScheduledEventsPolicy: &armcompute.ScheduledEventsPolicy{
	// 				ScheduledEventsAdditionalPublishingTargets: &armcompute.ScheduledEventsAdditionalPublishingTargets{
	// 					EventGridAndResourceGraph: &armcompute.EventGridAndResourceGraph{
	// 						Enable: to.Ptr(true),
	// 						ScheduledEventsAPIVersion: to.Ptr("2020-07-01"),
	// 					},
	// 				},
	// 				UserInitiatedRedeploy: &armcompute.UserInitiatedRedeploy{
	// 					AutomaticallyApprove: to.Ptr(true),
	// 				},
	// 				UserInitiatedReboot: &armcompute.UserInitiatedReboot{
	// 					AutomaticallyApprove: to.Ptr(true),
	// 				},
	// 				AllInstancesDown: &armcompute.AllInstancesDown{
	// 					AutomaticallyApprove: to.Ptr(true),
	// 				},
	// 			},
	// 			AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
	// 				Enabled: to.Ptr(true),
	// 				GracePeriod: to.Ptr("PT10M"),
	// 			},
	// 			ProvisioningState: to.Ptr("Creating"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}"),
	// 	},
	// }
}
