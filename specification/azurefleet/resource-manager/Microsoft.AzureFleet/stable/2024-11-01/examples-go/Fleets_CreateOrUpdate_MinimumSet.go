package armcomputefleet_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet"
)

// Generated from example definition: D:/w/Azure/azure-sdk-for-go/sdk/resourcemanager/computefleet/armcomputefleet/TempTypeSpecFiles/AzureFleet.Management/examples/2024-11-01/Fleets_CreateOrUpdate_MinimumSet.json
func ExampleFleetsClient_BeginCreateOrUpdate_fleetsCreateOrUpdateMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputefleet.NewClientFactory("1DC2F28C-A625-4B0E-9748-9885A3C9E9EB", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFleetsClient().BeginCreateOrUpdate(ctx, "rgazurefleet", "testFleet", armcomputefleet.Fleet{
		Properties: &armcomputefleet.FleetProperties{
			SpotPriorityProfile: &armcomputefleet.SpotPriorityProfile{
				Capacity:           to.Ptr[int32](2),
				MinCapacity:        to.Ptr[int32](1),
				EvictionPolicy:     to.Ptr(armcomputefleet.EvictionPolicyDelete),
				AllocationStrategy: to.Ptr(armcomputefleet.SpotAllocationStrategyPriceCapacityOptimized),
				Maintain:           to.Ptr(true),
			},
			RegularPriorityProfile: &armcomputefleet.RegularPriorityProfile{
				Capacity:           to.Ptr[int32](2),
				MinCapacity:        to.Ptr[int32](1),
				AllocationStrategy: to.Ptr(armcomputefleet.RegularPriorityAllocationStrategyLowestPrice),
			},
			VMSizesProfile: []*armcomputefleet.VMSizeProfile{
				{
					Name: to.Ptr("Standard_D2s_v3"),
				},
				{
					Name: to.Ptr("Standard_D4s_v3"),
				},
				{
					Name: to.Ptr("Standard_E2s_v3"),
				},
			},
			ComputeProfile: &armcomputefleet.ComputeProfile{
				BaseVirtualMachineProfile: &armcomputefleet.BaseVirtualMachineProfile{
					StorageProfile: &armcomputefleet.VirtualMachineScaleSetStorageProfile{
						ImageReference: &armcomputefleet.ImageReference{
							Publisher: to.Ptr("canonical"),
							Offer:     to.Ptr("0001-com-ubuntu-server-focal"),
							SKU:       to.Ptr("20_04-lts-gen2"),
							Version:   to.Ptr("latest"),
						},
						OSDisk: &armcomputefleet.VirtualMachineScaleSetOSDisk{
							Caching:      to.Ptr(armcomputefleet.CachingTypesReadWrite),
							CreateOption: to.Ptr(armcomputefleet.DiskCreateOptionTypesFromImage),
							OSType:       to.Ptr(armcomputefleet.OperatingSystemTypesLinux),
							ManagedDisk: &armcomputefleet.VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: to.Ptr(armcomputefleet.StorageAccountTypesStandardLRS),
							},
						},
					},
					OSProfile: &armcomputefleet.VirtualMachineScaleSetOSProfile{
						ComputerNamePrefix: to.Ptr("prefix"),
						AdminUsername:      to.Ptr("azureuser"),
						AdminPassword:      to.Ptr("TestPassword$0"),
						LinuxConfiguration: &armcomputefleet.LinuxConfiguration{
							DisablePasswordAuthentication: to.Ptr(false),
						},
					},
					NetworkProfile: &armcomputefleet.VirtualMachineScaleSetNetworkProfile{
						NetworkInterfaceConfigurations: []*armcomputefleet.VirtualMachineScaleSetNetworkConfiguration{
							{
								Name: to.Ptr("vmNameTest"),
								Properties: &armcomputefleet.VirtualMachineScaleSetNetworkConfigurationProperties{
									Primary:                     to.Ptr(true),
									EnableAcceleratedNetworking: to.Ptr(false),
									IPConfigurations: []*armcomputefleet.VirtualMachineScaleSetIPConfiguration{
										{
											Name: to.Ptr("vmNameTest"),
											Properties: &armcomputefleet.VirtualMachineScaleSetIPConfigurationProperties{
												Subnet: &armcomputefleet.APIEntityReference{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"),
												},
												Primary: to.Ptr(true),
												LoadBalancerBackendAddressPools: []*armcomputefleet.SubResource{
													{
														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"),
													},
												},
											},
										},
									},
									EnableIPForwarding: to.Ptr(true),
								},
							},
						},
					},
				},
				ComputeAPIVersion:        to.Ptr("2023-09-01"),
				PlatformFaultDomainCount: to.Ptr[int32](1),
			},
		},
		Tags: map[string]*string{
			"key": to.Ptr("fleets-test"),
		},
		Location: to.Ptr("eastus2euap"),
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
	// res = armcomputefleet.FleetsClientCreateOrUpdateResponse{
	// 	Fleet: &armcomputefleet.Fleet{
	// 		Properties: &armcomputefleet.FleetProperties{
	// 			ProvisioningState: to.Ptr(armcomputefleet.ProvisioningStateSucceeded),
	// 			SpotPriorityProfile: &armcomputefleet.SpotPriorityProfile{
	// 				Capacity: to.Ptr[int32](2),
	// 				MinCapacity: to.Ptr[int32](1),
	// 				EvictionPolicy: to.Ptr(armcomputefleet.EvictionPolicyDelete),
	// 				AllocationStrategy: to.Ptr(armcomputefleet.SpotAllocationStrategyPriceCapacityOptimized),
	// 				Maintain: to.Ptr(true),
	// 			},
	// 			RegularPriorityProfile: &armcomputefleet.RegularPriorityProfile{
	// 				Capacity: to.Ptr[int32](2),
	// 				MinCapacity: to.Ptr[int32](1),
	// 				AllocationStrategy: to.Ptr(armcomputefleet.RegularPriorityAllocationStrategyLowestPrice),
	// 			},
	// 			VMSizesProfile: []*armcomputefleet.VMSizeProfile{
	// 				{
	// 					Name: to.Ptr("Standard_D2s_v3"),
	// 				},
	// 				{
	// 					Name: to.Ptr("Standard_D4s_v3"),
	// 				},
	// 				{
	// 					Name: to.Ptr("Standard_E2s_v3"),
	// 				},
	// 			},
	// 			ComputeProfile: &armcomputefleet.ComputeProfile{
	// 				BaseVirtualMachineProfile: &armcomputefleet.BaseVirtualMachineProfile{
	// 					StorageProfile: &armcomputefleet.VirtualMachineScaleSetStorageProfile{
	// 						ImageReference: &armcomputefleet.ImageReference{
	// 							Publisher: to.Ptr("canonical"),
	// 							Offer: to.Ptr("0001-com-ubuntu-server-focal"),
	// 							SKU: to.Ptr("20_04-lts-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSDisk: &armcomputefleet.VirtualMachineScaleSetOSDisk{
	// 							Caching: to.Ptr(armcomputefleet.CachingTypesReadWrite),
	// 							CreateOption: to.Ptr(armcomputefleet.DiskCreateOptionTypesFromImage),
	// 							OSType: to.Ptr(armcomputefleet.OperatingSystemTypesLinux),
	// 							ManagedDisk: &armcomputefleet.VirtualMachineScaleSetManagedDiskParameters{
	// 								StorageAccountType: to.Ptr(armcomputefleet.StorageAccountTypesStandardLRS),
	// 							},
	// 						},
	// 					},
	// 					OSProfile: &armcomputefleet.VirtualMachineScaleSetOSProfile{
	// 						ComputerNamePrefix: to.Ptr("prefix"),
	// 						AdminUsername: to.Ptr("azureuser"),
	// 						LinuxConfiguration: &armcomputefleet.LinuxConfiguration{
	// 							DisablePasswordAuthentication: to.Ptr(false),
	// 						},
	// 					},
	// 					NetworkProfile: &armcomputefleet.VirtualMachineScaleSetNetworkProfile{
	// 						NetworkInterfaceConfigurations: []*armcomputefleet.VirtualMachineScaleSetNetworkConfiguration{
	// 							{
	// 								Name: to.Ptr("vmNameTest"),
	// 								Properties: &armcomputefleet.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 									Primary: to.Ptr(true),
	// 									EnableAcceleratedNetworking: to.Ptr(false),
	// 									IPConfigurations: []*armcomputefleet.VirtualMachineScaleSetIPConfiguration{
	// 										{
	// 											Name: to.Ptr("vmNameTest"),
	// 											Properties: &armcomputefleet.VirtualMachineScaleSetIPConfigurationProperties{
	// 												Subnet: &armcomputefleet.APIEntityReference{
	// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"),
	// 												},
	// 												Primary: to.Ptr(true),
	// 												LoadBalancerBackendAddressPools: []*armcomputefleet.SubResource{
	// 													{
	// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/backendAddressPools/{backendAddressPoolName}"),
	// 													},
	// 												},
	// 											},
	// 										},
	// 									},
	// 									EnableIPForwarding: to.Ptr(true),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				ComputeAPIVersion: to.Ptr("2023-09-01"),
	// 				PlatformFaultDomainCount: to.Ptr[int32](1),
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-10T00:16:34.8590555+00:00"); return t}()),
	// 			UniqueID: to.Ptr("a2f7fabd-bbc2-4a20-afe1-49fdb3885a28"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key": to.Ptr("fleets-test"),
	// 		},
	// 		Location: to.Ptr("eastus2euap"),
	// 		ID: to.Ptr("/subscriptions/7B0CD4DB-3381-4013-9B31-FB6E6FD0FF1C/resourceGroups/rgazurefleet/providers/Microsoft.AzureFleet/fleets/testFleet"),
	// 		Name: to.Ptr("testFleet"),
	// 		Type: to.Ptr("Microsoft.AzureFleet/fleets"),
	// 	},
	// }
}
